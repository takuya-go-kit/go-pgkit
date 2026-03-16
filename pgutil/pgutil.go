package pgutil

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

// IsNoRows reports whether err is or wraps pgx.ErrNoRows. Use after QueryRow when a missing row is acceptable.
func IsNoRows(err error) bool {
	return err != nil && errors.Is(err, pgx.ErrNoRows)
}

// IsPgUniqueViolation reports whether err is a PostgreSQL unique constraint violation (SQLSTATE 23505).
func IsPgUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return err != nil && errors.As(err, &pgErr) && pgErr.Code == "23505"
}

// TimestamptzToTime returns a pointer to the time, or nil if t.Valid is false.
func TimestamptzToTime(t pgtype.Timestamptz) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

// TimestamptzToTimeZero returns t.Time, or time.Time{} if t.Valid is false.
func TimestamptzToTimeZero(t pgtype.Timestamptz) time.Time {
	if !t.Valid {
		return time.Time{}
	}
	return t.Time
}

// TimeToTimestamptz converts *time.Time to pgtype.Timestamptz. Returns an invalid Timestamptz if t is nil.
func TimeToTimestamptz(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{}
	}
	return pgtype.Timestamptz{Time: *t, Valid: true}
}

// PtrTimeToTime returns *t, or time.Time{} if t is nil.
func PtrTimeToTime(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}
