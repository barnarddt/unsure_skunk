package rounds

import (
	"context"
	"database/sql"

	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"

	"unsure_skunk/skunk"
)

func LookupLastCompletedRound(ctx context.Context, dbc *sql.DB) (*skunk.Round,
	error) {
	return lookupWhere(ctx, dbc, "order by ext_id desc limit 1")
}

func ShiftToJoined(ctx context.Context, dbc *sql.DB, id int64) error {
	r, err := Lookup(ctx, dbc, id)
	if err != nil {
		return errors.Wrap(err, "failed to lookup round",
			j.KV("round", id))
	}

	return roundFSM.Update(ctx, dbc, r.Status, skunk.RoundStatusJoined,
		joined{ID: id})
}

func ShiftToCollected(ctx context.Context, dbc *sql.DB, id, rank int64) error {
	r, err := Lookup(ctx, dbc, id)
	if err != nil {
		return errors.Wrap(err, "failed to lookup round",
			j.KV("round", id))
	}

	return roundFSM.Update(ctx, dbc, r.Status, skunk.RoundStatusCollected,
		collected{ID: id, Rank: rank})
}

func LookupLatest(ctx context.Context, dbc *sql.DB, player string, round int64) (*skunk.Round, error) {
	return lookupWhere(ctx, dbc, "where ext_id=? and player=?", round, player)
}
