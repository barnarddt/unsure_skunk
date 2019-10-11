package ops

import (
	"github.com/corverroos/unsure/engine"
	"unsure_skunk/skunk"
	"unsure_skunk/skunk/db"
)

//go:generate genbackendsimpl
type Backends interface {
	GetPeers() []skunk.Client
	EngineClient() engine.Client
	SkunkDB() *db.SkunkDB
}
