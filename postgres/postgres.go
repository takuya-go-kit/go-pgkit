package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	defaultMaxConns        = 10
	defaultMinConns        = 0
	maxConnsLimit          = 10000
	defaultMaxConnLifetime = time.Hour
	defaultMaxConnIdleTime = 30 * time.Minute
	defaultHealthCheck     = 15 * time.Second
	defaultConnectTimeout  = 5 * time.Second
	pgConnRetryTimeout     = 30 * time.Second
)

// New creates a pgxpool with exponential backoff until the database is reachable. ctx can cancel the retry. Returns an error if ctx or cfg is nil, the URL is invalid, or pool limits are out of range.
func New(ctx context.Context, cfg *Config) (*pgxpool.Pool, error) {
	if ctx == nil {
		return nil, fmt.Errorf("postgres - New: context is nil")
	}
	if cfg == nil {
		return nil, fmt.Errorf("postgres - New: config is nil")
	}
	poolCfg, err := pgxpool.ParseConfig(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("postgres - New: invalid DSN format: %w", err)
	}
	maxConns := defaultMaxConns
	if cfg.MaxConns > 0 {
		maxConns = cfg.MaxConns
	}
	minConns := defaultMinConns
	if cfg.MinConns > 0 {
		minConns = cfg.MinConns
	}
	if cfg.MaxConns != 0 && (cfg.MaxConns < 1 || cfg.MaxConns > maxConnsLimit) {
		return nil, fmt.Errorf("postgres - New: MaxConns must be 0 (default) or 1..%d", maxConnsLimit)
	}
	if cfg.MinConns < 0 || cfg.MinConns > maxConnsLimit {
		return nil, fmt.Errorf("postgres - New: MinConns must be 0..%d", maxConnsLimit)
	}
	if minConns > maxConns {
		return nil, fmt.Errorf("postgres - New: MinConns (%d) must be <= MaxConns (%d)", minConns, maxConns)
	}
	poolCfg.MaxConns = int32(maxConns)
	poolCfg.MinConns = int32(minConns)
	poolCfg.MaxConnLifetime = defaultMaxConnLifetime
	poolCfg.MaxConnIdleTime = defaultMaxConnIdleTime
	poolCfg.HealthCheckPeriod = defaultHealthCheck
	poolCfg.ConnConfig.ConnectTimeout = defaultConnectTimeout

	var pool *pgxpool.Pool
	operation := func() error {
		var createErr error
		pool, createErr = pgxpool.NewWithConfig(ctx, poolCfg)
		if createErr != nil {
			return createErr
		}
		if pingErr := pool.Ping(ctx); pingErr != nil {
			pool.Close()
			pool = nil
			return pingErr
		}
		return nil
	}

	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = pgConnRetryTimeout
	if cfg.RetryTimeout > 0 {
		bo.MaxElapsedTime = cfg.RetryTimeout
	}
	if err := backoff.Retry(operation, backoff.WithContext(bo, ctx)); err != nil {
		return nil, fmt.Errorf("postgres - New: connection failed after retries: %w", err)
	}

	return pool, nil
}
