// Package postgres provides pgxpool creation with configurable connection limits and retry.
//
// # Configuration
//
// Config holds the connection URL and optional pool limits. URL is required (e.g. postgres://user:pass@host:5432/db?sslmode=disable).
// MaxConns and MinConns default to 10 and 0 when zero (so when only MinConns is set, MaxConns is still 10); MaxConns must be in 1..10000 when set, MinConns in 0..10000, and MinConns <= MaxConns.
// RetryTimeout overrides the default 30s window for connection retry. Do not log Config as-is; String and GoString mask the password in the URL. MaskURL(s) is a standalone helper for safe logging of any connection string. Error messages from New also mask the URL when reporting parse failures.
//
// # Creating a pool
//
// New creates a pgxpool.Pool with exponential backoff until the database is reachable (or ctx is cancelled). Pool settings (MaxConnLifetime, MaxConnIdleTime, HealthCheckPeriod) use sensible defaults. Use the returned pool for queries and call Close when shutting down.
package postgres
