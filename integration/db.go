package integration

import (
	"github.com/frenchie-foundation/go-lachesis/kvdb"
	"github.com/frenchie-foundation/go-lachesis/kvdb/leveldb"
	"github.com/frenchie-foundation/go-lachesis/kvdb/memorydb"
)

func DBProducer(dbdir string) kvdb.DbProducer {
	if dbdir == "inmemory" || dbdir == "" {
		return memorydb.NewProducer("")
	}

	return leveldb.NewProducer(dbdir)
}
