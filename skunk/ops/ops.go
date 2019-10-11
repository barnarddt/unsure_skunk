package ops

import (
	"context"
	"flag"
	"unsure_skunk/skunk/db/parts"

	"github.com/luno/fate"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
	"github.com/luno/reflex"

	"unsure_skunk/skunk"
	"unsure_skunk/skunk/db/rounds"
)

var player = flag.String("player", "skunky", "player name")

const (
	team = "skunkworx"
)

func GetPlayerName() string {
	return *player
}

func joinMatches(b Backends) reflex.Consumer {
	f := func(ctx context.Context, f fate.Fate, e *reflex.Event) error {
		// Skip uninteresting states.
		if !reflex.IsType(e.Type, skunk.RoundStatusJoin) {
			return fate.Tempt()
		}

		r, err := rounds.Lookup(ctx, b.SkunkDB().DB, e.ForeignIDInt())
		if err != nil {
			return errors.Wrap(err, "failed to lookup round",
				j.KV("round", e.ForeignIDInt()))
		}

		// Join the next round.
		joined, err := b.EngineClient().JoinRound(ctx, team, *player,
			r.ExternalID)
		if err != nil {
			return errors.Wrap(err, "failed to join round",
				j.KV("round", r.ExternalID))
		}

		// If joining was unsuccessful, skip the round.
		if !joined {
			return fate.Tempt()
		}

		// Join was successful, update your local state.
		err = rounds.ShiftToJoined(ctx, b.SkunkDB().ReplicaOrMaster(),
			e.ForeignIDInt())
		if err != nil {
			return errors.Wrap(err, "failed to update state to joined",
				j.KV("round", e.ForeignIDInt()))
		}

		return fate.Tempt()
	}

	return reflex.NewConsumer(skunk.ConsumerJoinRounds, f)
}

func collectParts(b Backends) reflex.Consumer {
	f := func(ctx context.Context, f fate.Fate, e *reflex.Event) error {
		// Skip uninteresting states.
		if !reflex.IsType(e.Type, skunk.RoundStatusCollect) {
			return fate.Tempt()
		}

		// Lookup the current round.
		r, err := rounds.Lookup(ctx, b.SkunkDB().DB, e.ForeignIDInt())
		if err != nil {
			return errors.Wrap(err, "failed to lookup round",
				j.KV("round", e.ForeignIDInt()))
		}

		// Attempt to collect parts from the engine.
		parts, err := b.EngineClient().CollectRound(ctx, team, *player,
			r.ExternalID)
		if err != nil {
			return errors.Wrap(err, "failed to collect parts for round",
				j.KV("round", r.ExternalID))
		}

		// TODO(Nick): Insert the parts.

		// Shift the round state to collected.
		err = rounds.ShiftToCollected(ctx, b.SkunkDB().DB, r.ID,
			int64(parts.Rank))
		if err != nil {
			return errors.Wrap(err, "failed to update state to collected",
				j.KV("round", r.ID))
		}

		return fate.Tempt()
	}

	return reflex.NewConsumer(skunk.ConsumerCollectParts, f)
}

func LookUpData(ctx context.Context, b Backends, round int64) ([]skunk.PartType, int, error) {
	part, err := parts.List(ctx, b.SkunkDB().DB, round)
	if err != nil {
		return nil, 0, err
	}

	ranker, err := rounds.LookupLatest(ctx, b.SkunkDB().DB, *player, round)
	if err != nil {
		return nil, 0, err
	}

	return part, ranker.Rank, nil
}
