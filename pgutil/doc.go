// Package pgutil provides helpers for working with pgx (jackc/pgx).
//
// # Error checks
//
// IsNoRows reports whether err is or wraps pgx.ErrNoRows (use after QueryRow when a missing row is expected).
// IsPgUniqueViolation reports whether err is a PostgreSQL unique constraint violation (SQLSTATE 23505).
//
// # Timestamp conversion
//
// pgx uses pgtype.Timestamptz for timestamp with time zone. TimestamptzToTime returns *time.Time or nil when invalid.
// TimestamptzToTimeZero returns time.Time or the zero value. TimeToTimestamptz converts *time.Time to pgtype.Timestamptz (invalid if nil).
// PtrTimeToTime dereferences a *time.Time or returns time.Time{} if nil.
package pgutil
