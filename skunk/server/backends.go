package server

import (
	"github.com/corverroos/unsure/engine"
	"unsure_skunk/skunk"
	"unsure_skunk/skunk/db"
)

//go:generate genbackendsimpl
type Backends interface {
	EngineClient() engine.Client
	SkunkDB() *db.SkunkDB
	GetPeers() []skunk.Client
}
