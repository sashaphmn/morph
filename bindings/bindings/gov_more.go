// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"morph-l2/bindings/solc"
)

const GovStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1017_storage\"},{\"astId\":1003,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1016_storage\"},{\"astId\":1005,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"batchBlockInterval\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"batchMaxBytes\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"batchTimeout\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_uint256\"},{\"astId\":1008,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"maxChunks\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"rollupEpoch\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_uint256\"},{\"astId\":1010,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"rollupEpochUpdateTime\",\"offset\":0,\"slot\":\"106\",\"type\":\"t_uint256\"},{\"astId\":1011,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"proposalInterval\",\"offset\":0,\"slot\":\"107\",\"type\":\"t_uint256\"},{\"astId\":1012,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"currentProposalID\",\"offset\":0,\"slot\":\"108\",\"type\":\"t_uint256\"},{\"astId\":1013,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"proposalData\",\"offset\":0,\"slot\":\"109\",\"type\":\"t_mapping(t_uint256,t_struct(ProposalData)1019_storage)\"},{\"astId\":1014,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"proposalInfos\",\"offset\":0,\"slot\":\"110\",\"type\":\"t_mapping(t_uint256,t_struct(ProposalInfo)1020_storage)\"},{\"astId\":1015,\"contract\":\"contracts/l2/staking/Gov.sol:Gov\",\"label\":\"votes\",\"offset\":0,\"slot\":\"111\",\"type\":\"t_mapping(t_uint256,t_struct(AddressSet)1018_storage)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_bytes32)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"bytes32[]\",\"numberOfBytes\":\"32\"},\"t_array(t_uint256)1016_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1017_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_struct(AddressSet)1018_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e struct EnumerableSetUpgradeable.AddressSet)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_struct(AddressSet)1018_storage\"},\"t_mapping(t_uint256,t_struct(ProposalData)1019_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e struct IGov.ProposalData)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_struct(ProposalData)1019_storage\"},\"t_mapping(t_uint256,t_struct(ProposalInfo)1020_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e struct IGov.ProposalInfo)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_struct(ProposalInfo)1020_storage\"},\"t_struct(AddressSet)1018_storage\":{\"encoding\":\"inplace\",\"label\":\"struct EnumerableSetUpgradeable.AddressSet\",\"numberOfBytes\":\"64\"},\"t_struct(ProposalData)1019_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IGov.ProposalData\",\"numberOfBytes\":\"160\"},\"t_struct(ProposalInfo)1020_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IGov.ProposalInfo\",\"numberOfBytes\":\"64\"},\"t_struct(Set)1021_storage\":{\"encoding\":\"inplace\",\"label\":\"struct EnumerableSetUpgradeable.Set\",\"numberOfBytes\":\"64\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var GovStorageLayout = new(solc.StorageLayout)

