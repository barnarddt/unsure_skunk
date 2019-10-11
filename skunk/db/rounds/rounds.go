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


func ShiftToJoined(ctx context.Context, dbc *sql.DB, roundID, extID int64) error {
	r, err := Lookup(ctx, dbc, roundID)
	if err != nil {
		return errors.Wrap(err, "failed to lookup round",
			j.KV("round", roundID))
	}

	err = roundFSM.Update(ctx, dbc, r.Status, skunk.RoundStatusJoined,
		joined{ID: roundID, ExternalID: extID})
	if err != nil {
		return errors.Wrap(err, "failed to shift to joined")
	}

	return nil
}