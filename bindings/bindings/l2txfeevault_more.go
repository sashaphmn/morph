// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"morph-l2/bindings/solc"
)

const L2TxFeeVaultStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/l2/system/L2TxFeeVault.sol:L2TxFeeVault\",\"label\":\"owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/l2/system/L2TxFeeVault.sol:L2TxFeeVault\",\"label\":\"minWithdrawAmount\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1002,\"contract\":\"contracts/l2/system/L2TxFeeVault.sol:L2TxFeeVault\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_address\"},{\"astId\":1003,\"contract\":\"contracts/l2/system/L2TxFeeVault.sol:L2TxFeeVault\",\"label\":\"recipient\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/l2/system/L2TxFeeVault.sol:L2TxFeeVault\",\"label\":\"totalProcessed\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_uint256\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L2TxFeeVaultStorageLayout = new(solc.StorageLayout)

var L2TxFeeVaultDeployedBin = "0x6080604052600436106100c6575f3560e01c806384411d6511610071578063f2fde38b1161004c578063f2fde38b1461021e578063feec756c1461023d578063ff4f35461461025c575f80fd5b806384411d65146101bf5780638da5cb5b146101d45780639e7adc79146101ff575f80fd5b8063457e1a49116100a1578063457e1a491461015c57806366d003ac1461017f578063715018a6146101ab575f80fd5b80632e1a7d4d146100d15780633cb747bf146100f25780633ccfd60b14610148575f80fd5b366100cd57005b5f80fd5b3480156100dc575f80fd5b506100f06100eb36600461098b565b61027b565b005b3480156100fd575f80fd5b5060025461011e9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b348015610153575f80fd5b506100f06104c8565b348015610167575f80fd5b5061017160015481565b60405190815260200161013f565b34801561018a575f80fd5b5060035461011e9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156101b6575f80fd5b506100f06104d5565b3480156101ca575f80fd5b5061017160045481565b3480156101df575f80fd5b505f5461011e9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561020a575f80fd5b506100f06102193660046109a2565b610560565b348015610229575f80fd5b506100f06102383660046109a2565b610656565b348015610248575f80fd5b506100f06102573660046109a2565b61075c565b348015610267575f80fd5b506100f061027636600461098b565b610852565b600154811015610338576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f4665655661756c743a207769746864726177616c20616d6f756e74206d75737460448201527f2062652067726561746572207468616e206d696e696d756d207769746864726160648201527f77616c20616d6f756e7400000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b47808211156103c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4665655661756c743a20696e73756666696369656e742062616c616e6365207460448201527f6f20776974686472617700000000000000000000000000000000000000000000606482015260840161032f565b60048054830190556003546040805184815273ffffffffffffffffffffffffffffffffffffffff90921660208301523382820152517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a1600254600354604080516020810182525f80825291517fb2267a7b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9485169463b2267a7b9488946104969491909216928592906004016109dc565b5f604051808303818588803b1580156104ad575f80fd5b505af11580156104bf573d5f803e3d5ffd5b50505050505050565b476104d28161027b565b50565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610555576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604482015260640161032f565b61055e5f610917565b565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146105e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604482015260640161032f565b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f1c928c417a10a21c3cddad148c5dba5d710e4b1442d6d8a36de345935ad84612905f90a35050565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146106d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604482015260640161032f565b73ffffffffffffffffffffffffffffffffffffffff8116610753576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f6e6577206f776e657220697320746865207a65726f2061646472657373000000604482015260640161032f565b6104d281610917565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146107dc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604482015260640161032f565b6003805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f7e1e96961a397c8aa26162fe259cc837afc95e33aad4945ddc61c18dabb7a6ad905f90a35050565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146108d2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604482015260640161032f565b600180549082905560408051828152602081018490527f0d3c80219fe57713b9f9c83d1e51426792d0c14d8e330e65b102571816140965910160405180910390a15050565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f6020828403121561099b575f80fd5b5035919050565b5f602082840312156109b2575f80fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146109d5575f80fd5b9392505050565b73ffffffffffffffffffffffffffffffffffffffff851681525f60208560208401526080604084015284518060808501525f5b81811015610a2b5786810183015185820160a001528201610a0f565b505f60a0828601015260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116850101925050508260608301529594505050505056fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(L2TxFeeVaultStorageLayoutJSON), L2TxFeeVaultStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2TxFeeVault"] = L2TxFeeVaultStorageLayout
	deployedBytecodes["L2TxFeeVault"] = L2TxFeeVaultDeployedBin
}
