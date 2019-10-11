package ops

import (
	"context"
	"github.com/luno/fate"
	"github.com/luno/reflex"
	"unsure_skunk/skunk"
)

func makeConsume(b Backends) reflex.Consumer {
	fn := func(ctx context.Context, fate fate.Fate, e *reflex.Event) error {
		if !reflex.IsAnyType(e.Type,
			skunk.RoundStatusJoined,
			skunk.RoundStatusCollected,
			skunk.RoundStatusSubmitted,
			skunk.RoundStatusSuccess,
			skunk.RoundStatusFailed,
		) {
			return fate.Tempt()
		}

		if !reflex.IsAnyType(e.Type, skunk.RoundStatusCollected) {
			// fetch parts from e.ForeignID
		}

		return fate.Tempt()
	}

	return reflex.NewConsumer(reflex.ConsumerName("skunk_consumer"), fn)
}
