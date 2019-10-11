package ops

import (
	"context"

	"github.com/corverroos/unsure"
	"github.com/luno/fate"
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
	consumer := Dummy
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}

func Dummy(backends Backends) reflex.Consumer {
	fn := func(ctx context.Context, fate fate.Fate, e *reflex.Event) error {

		return fate.Tempt()
	}

	return reflex.NewConsumer(reflex.ConsumerName("generic"), fn)
}
