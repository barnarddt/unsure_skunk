package ops

import (
	"context"
	"github.com/luno/fate"
	"github.com/luno/reflex"
	"unsure_skunk/skunk"
)

func makeConsume(b Backends, c skunk.Client) reflex.Consumer {
	fn := func(ctx context.Context, fate fate.Fate, e *reflex.Event) error {
		if !reflex.IsAnyType(e.Type,
			skunk.RoundStatusJoined,
			skunk.RoundStatusCollected,
			skunk.RoundStatusSubmitted,
		) {
			return fate.Tempt()
		}

		if reflex.IsType(e.Type, skunk.RoundStatusCollected) {
			if err := collectPeerParts(ctx, b, c, e); err != nil {
				return err
			}
		}

		if reflex.IsType(e.Type, skunk.RoundStatusSubmitted) {
			// find out next player
			if isLeader() {
				// submit parts
				return nil
			}
			return nil
		}

		return fate.Tempt()
	}

	return reflex.NewConsumer(reflex.ConsumerName("skunk_consumer"), fn)
}
