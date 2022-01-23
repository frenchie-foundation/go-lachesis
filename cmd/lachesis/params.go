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
		"enode://69e1a56cb64a90fc8774a6c5c25fcc5553155cccbc517b5cde110f0c6d14566aeded1b2829f0c23243c738cd084e8ebd96ce280dd0c0c207d613b44e1fd0d36e@ec2-18-192-119-74.eu-central-1.compute.amazonaws.com:5050",
		"enode://b632c65b1e9ec68ebecac3990c1ee1b6d768e07794318dfa9e4afba29fd0ed67093e05f9928b5f0f4fd3184dc439c2f43a8f11db5f3704dd4890213f97ba986b@ec2-18-197-95-103.eu-central-1.compute.amazonaws.com:5050",
		"enode://592054aef27ab44a3259a76190092b2517736caf825c73c06ed63b067d0350ecfde72a37e92551216f3d6f444b213bae612fbc47a89bc5fb8833899650ba3b56@ec2-18-184-83-6.eu-central-1.compute.amazonaws.com:5050",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
