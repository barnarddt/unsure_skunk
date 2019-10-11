package ops

import (
	"context"
	"flag"
	"strings"

	"github.com/luno/fate"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
	"github.com/luno/reflex"

	"unsure_skunk/skunk"
	"unsure_skunk/skunk/db/parts"
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

func skipLocalJoined(b Backends) reflex.Consumer {
	f := func(ctx context.Context, f fate.Fate, e *reflex.Event) error {
		// Skip uninteresting states.
		if !reflex.IsType(e.Type, skunk.RoundStatusJoined) {
			return fate.Tempt()
		}

		err := rounds.ShiftToCollect(ctx, b.SkunkDB().DB, e.ForeignIDInt())
		if err != nil {
			return errors.Wrap(err, "failed to update state to collect",
				j.KV("round", e.ForeignIDInt()))
		}

		return fate.Tempt()
	}

	return reflex.NewConsumer(skunk.ConsumerSkipLocalJoined, f)
}

func collectRemoteParts(b Backends) reflex.Consumer {
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
		pl, err := b.EngineClient().CollectRound(ctx, team, *player,
			r.ExternalID)
		if err != nil {
			return errors.Wrap(err, "failed to collect parts for round",
				j.KV("round", r.ExternalID))
		}

		localParts := make([]skunk.PartType, 0)
		for _, p := range pl.Players {
			localParts = append(localParts, skunk.PartType{
				RoundID: r.ExternalID,
				Player:  p.Name,
				Part:    int64(p.Part),
				Rank:    int64(pl.Rank),
			})
		}

		err = parts.CreateBatch(ctx, b.SkunkDB().DB, localParts)
		if err != nil {
			return errors.Wrap(err, "failed to insert remote parts",
				j.KV("round", r.ExternalID))
		}

		// Shift the round state to collected.
		err = rounds.ShiftToCollected(ctx, b.SkunkDB().DB, r.ID)
		if err != nil {
			return errors.Wrap(err, "failed to update state to collected",
				j.KV("round", r.ID))
		}

		return fate.Tempt()
	}

	return reflex.NewConsumer(skunk.ConsumerCollectParts, f)
}

func LookUpData(ctx context.Context, b Backends, round int64) ([]skunk.PartType, error) {
	part, err := parts.List(ctx, b.SkunkDB().DB, round)
	if err != nil {
		return nil, err
	}

	return part, nil
}

func collectPeerParts(ctx context.Context, b Backends, c skunk.Client, e *reflex.Event) error {
	r, err := rounds.Lookup(ctx, b.SkunkDB().DB, e.ForeignIDInt())
	if err != nil {
		return errors.Wrap(err, "failed round lookup")
	}

	part, err := c.GetData(ctx, r.ExternalID)
	if err != nil {
		return errors.Wrap(err, "failed to get data over rpc")
	}

	err = parts.CreateBatch(ctx, b.SkunkDB().DB, part)
	if err != nil {
		return errors.Wrap(err, "failed to create part")
	}

	return nil
}

func maybeSubmitParts(ctx context.Context, b Backends, c skunk.Client, e *reflex.Event) error {
	r, err := rounds.Lookup(ctx, b.SkunkDB().DB, e.ForeignIDInt())
	if err != nil {
		return errors.Wrap(err, "failed round lookup")
	}

	pl, err := parts.List(ctx, b.SkunkDB().DB, r.ExternalID)
	if err != nil {
		return errors.Wrap(err, "failed list all parts")
	}

	shouldSubmit := false
	for i, p := range pl {
		if i == int(r.Submitted) && strings.EqualFold(p.Player, *player) {
			shouldSubmit = true
		}
	}

	if !shouldSubmit {
		return rounds.IncrementSubmitted(ctx, b.SkunkDB().DB, r.ID)
	}

	err = rounds.ShiftToSubmit(ctx, b.SkunkDB().DB, r.ID)
	if err != nil {
		return errors.Wrap(err, "failed to update state to submit",
			j.KV("round", r.ID))
	}

	return nil
}

func submitParts(b Backends) reflex.Consumer {
	f := func(ctx context.Context, f fate.Fate, e *reflex.Event) error {
		// Skip uninteresting states.
		if !reflex.IsType(e.Type, skunk.RoundStatusSubmit) {
			return fate.Tempt()
		}

		// Lookup the current round.
		r, err := rounds.Lookup(ctx, b.SkunkDB().DB, e.ForeignIDInt())
		if err != nil {
			return errors.Wrap(err, "failed to lookup round",
				j.KV("round", e.ForeignIDInt()))
		}

		pl, err := parts.ListForPlayerAndRound(ctx, b.SkunkDB().DB, *player,
			r.ID)
		if err != nil {
			return errors.Wrap(err, "failed to lookup parts")
		}

		total := 0
		for _, p := range pl {
			total += int(p.Part)
		}

		// Submit the total of our parts to the engine.
		err = b.EngineClient().SubmitRound(ctx, team, *player, r.ExternalID,
			total)
		if err != nil {
			return errors.Wrap(err, "failed to submit part total",
				j.KV("round", r.ExternalID), j.KV("total", total))
		}

		// Shift the round to submitted.
		err = rounds.ShiftToSubmitted(ctx, b.SkunkDB().DB, r.ID)
		if err != nil {
			return errors.Wrap(err, "failed to update state to submitted")
		}

		return fate.Tempt()
	}

	return reflex.NewConsumer(skunk.ConsumerSubmitParts, f)
}