var GovDeployedBin = "0x608060405234801561000f575f80fd5b5060043610610179575f3560e01c80638da5cb5b116100d2578063b511328d11610088578063e5aec99511610063578063e5aec9951461038e578063ecded2ae14610397578063f2fde38b146103aa575f80fd5b8063b511328d14610334578063bb881e4114610372578063de00d3fd1461037b575f80fd5b8063929a9cbe116100b8578063929a9cbe146102b457806396dea936146102bd5780639f50395214610321575f80fd5b80638da5cb5b1461026f5780638e21d5fb1461028d575f80fd5b806349c1a5811161013257806377c793801161010d57806377c7938014610211578063807de4431461021a5780638596305214610266575f80fd5b806349c1a581146101dd5780636396619014610200578063715018a614610209575f80fd5b80632d7aa82b116101625780632d7aa82b146101a55780634063a84e146101b85780634428c1a4146101d4575f80fd5b80630121b93f1461017d5780630d61b51914610192575b5f80fd5b61019061018b366004611b1d565b6103bd565b005b6101906101a0366004611b1d565b61082b565b6101906101b3366004611b34565b61093f565b6101c1606b5481565b6040519081526020015b60405180910390f35b6101c1606a5481565b6101f06101eb366004611b94565b610e01565b60405190151581526020016101cb565b6101c1606c5481565b610190610e21565b6101c160675481565b6102417f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101cb565b6101c160655481565b60335473ffffffffffffffffffffffffffffffffffffffff16610241565b6102417f000000000000000000000000000000000000000000000000000000000000000081565b6101c160665481565b6102f96102cb366004611b1d565b606d6020525f9081526040902080546001820154600283015460038401546004909401549293919290919085565b604080519586526020860194909452928401919091526060830152608082015260a0016101cb565b61019061032f366004611b1d565b610e34565b61035d610342366004611b1d565b606e6020525f90815260409020805460019091015460ff1682565b604080519283529015156020830152016101cb565b6101c160685481565b610190610389366004611bc2565b610ef8565b6101c160695481565b6101f06103a5366004611b1d565b611256565b6101906103b8366004611bd8565b61129b565b5f73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016636d46e987336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401602060405180830381865afa158015610462573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104869190611bfa565b9050806104f4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f6f6e6c792073657175656e6365722063616e2070726f706f736500000000000060448201526064015b60405180910390fd5b5f828152606e6020526040902060010154829060ff1615610571576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f70726f706f73616c20616c726561647920617070726f7665640000000000000060448201526064016104eb565b5f818152606e60205260409020544211156105e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f70726f706f73616c206f7574206f66206461746500000000000000000000000060448201526064016104eb565b6105ff335f858152606f6020526040902090611352565b1561068c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f73657175656e63657220616c726561647920766f746520666f7220746869732060448201527f70726f706f73616c00000000000000000000000000000000000000000000000060648201526084016104eb565b6106a3335f858152606f6020526040902090611380565b505f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377d7dffb6040518163ffffffff1660e01b81526004015f60405180830381865afa15801561070d573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526107529190810190611c56565b90505f5b81518110156107e55761079b82828151811061077457610774611d34565b6020026020010151606f5f8881526020019081526020015f2061135290919063ffffffff16565b6107dd576107db8282815181106107b4576107b4611d34565b6020026020010151606f5f8881526020019081526020015f206113a190919063ffffffff16565b505b600101610756565b506003815160026107f69190611d8e565b6108009190611da5565b5f858152606f60205260409020610816906113c2565b111561082557610825846113cb565b50505050565b5f818152606e6020526040902060010154819060ff16156108a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f70726f706f73616c20616c726561647920617070726f7665640000000000000060448201526064016104eb565b5f818152606e602052604090205442111561091f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f70726f706f73616c206f7574206f66206461746500000000000000000000000060448201526064016104eb565b5f610929836116aa565b9050801561093a5761093a836113cb565b505050565b5f54610100900460ff161580801561095d57505f54600160ff909116105b806109765750303b15801561097657505f5460ff166001145b610a02576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104eb565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610a5e575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b5f8711610ac7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f696e76616c69642070726f706f73616c20696e74657276616c0000000000000060448201526064016104eb565b5f8311610b30576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c6964206d6178206368756e6b73000000000000000000000000000060448201526064016104eb565b5f8211610b99576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f696e76616c696420726f6c6c75702065706f636800000000000000000000000060448201526064016104eb565b85151580610ba657508415155b80610bb057508315155b610c16576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f696e76616c696420626174636820706172616d7300000000000000000000000060448201526064016104eb565b610c1e6117ba565b606b8790556065869055606685905560678490556068839055606982905542606a55604080515f8152602081018990527f9e890086ea51933fb82fde9166ba4d58ecb0fdb81559ee03743b7ac052f43f7b910160405180910390a1604080515f8152602081018890527fa044538eba1b21d23eb13fa35811ca9d1d7ff9ef1c81ee4dc594fca08412531b910160405180910390a1604080515f8152602081018790527f11b7e0f5b30d2753fcf7151b7a907cc343034c6a7572d56c261ae00c411d16a7910160405180910390a1604080515f8152602081018690527fab2cb47d396c5d12c082ac9b6512d332af2767ca8e1fa5bcef40fa6970626569910160405180910390a1604080515f8152602081018590527fd4cf36ce0d0f667d929d7bdf98e8774da275ea7f990c012c308516650d85839a910160405180910390a1604080515f8152602081018490527f9b20ee151d057f4f3ece7fdf4ca1370cf143f181760e7712b722572f2dcba88f910160405180910390a18015610df8575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b5f828152606f60205260408120610e189083611352565b90505b92915050565b610e29611858565b610e325f6118d9565b565b610e3c611858565b5f81118015610e4d57506065548114155b610eb3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f696e76616c6964206e65772070726f706f73616c20696e74657276616c00000060448201526064016104eb565b606580549082905560408051828152602081018490527f9e890086ea51933fb82fde9166ba4d58ecb0fdb81559ee03743b7ac052f43f7b910160405180910390a15050565b5f73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016636d46e987336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401602060405180830381865afa158015610f9d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fc19190611bfa565b90508061102a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f6f6e6c792073657175656e6365722063616e2070726f706f736500000000000060448201526064016104eb565b81608001355f03611097576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f696e76616c696420726f6c6c75702065706f636800000000000000000000000060448201526064016104eb565b5f826060013511611104576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f696e76616c6964206d6178206368756e6b73000000000000000000000000000060448201526064016104eb565b81351515806111165750602082013515155b806111245750604082013515155b61118a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f696e76616c696420626174636820706172616d7300000000000000000000000060448201526064016104eb565b606c8054905f61119983611ddd565b9091555050606c545f908152606d6020526040902082906111e5828281358155602082013560018201556040820135600282015560608201356003820155608082013560048201555050565b9050506040518060400160405280606b54426112019190611e14565b81525f6020918201819052606c548152606e82526040902082518155910151600190910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790555050565b5f818152606e602052604081206001015460ff161561127657505f919050565b5f828152606e602052604090205442111561129257505f919050565b610e1b826116aa565b6112a3611858565b73ffffffffffffffffffffffffffffffffffffffff8116611346576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104eb565b61134f816118d9565b50565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526001830160205260408120541515610e18565b5f610e188373ffffffffffffffffffffffffffffffffffffffff841661194f565b5f610e188373ffffffffffffffffffffffffffffffffffffffff841661199b565b5f610e1b825490565b5f818152606d60205260409020546065541461143a57606580545f838152606d60205260409081902054928390555190917fa044538eba1b21d23eb13fa35811ca9d1d7ff9ef1c81ee4dc594fca08412531b9161143091848252602082015260400190565b60405180910390a1505b5f818152606d6020526040902060010154606654146114af57606680545f838152606d60205260409081902060010154928390555190917f11b7e0f5b30d2753fcf7151b7a907cc343034c6a7572d56c261ae00c411d16a7916114a591848252602082015260400190565b60405180910390a1505b5f818152606d60205260409020600201546067541461152457606780545f838152606d60205260409081902060020154928390555190917fab2cb47d396c5d12c082ac9b6512d332af2767ca8e1fa5bcef40fa69706265699161151a91848252602082015260400190565b60405180910390a1505b5f818152606d60205260409020600301546068541461159957606880545f838152606d60205260409081902060030154928390555190917fd4cf36ce0d0f667d929d7bdf98e8774da275ea7f990c012c308516650d85839a9161158f91848252602082015260400190565b60405180910390a1505b5f818152606d60205260409020600401546069541461160d57606980545f838152606d6020908152604091829020600401805490945542606a55925481518381529384015290917f9b20ee151d057f4f3ece7fdf4ca1370cf143f181760e7712b722572f2dcba88f910160405180910390a1505b5f818152606e6020908152604091829020600190810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690911790556065546066546067546068546069548651948552948401929092528285015260608201526080810191909152905182917f146676d233683eb1ec2a813a7f97a7aa3241ae78af1ee6df4a4548c47178cbfa919081900360a00190a250565b5f807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377d7dffb6040518163ffffffff1660e01b81526004015f60405180830381865afa158015611714573d5f803e3d5ffd5b505050506040513d5f823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526117599190810190611c56565b90505f805b82518110156117975761177c83828151811061077457610774611d34565b1561178f5761178c826001611e14565b91505b60010161175e565b506003825160026117a89190611d8e565b6117b29190611da5565b109392505050565b5f54610100900460ff16611850576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104eb565b610e32611a7e565b60335473ffffffffffffffffffffffffffffffffffffffff163314610e32576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104eb565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f81815260018301602052604081205461199457508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610e1b565b505f610e1b565b5f8181526001830160205260408120548015611a75575f6119bd600183611e27565b85549091505f906119d090600190611e27565b9050818114611a2f575f865f0182815481106119ee576119ee611d34565b905f5260205f200154905080875f018481548110611a0e57611a0e611d34565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080611a4057611a40611e3a565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610e1b565b5f915050610e1b565b5f54610100900460ff16611b14576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104eb565b610e32336118d9565b5f60208284031215611b2d575f80fd5b5035919050565b5f805f805f8060c08789031215611b49575f80fd5b505084359660208601359650604086013595606081013595506080810135945060a0013592509050565b73ffffffffffffffffffffffffffffffffffffffff8116811461134f575f80fd5b5f8060408385031215611ba5575f80fd5b823591506020830135611bb781611b73565b809150509250929050565b5f60a08284031215611bd2575f80fd5b50919050565b5f60208284031215611be8575f80fd5b8135611bf381611b73565b9392505050565b5f60208284031215611c0a575f80fd5b81518015158114611bf3575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b8051611c5181611b73565b919050565b5f6020808385031215611c67575f80fd5b825167ffffffffffffffff80821115611c7e575f80fd5b818501915085601f830112611c91575f80fd5b815181811115611ca357611ca3611c19565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f83011681018181108582111715611ce657611ce6611c19565b604052918252848201925083810185019188831115611d03575f80fd5b938501935b82851015611d2857611d1985611c46565b84529385019392850192611d08565b98975050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8082028115828204841417610e1b57610e1b611d61565b5f82611dd8577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611e0d57611e0d611d61565b5060010190565b80820180821115610e1b57610e1b611d61565b81810381811115610e1b57610e1b611d61565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(GovStorageLayoutJSON), GovStorageLayout); err != nil {
		panic(err)
	}

	layouts["Gov"] = GovStorageLayout
	deployedBytecodes["Gov"] = GovDeployedBin
}
