package main

import (
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/frenchie-foundation/go-lachesis/kvdb"
)

// set RLP value
func set(table kvdb.KeyValueStore, key []byte, val interface{}) {
	buf, err := rlp.EncodeToBytes(val)
	if err != nil {
		panic(err)
	}

	if err := table.Put(key, buf); err != nil {
		panic(err)
	}
}

// get RLP value
func get(table kvdb.KeyValueStore, key []byte, to interface{}) interface{} {
	buf, err := table.Get(key)
	if err != nil {
		panic(err)
	}
	if buf == nil {
		return nil
	}

	err = rlp.DecodeBytes(buf, to)
	if err != nil {
		panic(err)
	}
	return to
}
