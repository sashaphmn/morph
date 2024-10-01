use crate::evm::{save_plonk_fixture, EvmProofFixture};
use anyhow::anyhow;
use morph_executor_client::{
    types::input::{AggregationInput, ShardInfo},
    verify_agg,
};
use morph_executor_host::get_blob_info;
use morph_executor_utils::read_env_var;
use sbv_primitives::{alloy_primitives::keccak256, types::BlockTrace, B256};
use sp1_sdk::{HashableKey, ProverClient, SP1Proof, SP1ProofWithPublicValues, SP1Stdin};
use std::time::Instant;

/// The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
pub const SHARD_VERIFIER_ELF: &[u8] =
    include_bytes!("../../client/shard/elf/riscv32im-succinct-zkvm-elf");

pub const AGG_VERIFIER_ELF: &[u8] =
    include_bytes!("../../client/agg/elf/riscv32im-succinct-zkvm-elf");

pub fn prove(
    shard_proofs: Vec<SP1ProofWithPublicValues>,
    blocks: Vec<Vec<BlockTrace>>,
    prove: bool,
) -> Result<Option<EvmProofFixture>, anyhow::Error> {
    // sp1_sdk::utils::setup_logger();
    let program_hash = keccak256(AGG_VERIFIER_ELF);
    println!("Program Hash [view on Explorer]:");
    println!("0x{}", hex::encode(program_hash));

    let client = ProverClient::new();
    let (_, vkey) = client.setup(SHARD_VERIFIER_ELF);
    let (stdin, agg_input) = get_agg_proof_stdin(blocks, shard_proofs, &vkey).unwrap();

    // Execute the program in native
    let native_rt =
        verify_agg(&agg_input).map_err(|e| anyhow!(format!("Native execution err: {:?}", e)))?;

    // Execute the program in sp1-vm
    let (mut public_values, execution_report) = client
        .execute(AGG_VERIFIER_ELF, stdin.clone())
        .run()
        .map_err(|e| anyhow!(format!("Program execution err: {:?}", e)))?;
    println!(
        "Program executed successfully, Number of cycles: {:?}",
        execution_report.total_instruction_count()
    );
    let public_values = B256::from_slice(&public_values.read::<[u8; 32]>());
    println!(
        "pi_hash generated with sp1-vm execution: {}",
        alloy::hex::encode_prefixed(public_values.as_slice())
    );
    assert_eq!(native_rt, public_values, "public_values == public_values ");
    println!("Values are correct!");

    if !prove {
        println!("Execution completed, No prove request, skipping...");
        return Ok(None);
    }

    // Start prove
    println!("Start proving...");
    // Setup the program for proving.
    let (pk, vk) = client.setup(AGG_VERIFIER_ELF);
    println!("Batch ELF Verification Key: {:?}", vk.vk.bytes32());

    // Generate plonk proof
    let start = Instant::now();
    let mut proof = client.prove(&pk, stdin).plonk().run().expect("proving failed");
    let duration_mins = start.elapsed().as_secs() / 60;
    println!("Successfully generated proof!, time use: {:?} minutes", duration_mins);

    // Verify the proof.
    client.verify(&proof, &vk).expect("failed to verify proof");
    println!("Successfully verified proof!");

    // Deserialize the public values.
    let pi_bytes = proof.public_values.read::<[u8; 32]>();
    println!("pi_hash generated with sp1-vm prove: {}", alloy::hex::encode_prefixed(pi_bytes));
    let fixture = EvmProofFixture {
        vkey: vk.bytes32().to_string(),
        public_values: B256::from_slice(&pi_bytes).to_string(),
        proof: alloy::hex::encode_prefixed(proof.bytes()),
    };

    if read_env_var("SAVE_FIXTURE", false) {
        save_plonk_fixture(&fixture);
    }
    Ok(Some(fixture))
}

/// Get the stdin for the aggregation proof.
fn get_agg_proof_stdin(
    blocks: Vec<Vec<BlockTrace>>,
    shard_proofs: Vec<SP1ProofWithPublicValues>,
    shard_vkey: &sp1_sdk::SP1VerifyingKey,
) -> Result<(SP1Stdin, AggregationInput), anyhow::Error> {
    let mut proofs: Vec<SP1Proof> = Vec::with_capacity(shard_proofs.len());
    let mut shard_infos: Vec<ShardInfo> = Vec::with_capacity(shard_proofs.len());

    for proof in shard_proofs {
        proofs.push(proof.proof.clone());
        shard_infos.push(proof.public_values.clone().read());
    }

    // Write the proofs.
    let mut stdin = SP1Stdin::new();
    for proof in proofs {
        let SP1Proof::Compressed(compressed_proof) = proof else {
            panic!();
        };
        stdin.write_proof(*compressed_proof, shard_vkey.vk.clone());
    }

    let flattened_blocks = &blocks.clone().into_iter().flatten().collect();
    let agg_input = AggregationInput {
        shard_infos,
        shard_vkey: shard_vkey.hash_u32(),
        blob_info: get_blob_info(flattened_blocks).unwrap(),
        l2_traces: blocks,
    };

    // Write the aggregation inputs to the stdin.
    stdin.write(&agg_input);

    Ok((stdin, agg_input))
}

#[cfg(test)]
mod tests {
    use morph_executor_client::{
        types::{
            blob::{decode_transactions, get_origin_batch},
            input::BlobInfo,
        },
        BlobVerifier,
    };
    use morph_executor_host::{encode_blob, populate_kzg};
    use sbv_primitives::{alloy_primitives::hex, types::TypedTransaction};
    #[test]
    fn test_blob() {
        //blob to txn
        let blob_bytes = load_zstd_blob();
        println!("blob_bytes len: {:?}", blob_bytes.len());

        let origin_batch = get_origin_batch(&blob_bytes).unwrap();
        println!("origin_batch len: {:?}", origin_batch.len());

        let tx_list: Vec<TypedTransaction> = decode_transactions(origin_batch.as_slice());
        println!("decoded tx_list_len: {:?}", tx_list.len());

        //txn to blob
        let mut tx_bytes: Vec<u8> = vec![];
        let x = tx_list.iter().flat_map(|tx| tx.rlp()).collect::<Vec<u8>>();
        tx_bytes.extend(x);
        assert!(tx_bytes == origin_batch, "tx_bytes==origin_batch");
        // tx_bytes[121..1000].fill(0);
        let blob = encode_blob(tx_bytes);

        std::env::set_var("TRUSTED_SETUP_4844", "../../configs/4844_trusted_setup.txt");
        let blob_info: BlobInfo = populate_kzg(&blob).unwrap();

        let (versioned_hash, batch_data) = BlobVerifier::verify(&blob_info).unwrap();
        println!(
            "versioned_hash: {:?}, batch_data len: {:?}",
            hex::encode(versioned_hash.as_slice()),
            batch_data.len()
        );
    }

    pub fn load_zstd_blob() -> [u8; 131072] {
        use sbv_primitives::alloy_primitives::hex;
        use std::{fs, path::Path};

        //https://holesky.etherscan.io/blob/0x018494ae7657bebd9e590baf3736ac9207a5d2275ef98c025dad3232b7875278?bid=2391294
        //https://explorer-holesky.morphl2.io/batches/223946
        let blob_data_path = Path::new("../../testdata/blob/sp1_batch.data");
        let data = fs::read_to_string(blob_data_path).expect("Unable to read file");
        let hex_data: Vec<u8> = hex::decode(data.trim()).unwrap();
        let mut array = [0u8; 131072];
        array.copy_from_slice(&hex_data);
        array
    }
}
