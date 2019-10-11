package ops

import (
	"github.com/corverroos/unsure"
	"github.com/luno/reflex"

	"unsure_skunk/skunk"
	"unsure_skunk/skunk/db/cursors"
)

func StartLoops(b Backends) {
	for _, peer := range b.GetPeers() {
		go consumePeerEvents(b, peer)
	}
}

func consumePeerEvents(b Backends, peer skunk.Client) {
	consumable := reflex.NewConsumable(peer.Stream,
		cursors.ToStore(b.SkunkDB().DB))
	consumer := makeConsume(b)
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}
