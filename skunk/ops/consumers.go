package ops

import (
	"context"
	"database/sql"

	"github.com/corverroos/unsure/engine"
	"github.com/luno/fate"
	"github.com/luno/jettison/errors"
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
			r, err := rounds.Lookup(ctx, b.SkunkDB().DB, e.ForeignIDInt())

			part, rank, err := c.GetData(ctx, r.ExternalID)
			if err != nil {
				return errors.Wrap(err, "failed to get data over rpc")
			}

		}

		return fate.Tempt()
	}

	return reflex.NewConsumer(reflex.ConsumerName("skunk_consumer"), fn)
}

func makeEngineConsume(b Backends) reflex.Consumer {
	fn := func(ctx context.Context, fate fate.Fate, e *reflex.Event) error {

		if engine.EventType(e.Type.ReflexType()) == engine.EventTypeMatchStarted {
			// Check if a previous round exists
			var extID int64

			r, err := rounds.LookupLastCompletedRound(ctx, b.SkunkDB().DB)
			if err == nil {
				extID = r.ExternalID + 1
			} else if !errors.Is(err, sql.ErrNoRows) {
				return err
			}

			_, err = rounds.ShitToJoin(ctx, b.SkunkDB().DB, GetPlayerName(), extID)
			if err != nil {
				return err
			}

			return nil
		}

		return fate.Tempt()
	}

	return reflex.NewConsumer(reflex.ConsumerName("skunk_engine_consumer"), fn)
}
