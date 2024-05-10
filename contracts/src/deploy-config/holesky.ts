const config = {
    // Global configuration
    l1FeeVaultRecipient: '0x4844c339e748fD44825856E5B877a7A9B5BcE559',
    contractAdmin: '0x4844c339e748fD44825856E5B877a7A9B5BcE559',

    // chainID config
    l1ChainID: 17000,
    l2ChainID: 2810,

    // L1MessageQueue config
    l1MessageQueueMaxGasLimit: 100000000,

    // gasPriceOracle config
    l2BaseFee: 0.1,  // Gwei 

    // rollup contract config
    // initialize config
    rollupMaxNumTxInChunk: 100,
    finalizationPeriodSeconds: 600,
    rollupProofWindow: 86400,
    // challenge config
    // TODO to be removed later
    rollupProposer: '0x48446b07A04307F320fa34790116d73aD4A51810',
    rollupChallenger: '0x48445d5DDb25a4f23C6d85D87f15B39Fe94e6867',
    // genesis config
    rollupGenesisStateRoot: '',
    withdrawRoot: '',
    batchHeader: '',

    // staking config
    // staking initialize config
    stakingSequencerSize: 4,
    stakingLockNumber: 3,
    stakingMinDeposit: 1, // limit
    // register sequencers
    l2SequencerAddresses: [
        "0x48446b07A04307F320fa34790116d73aD4A51810",
        "0x48447a9d66D8eAE067e59E3F9f7E6B5f1b5E7581",
        "0x4844631425Ed76D79164c37CA9d558bd5FB12662",
        "0x4844ede12C743845dfae266A6BE891F114702583",
        "0x4844741d4D5c7f943365368e9a9c891Da1EBBb24",
        "0x4844C03f82fd15765Dd2547dfF7CDb05F62EF735",
        "0x4844934E51b52b0f8a4200752028c4a6414c45c6"
    ],
    l2SequencerTmKeys: [
        "0xa4f3a7fb882c3159398f520e3c89a39dc9818235fea07789995a9ce7b4aba31d",
        "0x0535824f72998615ac8265c12c9cfc63d315d8f3e849aa1857f42b11cfc5d72c",
        "0x3a50197cbfdbe5bbd85ff13b0d55ddc08ecd5dbddbc012880ba1358562e192bd",
        "0x814ce757e9727768f25ead4619efa552d9f8facad85236c0a774d0b21fae28ea",
        "0x9d2992e94cc50cd8bb881f5eb3ad3828cfbb5e2bf922d988f311df0971a6cc9c",
        "0x61237ca5a45b3b9aebd6cfaacdc0020775eaeecbd0c4c61021659b6d581f606f",
        "0xa8ec67073d74afa3791c369d388c8dc2eb63eb3f519f3975b7e894e0bba155f6"
    ],
    l2SequencerBlsKeys: [
        "0x0000000000000000000000000000000017c22eb22cfacc028f8556a10ed0dcf56f4189c03c57103498bebc38df3f04202a1bf2ae6361bd0fd0c606e5a5d1122500000000000000000000000000000000115359ea50bb5e0bc9a72d5d96dbe5425aad06b7c4f9b18c30534d8237de9f3d94df3fd38e2e19bf572cda4e76dfd7200000000000000000000000000000000001584eb7275c2678f41b1649c4de05acecffa73f998c30995e934b01029257a6466964b68d4ea34c64ac1807f57bad0c0000000000000000000000000000000016d8f19d0fb3f0428c5fe338ca6bdb4b93e49f106267a106b11bde1bc4b795a119cbd8bfcd21d5be79b83fabc43c1d1e",
        "0x0000000000000000000000000000000018ca6c24a492491ccbc97b0743c52e188ccbf4c42fe858a22ece10c96d99af4c80a3c73b2b199f1bc35adf2b42be86e700000000000000000000000000000000028ebe1896404b3a9f3609d61b55c8a13aed4c2ba951fb081d0982bebcd072ea2fa5a90bd581565e4277be1b86b201c80000000000000000000000000000000009f5d61b5b90288ba4650db5865235935447e4999ed773965db1cc663d14bd30e364244f8f97cbdb9483591366d11db00000000000000000000000000000000015d1d71cb633e1a33b7e7432b268d94453704f0fcd54dbf223d294c969c4ea9285d807a42d785d6198d3a361be98fdf7",
        "0x000000000000000000000000000000000c666fbafcde1b34ac662d6edd8ddaca29fb406c27bdb4b45114409d1d6b3a1415efcb88a9adf94a925f255a7e3b4cb300000000000000000000000000000000097f204c34019966d2f1f2e57572e534d51f93f5c2e372bf15bdb031eefa72398fb5c773c217365df980c9c6b3480962000000000000000000000000000000000b1d419b65cf9d415ae3dc213f41296623517efeea8a6b97ed31aacfd75a9d0c15b1b59091fa185981ce1b45b85291240000000000000000000000000000000007085c7d37c0b2ac925b75b1f86ea93cbfc60909ddabbc7a1d89cc24b1766e4bdfc3e88cf405a7cd8e8137050777a8e7",
        "0x0000000000000000000000000000000017cb2706bf579e0043486627a62d7b4b9f868585bab1eb584253e54d57cc6e8d5851d4260df693f2e671eadf5daac13b000000000000000000000000000000001176bb97fd0c8d1ecec2d240ffe9db99305ef00fdd2341590970dde0e34f26c672b5bb3ecd2e91fa66c349cba9f2f77a000000000000000000000000000000000220ff827a48e8a21e7ec9c73432fa47453d275de48fbc1988dad49e58123a91ba1e8f0463c2f6913a5efff41f3100a2000000000000000000000000000000000219da2bc157d03525b20e6ec00d2edb312ab654c107465dfe9c2b19e2c07639c02bbbe19ba5ab0202110ba7b7d72c90",
        "0x000000000000000000000000000000000448e77ab0dbdabf8a77ad9096c2f27cd1be16c15d5e15f68754c213a0155637c081667725ce0a386d76eb8f89c5101100000000000000000000000000000000031fcb97aa3e4aa9961ac769d324bdcd753dcef33b2507abf7dc33fba2e3ae49c292ec5afd9ef5df4ea1ab55a27016b60000000000000000000000000000000015573598061293e23175d4f2be97b70e47554ef2e634e3b238a36c5834681aee3cd9547966e131715c3dd997d406e01600000000000000000000000000000000183350a36939eefa711bf26a2ae8e74db7ece0e1afca17531df55bc55a55cf019fa1adaadd608943540c066aedf00dc0",
        "0x0000000000000000000000000000000000612795f4bed48dd8f8aa60474707e308db27a4221b0d16787681313f0f992bd6beeb0552e2eeea512a0acf6a20d9b6000000000000000000000000000000000be974eab4d6cead817196303207b8136f2ec1f441e381a02f1db87192c27a57fc160188e73d8576d2d7a6c06cf85ec7000000000000000000000000000000000f925f393553594cb8adbd674acc6fc9664a487aa61b26c91fa926c39979c1ca0a40440f4f064c6fa775de82c980f8bd000000000000000000000000000000000641bba2ffb13ca4dec1ecedd00ccfcebb288e7af6b9dce828398f83743c19892441d591e12c2c0e7ae36c27bcc10c42",
        "0x000000000000000000000000000000000b948f7891173ce897579e6e400ec5c0b78b51dd07ef4efe4a84d96b034c67e6a84e70a393481c5f39a5d3832bfed4770000000000000000000000000000000009d0ec3a0aab7d81a2dd853b81f2ec15a467d7a124424a7105a00b3cfda7532177cf174b58ac68a7a3f31ccd6b08b8e60000000000000000000000000000000015a97af231023c6f3468aa5e9919a11dfb51d99a0adeae22f29737c2c8132880870356221312232b93fffb886e0355c50000000000000000000000000000000004b10c90b887342c3b203aa8b55241df94feab4e634ce16490adc790b8fe19991ec485e65e3d43b9b1d4af0f6c9d039f"
    ]
}

export default config 
