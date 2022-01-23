package main

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		// mainnet
		// "enode://<hash>@<ip>:5050",
		// "enode://<hash>@<ip>:5050",
		// "enode://<hash>@<ip>:5050",
		// "enode://<hash>@<ip>:5050",
		// "enode://<hash>@<ip>:5050",
		// testnet
		// "enode://<hash>@<ip>:5050",
		// "enode://<hash>@<ip>:5050",
		// "enode://<hash>@<ip>:5050",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
