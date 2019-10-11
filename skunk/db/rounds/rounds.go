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
	return lookupWhere(ctx, dbc, "1=1 order by ext_id desc limit 1")
}

func ShitToJoin(ctx context.Context, dbc *sql.DB, player string, extID int64) (int64, error) {
	id, err := roundFSM.Insert(ctx, dbc, ready{Player: player, ExternalID: extID})
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert round")
	}

	return id, nil
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

func ShiftToCollect(ctx context.Context, dbc *sql.DB, id int64) error {
	r, err := Lookup(ctx, dbc, id)
	if err != nil {
		return errors.Wrap(err, "failed to lookup round",
			j.KV("round", id))
	}

	return roundFSM.Update(ctx, dbc, r.Status, skunk.RoundStatusCollect,
		empty{ID: id})
}

func ShiftToCollected(ctx context.Context, dbc *sql.DB, id int64) error {
	r, err := Lookup(ctx, dbc, id)
	if err != nil {
		return errors.Wrap(err, "failed to lookup round",
			j.KV("round", id))
	}

	return roundFSM.Update(ctx, dbc, r.Status, skunk.RoundStatusCollected,
		collected{ID: id})
}

func ShiftToSubmit(ctx context.Context, dbc *sql.DB, id int64) error {
	r, err := Lookup(ctx, dbc, id)
	if err != nil {
		return errors.Wrap(err, "failed to lookup round",
			j.KV("round", id))
	}

	return roundFSM.Update(ctx, dbc, r.Status, skunk.RoundStatusSubmitted,
		empty{ID: id})
}

func ShiftToSubmitted(ctx context.Context, dbc *sql.DB, id int64) error {
	r, err := Lookup(ctx, dbc, id)
	if err != nil {
		return errors.Wrap(err, "failed to lookup round",
			j.KV("round", id))
	}

	return roundFSM.Update(ctx, dbc, r.Status, skunk.RoundStatusSubmitted,
		empty{ID: id})
}

func LookupLatest(ctx context.Context, dbc *sql.DB, player string, round int64) (*skunk.Round, error) {
	return lookupWhere(ctx, dbc, "where ext_id=? and player=?", round, player)
}
