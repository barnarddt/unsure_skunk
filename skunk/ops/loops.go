package ops

import (
	"context"
	"time"

	"github.com/corverroos/unsure"
	"github.com/corverroos/unsure/engine"
	"github.com/luno/fate"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/log"
	"github.com/luno/reflex"

	"unsure_skunk/skunk"
	"unsure_skunk/skunk/db/cursors"
)

func StartLoops(b Backends) {
	for _, peer := range b.GetPeers() {
		go consumePeerEvents(b, peer)
	}

	go startMatchForever(b)
}

func consumePeerEvents(b Backends, peer skunk.Client) {
	consumable := reflex.NewConsumable(peer.Stream,
		cursors.ToStore(b.SkunkDB().DB))
	consumer := makeConsume(b)
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}

func startMatchForever(b Backends) {
	for {
		ctx := unsure.ContextWithFate(context.Background(), unsure.DefaultFateP())

		err := b.EngineClient().StartMatch(ctx, team, len(b.GetPeers()))

		if errors.Is(err, engine.ErrActiveMatch) {
			// Match active, just ignore
			return
		} else if err != nil {
			log.Error(ctx, errors.Wrap(err, "start match error"))
		} else {
			log.Info(ctx, "match started")
			return
		}

		time.Sleep(time.Second)
	}
}

func Dummy(backends Backends) reflex.Consumer {
	fn := func(ctx context.Context, fate fate.Fate, e *reflex.Event) error {

		return fate.Tempt()
	}

	return reflex.NewConsumer(reflex.ConsumerName("generic"), fn)
}
