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

func ShitToJoin(ctx context.Context, dbc *sql.DB, player string) (int64, error) {
	id, err := roundFSM.Insert(ctx, dbc, ready{Player: player})
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert round")
	}

	return id, nil
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