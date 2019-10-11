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

		if reflex.IsType(e.Type, skunk.RoundStatusCollected) {
			if err := collectPeerParts(ctx, b, c, e); err != nil {
				return err
			}
		}

		if reflex.IsType(e.Type, skunk.RoundStatusSubmitted) {
			if err := updateSubmitState(ctx, b, c, e); err != nil {
				return err
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
			var ExternalID int64

			r, err := rounds.LookupLastCompletedRound(ctx, b.SkunkDB().DB)
			if err == nil {
				ExternalID = r.ExternalID + 1
			} else if !errors.Is(err, sql.ErrNoRows) {
				return err
			}

			_, err = rounds.ShitToJoin(ctx, b.SkunkDB().DB, GetPlayerName(), ExternalID)
			if err != nil {
				return err
			}

			return nil
		}

		return fate.Tempt()
	}

	return reflex.NewConsumer(reflex.ConsumerName("skunk_engine_consumer"), fn)
}
