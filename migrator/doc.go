// Package migrator provides migration runners in subpackages. Use the one that matches your migration layout.
//
// # goose (migrator/goose)
//
// Run(ctx, connStr, migrationsPath) runs pressly/goose "up" migrations. SQL files use +goose Up/Down directives.
// connStr and migrationsPath must be non-empty. migrationsPath is cleaned with filepath.Clean and should be under application control (not user input). ctx is used for cancellation.
//
// # migrate (migrator/migrate)
//
// Run(ctx, connURL, migrationsPath) runs golang-migrate "up" from file://migrationsPath. Expects separate .up.sql and .down.sql files. connURL and migrationsPath must be non-empty. migrationsPath is cleaned and should be under application control. ctx is checked before starting; if already cancelled, Run returns immediately. The underlying library does not accept context for Up(), so a migration in progress cannot be cancelled—long migrations may delay shutdown. ErrNoChange is ignored.
//
// # testutil (migrator/testutil)
//
// StartPostgres(t) starts a PostgreSQL container via testcontainers-go and returns a connection string. t.Cleanup terminates the container. Use in integration tests for goose and migrate.
package migrator
