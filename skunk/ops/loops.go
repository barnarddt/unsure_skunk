package ops

import (
	"context"
	"time"
	"unsure_skunk/skunk"
	"unsure_skunk/skunk/db/cursors"
	"unsure_skunk/skunk/db/events"

	"github.com/corverroos/unsure"
	"github.com/corverroos/unsure/engine"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/log"
	"github.com/luno/reflex"
)

func StartLoops(b Backends) {

	for _, peer := range b.GetPeers() {
		go consumePeerEvents(b, peer)
	}

	go consumeEngineEvents(b)
	go startMatchForever(b)
	go joinMatchesForever(b)
	go collectRemotePartsForever(b)
	go submitPartsForever(b)
}

func consumePeerEvents(b Backends, peer skunk.Client) {
	consumable := reflex.NewConsumable(peer.Stream,
		cursors.ToStore(b.SkunkDB().DB))
	consumer := makeConsume(b, peer)
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}

func consumeEngineEvents(b Backends) {
	consumable := reflex.NewConsumable(b.EngineClient().Stream,
		cursors.ToStore(b.SkunkDB().DB))
	consumer := makeEngineConsume(b)
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}

func startMatchForever(b Backends) {
	for {
		ctx := unsure.ContextWithFate(context.Background(), unsure.DefaultFateP())

		err := b.EngineClient().StartMatch(ctx, team, len(b.GetPeers())+1)

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

func joinMatchesForever(b Backends) {
	consumable := reflex.NewConsumable(events.ToStream(b.SkunkDB().DB),
		cursors.ToStore(b.SkunkDB().DB))
	consumer := joinMatches(b)
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}

func collectRemotePartsForever(b Backends) {
	consumable := reflex.NewConsumable(events.ToStream(b.SkunkDB().DB),
		cursors.ToStore(b.SkunkDB().DB))
	consumer := collectRemoteParts(b)
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}

func submitPartsForever(b Backends) {
	consumable := reflex.NewConsumable(events.ToStream(b.SkunkDB().DB),
		cursors.ToStore(b.SkunkDB().DB))
	consumer := submitParts(b)
	unsure.ConsumeForever(unsure.FatedContext, consumable.Consume, consumer)
}
