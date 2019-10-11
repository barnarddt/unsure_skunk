package ops

import (
	"context"
	"google.golang.org/api/admin/directory/v1"
	"strconv"

	"github.com/corverroos/unsure/engine"
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

		if reflex.IsType(e.Type, engine.EventTypeRoundJoin) {
			roundID, err := strconv.ParseInt(e.ForeignID, 10, 64)
			if err != nil {
				return err
			}

			_, err = rounds.ShitToJoin(ctx, b.SkunkDB().DB, GetPlayerName(), roundID)
			if err != nil {
				return err
			}

			return nil
		}

		if reflex.IsType(e.Type, engine.EventTypeRoundSubmit) {
			err := updateSubmitState(ctx, b, e)
			if err != nil {
				return err
			}
		}

		return fate.Tempt()
	}

	return reflex.NewConsumer(reflex.ConsumerName("skunk_engine_consumer"), fn)
}
