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
	return lookupWhere(ctx, dbc, "order by external_id desc limit 1")
}

func ShiftToJoined(ctx context.Context, dbc *sql.DB, id int64) error {
	r, err := Lookup(ctx, dbc, id)
	if err != nil {
		return errors.Wrap(err, "failed to lookup round",
			j.KV("round", id))
	}

	err = roundFSM.Update(ctx, dbc, r.Status, skunk.RoundStatusJoined,
		joined{ID: id})
	if err != nil {
		return errors.Wrap(err, "failed to shift to joined")
	}

	return nil
}