package ops

import (
	"context"
	"database/sql"
	"flag"

	"github.com/luno/fate"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
	"github.com/luno/reflex"

	"unsure_skunk/skunk"
	"unsure_skunk/skunk/db/rounds"
)

var player = flag.String("player", "loser", "player name")

const (
	team = "skunkworx"
)

func joinMatches(b Backends) reflex.Consumer {
	f := func(ctx context.Context, f fate.Fate, e *reflex.Event) error {
		// Skip uninteresting states.
		if !reflex.IsType(e.Type, skunk.RoundStatusJoin) {
			return fate.Tempt()
		}

		// Default nextRoundID will be 0, i.e. The first round.
		var nextRoundID int64

		// Lookup the most recent round.
		lastMatch, err := rounds.LookupLastCompletedRound(ctx,
			b.SkunkDB().ReplicaOrMaster())
		if err == nil { // Found a completed round, increment the id.
			nextRoundID = lastMatch.ExternalID + 1
		} else if !errors.Is(err, sql.ErrNoRows){ // Unexpected error.
			return errors.Wrap(err, "failed to find last completed round")
		}

		// Join the next round.
		joined, err := b.EngineClient().JoinRound(ctx, team, *player,
			nextRoundID)
		if err != nil {
			return errors.Wrap(err, "failed to join round",
				j.KV("round", nextRoundID))
		}

		// If joining was unsuccessful, skip the round.
		if !joined {
			return fate.Tempt()
		}

		// Join was successful, update your local state.
		err = rounds.ShiftToJoined(ctx, b.SkunkDB().ReplicaOrMaster(),
			e.ForeignIDInt(), nextRoundID)
		if err != nil {
			return errors.Wrap(err, "failed to update state to joined",
				j.KV("round", e.ForeignIDInt()))
		}

		return fate.Tempt()
	}

	return reflex.NewConsumer(skunk.ConsumerJoinRounds, f)
}