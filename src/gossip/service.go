package gossip

import (
	"fmt"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discv5"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/rpc"
	"math/rand"
	"sync"

	"github.com/Fantom-foundation/go-lachesis/src/crypto"
	"github.com/Fantom-foundation/go-lachesis/src/cryptoaddr"
	"github.com/Fantom-foundation/go-lachesis/src/hash"
	"github.com/Fantom-foundation/go-lachesis/src/inter"
	"github.com/Fantom-foundation/go-lachesis/src/lachesis"
	"github.com/Fantom-foundation/go-lachesis/src/logger"
)

// Service implements go-ethereum/node.Service interface.
type Service struct {
	config *lachesis.Net

	wg   sync.WaitGroup
	done chan struct{}

	// server
	Name   string
	Topics []discv5.Topic

	serverPool *serverPool

	// my identity
	me         hash.Peer
	privateKey *crypto.PrivateKey

	// application
	store    *Store
	engine   Consensus
	engineMu *sync.RWMutex
	emitter  *Emitter

	mux *event.TypeMux

	// application protocol
	pm *ProtocolManager

	logger.Instance
}

func NewService(config *lachesis.Net, mux *event.TypeMux, store *Store, engine Consensus) (*Service, error) {
	engine = &StoreAwareEngine{
		engine: engine,
		store:  store,
	}

	svc := &Service{
		config: config,

		done: make(chan struct{}),

		Name: fmt.Sprintf("Node-%d", rand.Int()),

		store:  store,
		engine: engine,

		engineMu: new(sync.RWMutex),

		mux: mux,

		Instance: logger.MakeInstance(),
	}

	engine.Bootstrap(svc.ApplyBlock)

	trustedNodes := []string{}

	svc.serverPool = newServerPool(store.table.Peers, svc.done, &svc.wg, trustedNodes)

	var err error
	svc.pm, err = NewProtocolManager(config, downloader.FullSync, config.Genesis.NetworkId, svc.mux, &dummyTxPool{}, svc.engineMu, store, engine)

	return svc, err
}

func (s *Service) makeEmitter() *Emitter {
	return NewEmitter(s.config, s.me, s.privateKey, s.engineMu, s.store, s.engine, func(emitted *inter.Event) {
		// svc.engineMu is locked here

		err := s.engine.ProcessEvent(emitted)
		if err != nil {
			s.Fatalf("Self-event connection failed: %s", err.Error())
		}

		err = s.pm.mux.Post(emitted) // PM listens and will broadcast it
		if err != nil {
			s.Fatalf("Failed to post self-event: %s", err.Error())
		}
	})
}

// Protocols returns protocols the service can communicate on.
func (s *Service) Protocols() []p2p.Protocol {
	protos := make([]p2p.Protocol, len(ProtocolVersions))
	for i, vsn := range ProtocolVersions {
		protos[i] = s.pm.makeProtocol(vsn)
		protos[i].Attributes = []enr.Entry{s.currentEnr()}
	}
	return protos
}

// APIs returns api methods the service wants to expose on rpc channels.
func (s *Service) APIs() []rpc.API {
	return []rpc.API{}
}

// Start method invoked when the node is ready to start the service.
func (s *Service) Start(srv *p2p.Server) error {

	var genesis hash.Hash
	genesis = s.engine.GetGenesisHash()
	s.Topics = []discv5.Topic{
		discv5.Topic("lachesis@" + genesis.Hex()),
	}

	if srv.DiscV5 != nil {
		for _, topic := range s.Topics {
			topic := topic
			go func() {
				s.Info("Starting topic registration")
				defer s.Info("Terminated topic registration")

				srv.DiscV5.RegisterTopic(topic, s.done)
			}()
		}
	}
	s.privateKey = (*crypto.PrivateKey)(srv.PrivateKey)
	s.me = cryptoaddr.AddressOf(s.privateKey.Public())

	s.pm.Start(srv.MaxPeers)

	s.emitter = s.makeEmitter()

	return nil
}

// Stop method invoked when the node terminates the service.
func (s *Service) Stop() error {
	fmt.Println("Service stopping...")
	s.pm.Stop()
	s.wg.Wait()
	return nil
}
