const config = {
    // Global configuration
    contractAdmin: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    l1FeeVaultRecipient: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',

    // chainID config
    l1ChainID: 900,
    l2ChainID: 53077,

    // L1MessageQueue config 
    l1MessageQueueMaxGasLimit: 100000000,

    // gasPriceOracle config
    l2BaseFee: 0.1,  // Gwei 

    // rollup contract config
    // initialize config
    rollupMaxNumTxInChunk: 100,
    finalizationPeriodSeconds: 10,
    rollupProofWindow: 86400,
    // challenge config
    rollupProposer: '0x70997970C51812dc3A010C7d01b50e0d17dc79C8',
    rollupChallenger: '0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65',
    // genesis config
    rollupGenesisStateRoot: '0x06374560ac61775a9ee81644c875754fc69029a05a32becf6ebc2c801070db0e',
    withdrawRoot: '0x27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757',
    batchHeader: '0x000000000000000000000000000000000000000000000000001c886dcf76b0d43e796488c6ed2776179e359d8f567f1e06faee36357b06fa04010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c4440140000000000000000000000000000000000000000000000000000000000000000',

    // staking config
    // staking Cross-Chain config
    stakingCrossChainGaslimitAdd: 1000000,
    stakingCrossChainGaslimitRemove: 8000000,

    // staking initialize config
    stakingLockNumber: 3,
    stakingMinDeposit: 1, // limit
    stakingChallengerRewardPercentage: 100,
    // register sequencers
    l2SequencerAddresses: [
        "0x783698dCDEBdc96785c5c60ED96113612bA09c2b",
        "0x310824AA27a29D269d2F9C0a8563C0e3C98dD226",
        "0x343C5154FFe47c8a07DF5ea6846404e68E9809A2",
        "0xaaC606d51De6A5aBF0d1B9dbd5ed5Ff2Ac2e521B"
    ],
    l2SequencerTmKeys: [
        "0x5280d0eee2a64d3ad29480d15ffd1b048ce5908f180b5ccd65cc3dcf00941abb",
        "0xb798eb74c06721d54c659e9ea2bc232a7f95e96d234cc87186b2ab8f43db6935",
        "0xecffad01129786ba9c6293aa664f952894dc4019762804799dfa1ffb6e4ed040",
        "0x8d29695bb4157c6960adb486be1d9a0c9f728524d09bffcc4f8932ed15221c4a"
    ],
    l2SequencerBlsKeys: [
        "0x00000000000000000000000000000000095ad465c2895ee825c7d4f1b60a18734db57d4108369e47c6e3a94ee15846f825c06dad5d98f503bd31ece1d9f94b11000000000000000000000000000000000c5d6ba04bc9b9674dd2acbfc5caed3976c1b8be2ec90a03d78dffe924648b4fba82225aff43c744310c6a60185b75ac000000000000000000000000000000000fce6be001c871a11b9db1c6c15f0a6999de5646941a74486206dc784f0b3ffe11799212f3f44ef754b4a0f1ecf85639000000000000000000000000000000000b2f06634e5ea719682c30911c94dfb560f0b7656b5c34a871ea035e3fe7b041885420f8fe1e251f1cce5cdb7514869e",
        "0x0000000000000000000000000000000010173aeac4ff317e8e60493f962b91dbd27614e1f6594e17d18a02968bd1fd698b6703092ab8622cd22d6948d9421156000000000000000000000000000000000801aea15697ab4d7a808be45377e4f0d2f54857fdc04031e476402ff16c66a6cbcc5f09a84bf85400c8afbabed006600000000000000000000000000000000015fc71b2c4e81148274e6169c9c9aace8c34fa6030547650242b6c32527dd23a996416e32640bce4f495a0afabc7dbb900000000000000000000000000000000088c4a0dffccc96bce47aef0e176b129457a5f3ae1651b132ddb418e9f7b5850a38c6fec1be6d169eb88dc1619648bf4",
        "0x0000000000000000000000000000000003fd9468a8ceffc1b696874517777ef8bfdc9a1bade95c480ee2624903e648c1caf01c65de5b4fda8876a3a0e8d9f0890000000000000000000000000000000004c02f3609a0f61d12fe737dcbb047d5253bd3ff905b55c0e0f932b476fd77d172a58b72ef0f506407870988dd6038220000000000000000000000000000000017fa5765899f60f7a58f8ccdaaa295cde55992231710672692ba6a71a4faa9572f728f438ded65576a570d57e19fd304000000000000000000000000000000001226138813bde98af3464ed03649d8c731bc4e5cb3d26b53bf7483f4105d18bbb3f19e23905119e156e7d003d2fd125c",
        "0x00000000000000000000000000000000109bf02a2636c0dc1968b0a50db77251eb090c3e9f51e2a2bc60c4ac72213f41f01f0a34e92c2e0625bd62e28e27edb500000000000000000000000000000000139969bd92522113c0615659874d1fae311ad8152d0584c7b57ffc14927067486dcf86413c5684fccc1163ee2d45c1c1000000000000000000000000000000000f172603f70a0730d100ad6d28bde477195987062e8ade83b82d093935d956ff20ca768c26263577b094f1cb756adc400000000000000000000000000000000010dde3acca00b4ff1b4976500a8f97e92246f43f78cadc95c4993dfc4f4c501c33d42a4bf52587f4931287b59623149c"
    ],
}

export default config 
