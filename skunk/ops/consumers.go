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
			skunk.RoundStatusJoin,
			skunk.RoundStatusJoined,
			skunk.RoundStatusCollect,
			skunk.RoundStatusCollected,
			skunk.RoundStatusSubmit,
			skunk.RoundStatusSubmitted,
			skunk.RoundStatusSuccess,
			skunk.RoundStatusFailed,
			skunk.RoundStatusSubmit,
		) {
			return fate.Tempt()
		}

		return fate.Tempt()
	}

	return reflex.NewConsumer(reflex.ConsumerName("skunk_consumer"), fn)
}
