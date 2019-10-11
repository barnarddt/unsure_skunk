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
