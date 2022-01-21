package proxy

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/frenchie-foundation/go-lachesis/lachesis/genesis/proxy/proxypos"
)

// GetContractBin is SFC contract first implementation bin code for mainnet
// Must be compiled with bin-runtime flag
func GetContractBin() []byte {
	return hexutil.MustDecode("0x60806040526004361061005a5760003560e01c80635c60da1b116100435780635c60da1b146101315780638f2839701461016f578063f851a440146101af5761005a565b80633659cfe6146100645780634f1ef286146100a4575b6100626101c4565b005b34801561007057600080fd5b506100626004803603602081101561008757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101de565b610062600480360360408110156100ba57600080fd5b73ffffffffffffffffffffffffffffffffffffffff82351691908101906040810160208201356401000000008111156100f257600080fd5b82018360208201111561010457600080fd5b8035906020019184600183028401116401000000008311171561012657600080fd5b509092509050610232565b34801561013d57600080fd5b50610146610306565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561017b57600080fd5b506100626004803603602081101561019257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610315565b3480156101bb57600080fd5b5061014661041d565b6101cc6101dc565b6101dc6101d7610427565b61044c565b565b6101e6610470565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102275761022281610495565b61022f565b61022f6101c4565b50565b61023a610470565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102f95761027683610495565b60008373ffffffffffffffffffffffffffffffffffffffff1683836040518083838082843760405192019450600093509091505080830381855af49150503d80600081146102e0576040519150601f19603f3d011682016040523d82523d6000602084013e6102e5565b606091505b50509050806102f357600080fd5b50610301565b6103016101c4565b505050565b6000610310610427565b905090565b61031d610470565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102275773ffffffffffffffffffffffffffffffffffffffff81166103bc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603681526020018061058f6036913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6103e5610470565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301528051918290030190a1610222816104e2565b6000610310610470565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e80801561046b573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b61049e81610506565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b61050f81610588565b610564576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001806105c5603b913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b3b15159056fe43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f787920746f20746865207a65726f206164647265737343616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373a265627a7a723158200da8e66fafe34c061153fc15cc5ac6b9a91fb1a412f12ff9a7ebe2dd4401368e64736f6c634300050b0032")
}

// AssembleStorage builds genesis storage for the Upgradability contract
func AssembleStorage(admin common.Address, implementation common.Address, storage map[common.Hash]common.Hash) map[common.Hash]common.Hash {
	if storage == nil {
		storage = make(map[common.Hash]common.Hash)
	}
	storage[proxypos.Admin()] = admin.Hash()
	storage[proxypos.Implementation()] = implementation.Hash()
	return storage
}
