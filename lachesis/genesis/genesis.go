package genesis

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/frenchie-foundation/go-lachesis/inter"
	"github.com/frenchie-foundation/go-lachesis/inter/pos"
	"github.com/frenchie-foundation/go-lachesis/lachesis/genesis/proxy"
	"github.com/frenchie-foundation/go-lachesis/lachesis/genesis/sfc"
	"github.com/frenchie-foundation/go-lachesis/utils"
)

var (
	genesisTime = inter.Timestamp(1577419000 * time.Second)
)

type Genesis struct {
	Alloc     VAccounts
	Time      inter.Timestamp
	ExtraData []byte
}

func preDeploySfc(g Genesis, implCode []byte) Genesis {
	// pre deploy SFC impl
	g.Alloc.Accounts[sfc.ContractAddressV1] = Account{
		Code:    implCode, // impl account has only code, balance and storage is in proxy account
		Balance: big.NewInt(0),
	}
	// pre deploy SFC proxy
	storage := sfc.AssembleStorage(g.Alloc.Validators, g.Time, g.Alloc.SfcContractAdmin, nil)
	storage = proxy.AssembleStorage(g.Alloc.SfcContractAdmin, sfc.ContractAddressV1, storage) // Add storage of proxy
	g.Alloc.Accounts[sfc.ContractAddress] = Account{
		Code:    proxy.GetContractBin(),
		Storage: storage,
		Balance: g.Alloc.Validators.TotalStake(),
	}
	return g
}

// FakeGenesis generates fake genesis with n-nodes.
func FakeGenesis(accs VAccounts) Genesis {
	g := Genesis{
		Alloc: accs,
		Time:  genesisTime,
	}
	g = preDeploySfc(g, sfc.GetTestContractBinV1())
	return g
}

// MainGenesis returns builtin genesis keys of mainnet.
func MainGenesis() Genesis {
	g := Genesis{
		Time: genesisTime,
		Alloc: VAccounts{
			Accounts: Accounts{
				common.HexToAddress("0x48e65A90F3FD40c4212872F9189aF7a3EF3f6e95"): Account{Balance: utils.ToFren(2000000100)},
				common.HexToAddress("0x4DD6920836e42Ef99a70D94e0C4487402881FA81"): Account{Balance: utils.ToFren(100)},
				common.HexToAddress("0xB201c0DA6556e1b3f849Ac917Bfd4AAbA2b763F6"): Account{Balance: utils.ToFren(100)},
				common.HexToAddress("0xCA5D0667518bf2e6291D3E5C70ef34b0dc206C19"): Account{Balance: utils.ToFren(100)},
				// it is recommended to use 13 accounts
			},
			Validators: pos.GValidators{
				pos.GenesisValidator{
					ID:      1,
					Address: common.HexToAddress("0x4DD6920836e42Ef99a70D94e0C4487402881FA81"),
					Stake:   utils.ToFren(10000000),
				},
				pos.GenesisValidator{
					ID:      2,
					Address: common.HexToAddress("0xB201c0DA6556e1b3f849Ac917Bfd4AAbA2b763F6"),
					Stake:   utils.ToFren(10000000),
				},
				pos.GenesisValidator{
					ID:      3,
					Address: common.HexToAddress("0xCA5D0667518bf2e6291D3E5C70ef34b0dc206C19"),
					Stake:   utils.ToFren(10000000),
				},
				// it is recommended to use 12 validators
			},
			SfcContractAdmin: common.HexToAddress("0x48e65A90F3FD40c4212872F9189aF7a3EF3f6e95"),
		},
	}
	g = preDeploySfc(g, sfc.GetMainContractBinV1())
	return g
}

// TestGenesis returns builtin genesis keys of testnet.
func TestGenesis() Genesis {
	g := Genesis{
		Time: genesisTime,
		Alloc: VAccounts{
			Accounts: Accounts{
				common.HexToAddress("0x48e65A90F3FD40c4212872F9189aF7a3EF3f6e95"): Account{Balance: utils.ToFren(2000000100)},
				common.HexToAddress("0x4DD6920836e42Ef99a70D94e0C4487402881FA81"): Account{Balance: utils.ToFren(400)},
				common.HexToAddress("0xB201c0DA6556e1b3f849Ac917Bfd4AAbA2b763F6"): Account{Balance: utils.ToFren(400)},
				common.HexToAddress("0xCA5D0667518bf2e6291D3E5C70ef34b0dc206C19"): Account{Balance: utils.ToFren(400)},
			},
			Validators: pos.GValidators{
				pos.GenesisValidator{
					ID:      1,
					Address: common.HexToAddress("0x4DD6920836e42Ef99a70D94e0C4487402881FA81"),
					Stake:   utils.ToFren(40000000),
				},
				pos.GenesisValidator{
					ID:      2,
					Address: common.HexToAddress("0xB201c0DA6556e1b3f849Ac917Bfd4AAbA2b763F6"),
					Stake:   utils.ToFren(40000000),
				},
				pos.GenesisValidator{
					ID:      3,
					Address: common.HexToAddress("0xCA5D0667518bf2e6291D3E5C70ef34b0dc206C19"),
					Stake:   utils.ToFren(40000000),
				},
			},
			SfcContractAdmin: common.HexToAddress("0x48e65A90F3FD40c4212872F9189aF7a3EF3f6e95"),
		},
	}
	g = preDeploySfc(g, sfc.GetTestContractBinV1())
	return g
}
