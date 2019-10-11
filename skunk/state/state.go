package state

import (
	"flag"
	"strings"

	"github.com/corverroos/unsure/skunk/client"

	"github.com/corverroos/unsure/engine"
	engine_client "github.com/corverroos/unsure/engine/client"
	"github.com/corverroos/unsure/skunk"
	"github.com/corverroos/unsure/skunk/db"
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
		peer, err := client.New(client.WithAddress(peerAddr))
		if err != nil {
			return nil, err
		}
		s.peers = append(s.peers, peer)
	}

	return &s, nil
}
