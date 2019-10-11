package parts

import (
	"context"
	"database/sql"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"

	"unsure_skunk/skunk"
)

var ErrURoundIDInvalid = errors.New("Invalid round ID", j.C("ERR_123"))
var ErrPlayerInvalid = errors.New("Invalid player name", j.C("ERR_456"))

func GetPart(ctx context.Context, dbc *sql.DB, roundID int64, player string) (*skunk.PartType, error) {
	if roundID < 0 {
		return nil, ErrURoundIDInvalid
	}

	if player == "" {
		return nil, ErrPlayerInvalid
	}

	return lookupWhere(ctx, dbc, "round_id=? and player=?", roundID, player)
}

func Create(ctx context.Context, tx *sql.Tx, part skunk.PartType) error {
	_, err := tx.ExecContext(ctx, "insert into parts set round_id=?, " +
		"player=?, rank=?, part=?, created_at=now()", part.RoundID,
		part.Player, part.Rank, part.Part)
	if err != nil {
		return errors.Wrap(err, "failed to insert part")
	}

	return nil
}

func CreateBatch(ctx context.Context, dbc *sql.DB, parts []skunk.PartType) error {
	tx, err := dbc.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to start transaction")
	}
	defer tx.Rollback()

	for _, p := range parts {
		err := Create(ctx, tx, p)
		if err != nil {
			return errors.Wrap(err, "failed to insert part")
		}
	}

	return tx.Commit()
}

func List(ctx context.Context, dbc *sql.DB, roundID int64) ([]skunk.PartType, error) {
	if roundID < 0 {
		return nil, ErrURoundIDInvalid
	}

	return listWhere(ctx, dbc, "round_id=?", roundID)
}
