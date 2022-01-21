package gossip

import (
	"time"

	"github.com/frenchie-foundation/go-lachesis/app"
	"github.com/frenchie-foundation/go-lachesis/kvdb"
	"github.com/frenchie-foundation/go-lachesis/kvdb/flushable"
	"github.com/frenchie-foundation/go-lachesis/kvdb/leveldb"
	"github.com/frenchie-foundation/go-lachesis/kvdb/memorydb"
)

func cachedStore() *Store {
	mems := memorydb.NewProducer("", withDelay)
	dbs := flushable.NewSyncedPool(mems)
	cfg := LiteStoreConfig()

	return NewStore(dbs, cfg, app.LiteStoreConfig())
}

func nonCachedStore() *Store {
	mems := memorydb.NewProducer("", withDelay)
	dbs := flushable.NewSyncedPool(mems)
	cfg := StoreConfig{}

	return NewStore(dbs, cfg, app.LiteStoreConfig())
}

func realStore(dir string) *Store {
	disk := leveldb.NewProducer(dir)
	dbs := flushable.NewSyncedPool(disk)
	cfg := LiteStoreConfig()

	return NewStore(dbs, cfg, app.LiteStoreConfig())
}

func withDelay(db kvdb.KeyValueStore) kvdb.KeyValueStore {
	mem, ok := db.(*memorydb.Database)
	if ok {
		mem.SetDelay(time.Millisecond)

	}

	return db
}
