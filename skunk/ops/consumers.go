package ops

import (
	"context"
	"github.com/luno/fate"
	"github.com/luno/reflex"
	"github.com/pkg/errors"
	"unsure_skunk/skunk"
	"unsure_skunk/skunk/db/parts"
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
			r, err := rounds.Lookup(ctx, b.SkunkDB().DB, e.ForeignIDInt())

			part, rank, err := c.GetData(ctx, r.ExternalID)
			if err != nil {
				return errors.Wrap(err, "failed to get data over rpc")
			}

			err = parts.Create(ctx, b.SkunkDB().DB, part[0], rank)
			if err != nil {
				return errors.Wrap(err, "failed to create part")
			}
		}

		return fate.Tempt()
	}

	return reflex.NewConsumer(reflex.ConsumerName("skunk_consumer"), fn)
}
