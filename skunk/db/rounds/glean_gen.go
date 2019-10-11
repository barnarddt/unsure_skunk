// Code generated by glean from glean.go:5. DO NOT EDIT.
package rounds

import (
	"context"
	"database/sql"

	"unsure_skunk/skunk"
)

const cols = " `id`, `external_id`, `player`, `rank`, `status`, `created_at`, `updated_at` "
const selectPrefix = "select " + cols + " from rounds where "

func Lookup(ctx context.Context, dbc dbc, id int64) (*skunk.Round, error) {
	return lookupWhere(ctx, dbc, "id=?", id)
}

// lookupWhere queries the rounds table with the provided where clause, then scans
// and returns a single row.
func lookupWhere(ctx context.Context, dbc dbc, where string, args ...interface{}) (*skunk.Round, error) {
	return scan(dbc.QueryRowContext(ctx, selectPrefix+where, args...))
}

// listWhere queries the rounds table with the provided where clause, then scans
// and returns all the rows.
func listWhere(ctx context.Context, dbc dbc, where string, args ...interface{}) ([]skunk.Round, error) {

	rows, err := dbc.QueryContext(ctx, selectPrefix+where, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []skunk.Round
	for rows.Next() {
		r, err := scan(rows)
		if err != nil {
			return nil, err
		}
		res = append(res, *r)
	}

	return res, rows.Err()
}

func scan(row row) (*skunk.Round, error) {
	var g glean

	err := row.Scan(&g.ID, &g.ExternalID, &g.Player, &g.Rank, &g.Status, &g.CreatedAt, &g.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &skunk.Round{
		ID:         g.ID,
		ExternalID: g.ExternalID,
		Player:     g.Player,
		Rank:       g.Rank,
		Status:     g.Status,
		CreatedAt:  g.CreatedAt,
		UpdatedAt:  g.UpdatedAt,
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