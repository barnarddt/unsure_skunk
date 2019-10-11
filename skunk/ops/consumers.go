package ops

import (
	"context"
	"github.com/luno/fate"
	"github.com/luno/reflex"
	"unsure_skunk/skunk"
	"unsure_skunk/skunk/db/rounds"
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

		if !reflex.IsType(e.Type, skunk.RoundStatusCollected) {
			// fetch parts from e.ForeignID
			//part, err := parts.Lookup(ctx, b.SkunkDB(), e.ForeignIDInt())
			r, err := rounds.Lookup(ctx, b.SkunkDB(),e.ForeignID)


			c.GetParts(ctx,  r. int64, player string)
		}

		return fate.Tempt()
	}

	return reflex.NewConsumer(reflex.ConsumerName("skunk_consumer"), fn)
}
