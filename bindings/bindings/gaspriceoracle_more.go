// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"morph-l2/bindings/solc"
)

const GasPriceOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/l2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/l2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"l1BaseFee\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1002,\"contract\":\"contracts/l2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"contracts/l2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_uint256\"},{\"astId\":1004,\"contract\":\"contracts/l2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"allowListEnabled\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_bool\"},{\"astId\":1005,\"contract\":\"contracts/l2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"isAllowed\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_mapping(t_address,t_bool)\"},{\"astId\":1006,\"contract\":\"contracts/l2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"l1BlobBaseFee\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"contracts/l2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"commitScalar\",\"offset\":0,\"slot\":\"7\",\"type\":\"t_uint256\"},{\"astId\":1008,\"contract\":\"contracts/l2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"blobScalar\",\"offset\":0,\"slot\":\"8\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"contracts/l2/system/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"isCurie\",\"offset\":0,\"slot\":\"9\",\"type\":\"t_bool\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_bool\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var GasPriceOracleStorageLayout = new(solc.StorageLayout)

var GasPriceOracleDeployedBin = "0x608060405234801561000f575f80fd5b5060043610610184575f3560e01c806384189161116100dd578063de26c4a111610088578063efeadb6d11610063578063efeadb6d146102f8578063f2fde38b1461030b578063f45e65d81461031e575f80fd5b8063de26c4a1146102bf578063e3de72a5146102d2578063e88a60ad146102e5575f80fd5b8063a911d77f116100b8578063a911d77f14610282578063babcc5391461028a578063bede39b5146102ac575f80fd5b8063841891611461023f5780638da5cb5b14610248578063944b247f1461026f575f80fd5b806339455d3a1161013d5780636a5e67e5116101185780636a5e67e51461021b5780637046559714610224578063715018a614610237575f80fd5b806339455d3a146101ec57806349948e0e146101ff578063519b4bd314610212575f80fd5b806322bd5c1c1161016d57806322bd5c1c146101c157806323e524ac146101ce5780633577afc5146101d7575f80fd5b80630c18c1621461018857806313dad5be146101a4575b5f80fd5b61019160025481565b6040519081526020015b60405180910390f35b6009546101b19060ff1681565b604051901515815260200161019b565b6004546101b19060ff1681565b61019160075481565b6101ea6101e5366004610f76565b610327565b005b6101ea6101fa366004610f8d565b610432565b61019161020d366004611029565b610538565b61019160015481565b61019160085481565b6101ea610232366004610f76565b610562565b6101ea610672565b61019160065481565b5f5460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101ea61027d366004610f76565b610685565b6101ea610793565b6101b16102983660046110f9565b60056020525f908152604090205460ff1681565b6101ea6102ba366004610f76565b6108c5565b6101916102cd366004611029565b61098c565b6101ea6102e03660046111be565b6109a9565b6101ea6102f3366004610f76565b610b0e565b6101ea610306366004611278565b610c1c565b6101ea6103193660046110f9565b610cc5565b61019160035481565b336103465f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff161480610383575060045460ff1680156103835750335f9081526005602052604090205460ff165b6103b9576040517f297b056200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b621c9c388111156103f6576040517fae85900a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527f32740b35c0ea213650f60d44366b4fb211c9033b50714e4a1d34e65d5beb9bb4906020015b60405180910390a150565b336104515f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff16148061048e575060045460ff16801561048e5750335f9081526005602052604090205460ff165b6104c4576040517f297b056200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600182905560068190556040518281527f351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c449060200160405180910390a16040518181527f9a14bfb5d18c4c3cf14cae19c23d7cf1bcede357ea40ca1f75cd49542c71c2149060200160405180910390a15050565b6009545f9060ff16156105545761054e82610d81565b92915050565b61054e82610dc7565b919050565b336105815f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614806105be575060045460ff1680156105be5750335f9081526005602052604090205460ff165b6105f4576040517f297b056200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610604633b9aca006103e86112be565b81111561063d576040517f3c89fbd600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60038190556040518181527f3336cd9708eaf2769a0f0dc0679f30e80f15dcd88d1921b5a16858e8b85c591a90602001610427565b61067a610e0a565b6106835f610e8a565b565b336106a45f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614806106e1575060045460ff1680156106e15750335f9081526005602052604090205460ff165b610717576040517f297b056200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610725633b9aca00806112be565b81111561075e576040517f874f603100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60078190556040518181527f2ab3f5a4ebbcbf3c24f62f5454f52f10e1a8c9dcc5acac8f19199ce881a6a10890602001610427565b336107b25f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1614806107ef575060045460ff1680156107ef5750335f9081526005602052604090205460ff165b610825576040517f297b056200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60095460ff1615610862576040517f79f9c57500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600980547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091556040519081527f1d876a458a15cb9b74fce42f51ef6e9427d75b6f736892a2e292d93b28e7625c9060200160405180910390a1565b336108e45f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff161480610921575060045460ff1680156109215750335f9081526005602052604090205460ff165b610957576040517f297b056200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527f351fb23757bb5ea0546c85b7996ddd7155f96b939ebaa5ff7bc49c75f27f2c4490602001610427565b6009545f9060ff16156109a057505f919050565b61054e82610efe565b6109b1610e0a565b80518251146109ec576040517f1b9c61c500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8251811015610b0957818181518110610a0957610a096112d5565b602002602001015160055f858481518110610a2657610a266112d5565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550828181518110610a8f57610a8f6112d5565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167fd9739f45a01ce092c5cdb3d68f63d63d21676b1c6c0b4f9cbc6be4cf5449595a838381518110610ae057610ae06112d5565b6020026020010151604051610af9911515815260200190565b60405180910390a26001016109ee565b505050565b33610b2d5f5473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff161480610b6a575060045460ff168015610b6a5750335f9081526005602052604090205460ff165b610ba0576040517f297b056200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610bae633b9aca00806112be565b811115610be7576040517ff37ec21500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60088190556040518181527f6b332a036d8c3ead57dcb06c87243bd7a2aed015ddf2d0528c2501dae56331aa90602001610427565b610c24610e0a565b60045460ff16151581151503610c66576040517fd5d1b79c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168215159081179091556040519081527f16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb390602001610427565b610ccd610e0a565b73ffffffffffffffffffffffffffffffffffffffff8116610d75576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b610d7e81610e8a565b50565b5f633b9aca006006548351600854610d9991906112be565b610da391906112be565b600154600754610db391906112be565b610dbd9190611302565b61054e9190611315565b5f80610dd283610efe565b90505f60015482610de391906112be565b9050633b9aca0060035482610df891906112be565b610e029190611315565b949350505050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610683576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610d6c565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80515f908190815b81811015610f6757848181518110610f2057610f206112d5565b01602001517fff00000000000000000000000000000000000000000000000000000000000000165f03610f5857600483019250610f5f565b6010830192505b600101610f06565b50506002540160400192915050565b5f60208284031215610f86575f80fd5b5035919050565b5f8060408385031215610f9e575f80fd5b50508035926020909101359150565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561102157611021610fad565b604052919050565b5f602080838503121561103a575f80fd5b823567ffffffffffffffff80821115611051575f80fd5b818501915085601f830112611064575f80fd5b81358181111561107657611076610fad565b6110a6847fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610fda565b915080825286848285010111156110bb575f80fd5b80848401858401375f90820190930192909252509392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461055d575f80fd5b5f60208284031215611109575f80fd5b611112826110d6565b9392505050565b5f67ffffffffffffffff82111561113257611132610fad565b5060051b60200190565b8035801515811461055d575f80fd5b5f82601f83011261115a575f80fd5b8135602061116f61116a83611119565b610fda565b8083825260208201915060208460051b870101935086841115611190575f80fd5b602086015b848110156111b3576111a68161113c565b8352918301918301611195565b509695505050505050565b5f80604083850312156111cf575f80fd5b823567ffffffffffffffff808211156111e6575f80fd5b818501915085601f8301126111f9575f80fd5b8135602061120961116a83611119565b82815260059290921b84018101918181019089841115611227575f80fd5b948201945b8386101561124c5761123d866110d6565b8252948201949082019061122c565b96505086013592505080821115611261575f80fd5b5061126e8582860161114b565b9150509250929050565b5f60208284031215611288575f80fd5b6111128261113c565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808202811582820484141761054e5761054e611291565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b8082018082111561054e5761054e611291565b5f82611348577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b50049056fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(GasPriceOracleStorageLayoutJSON), GasPriceOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["GasPriceOracle"] = GasPriceOracleStorageLayout
	deployedBytecodes["GasPriceOracle"] = GasPriceOracleDeployedBin
}
