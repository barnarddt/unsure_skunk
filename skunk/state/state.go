package state

import (
	"flag"
	"strings"
	"unsure_skunk/skunk"
	"unsure_skunk/skunk/client"
	"unsure_skunk/skunk/db"

	"github.com/corverroos/unsure/engine"
	engine_client "github.com/corverroos/unsure/engine/client"
)

var peers = flag.String("peer_addresses", "", "host:port|host:port of peer skunk gRPC service")

type State struct {
	skunkDB      *db.SkunkDB
	engineClient engine.Client
	peers        []skunk.Client
}

func (s *State) SkunkDB() *db.SkunkDB {
	return s.skunkDB
}

func (s *State) EngineClient() engine.Client {
	return s.engineClient
}

func (s *State) GetPeers() []skunk.Client {
	return s.peers
}

// New returns a new engine state.
func New() (*State, error) {
	var (
		s   State
		err error
	)

	s.skunkDB, err = db.Connect()
	if err != nil {
		return nil, err
	}

	enc, err := engine_client.New()
	if err != nil {
		return nil, err
	}
	s.engineClient = enc

	pa := strings.Split(*peers, "|")
	for _, peerAddr := range pa {
		if peerAddr == "" {
			continue
		}
		peer, err := client.New(client.WithAddress(peerAddr))
		if err != nil {
			return nil, err
		}
		s.peers = append(s.peers, peer)
	}

	return &s, nil
}
