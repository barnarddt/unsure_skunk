// Code generated by shiftgen at shift.go:10. DO NOT EDIT.

package rounds

import (
	"context"
	"database/sql"
	"strings"
	"time"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
	"github.com/luno/shift"
)

// Insert inserts a new rounds table entity. All the fields of the 
// ready receiver are set, as well as status, created_at and updated_at. 
// The newly created entity id is returned on success or an error.
func (一 ready) Insert(ctx context.Context, tx *sql.Tx,st shift.Status) (int64, error) {
	var (
		q    strings.Builder
		args []interface{}
	)

	q.WriteString("insert into rounds set `status`=?, `created_at`=?, `updated_at`=? ")
	args = append(args, st.Enum(), time.Now(), time.Now())

	q.WriteString(", `player`=?")
	args = append(args, 一.Player)

	q.WriteString(", `external_id`=?")
	args = append(args, 一.ExternalID)

	q.WriteString(", `submitted`=?")
	args = append(args, 一.Submitted)

	res, err := tx.ExecContext(ctx, q.String(), args...)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update updates the status of a rounds table entity. All the fields of the
// joined receiver are updated, as well as status and updated_at. 
// The entity id is returned on success or an error.
func (一 joined) Update(ctx context.Context, tx *sql.Tx,from shift.Status, 
	to shift.Status) (int64, error) {
	var (
		q    strings.Builder
		args []interface{}
	)

	q.WriteString("update rounds set `status`=?, `updated_at`=? ")
	args = append(args, to.Enum(), time.Now())

	q.WriteString(" where `id`=? and `status`=?")
	args = append(args, 一.ID, from.Enum())

	res, err := tx.ExecContext(ctx, q.String(), args...)
	if err != nil {
		return 0, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, errors.Wrap(shift.ErrRowCount, "joined", j.KV("count", n))
	}

	return 一.ID, nil
}

// Update updates the status of a rounds table entity. All the fields of the
// collected receiver are updated, as well as status and updated_at. 
// The entity id is returned on success or an error.
func (一 collected) Update(ctx context.Context, tx *sql.Tx,from shift.Status, 
	to shift.Status) (int64, error) {
	var (
		q    strings.Builder
		args []interface{}
	)

	q.WriteString("update rounds set `status`=?, `updated_at`=? ")
	args = append(args, to.Enum(), time.Now())

	q.WriteString(" where `id`=? and `status`=?")
	args = append(args, 一.ID, from.Enum())

	res, err := tx.ExecContext(ctx, q.String(), args...)
	if err != nil {
		return 0, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, errors.Wrap(shift.ErrRowCount, "collected", j.KV("count", n))
	}

	return 一.ID, nil
}

// Update updates the status of a rounds table entity. All the fields of the
// empty receiver are updated, as well as status and updated_at. 
// The entity id is returned on success or an error.
func (一 empty) Update(ctx context.Context, tx *sql.Tx,from shift.Status, 
	to shift.Status) (int64, error) {
	var (
		q    strings.Builder
		args []interface{}
	)

	q.WriteString("update rounds set `status`=?, `updated_at`=? ")
	args = append(args, to.Enum(), time.Now())

	q.WriteString(" where `id`=? and `status`=?")
	args = append(args, 一.ID, from.Enum())

	res, err := tx.ExecContext(ctx, q.String(), args...)
	if err != nil {
		return 0, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, errors.Wrap(shift.ErrRowCount, "empty", j.KV("count", n))
	}

	return 一.ID, nil
}
