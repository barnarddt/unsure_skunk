// Code generated by glean from glean.go:5. DO NOT EDIT.
package parts

import (
	"context"
	"database/sql"

	"unsure_skunk/skunk"
)

const cols = " `id`, `round_id`, `player`, `part`, `created_at` "
const selectPrefix = "select " + cols + " from parts where "

func Lookup(ctx context.Context, dbc dbc, id int64) (*skunk.PartType, error) {
	return lookupWhere(ctx, dbc, "id=?", id)
}

// lookupWhere queries the parts table with the provided where clause, then scans
// and returns a single row.
func lookupWhere(ctx context.Context, dbc dbc, where string, args ...interface{}) (*skunk.PartType, error) {
	return scan(dbc.QueryRowContext(ctx, selectPrefix+where, args...))
}

// listWhere queries the parts table with the provided where clause, then scans
// and returns all the rows.
func listWhere(ctx context.Context, dbc dbc, where string, args ...interface{}) ([]skunk.PartType, error) {

	rows, err := dbc.QueryContext(ctx, selectPrefix+where, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []skunk.PartType
	for rows.Next() {
		r, err := scan(rows)
		if err != nil {
			return nil, err
		}
		res = append(res, *r)
	}

	return res, rows.Err()
}

func scan(row row) (*skunk.PartType, error) {
	var g glean

	err := row.Scan(&g.ID, &g.RoundID, &g.Player, &g.Part, &g.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &skunk.PartType{
		ID:        g.ID,
		RoundID:   g.RoundID,
		Player:    g.Player,
		Part:      g.Part,
		CreatedAt: g.CreatedAt,
	}, nil
}

// dbc is a common interface for *sql.DB and *sql.Tx.
type dbc interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// row is a common interface for *sql.Rows and *sql.Row.
type row interface {
	Scan(dest ...interface{}) error
}
