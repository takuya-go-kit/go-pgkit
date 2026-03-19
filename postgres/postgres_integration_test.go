//go:build integration

package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/TakuyaYagam1/go-pgkit/migrator/testutil"
)

func TestNew(t *testing.T) {
	connStr := testutil.StartPostgres(t)

	pool, err := New(context.Background(), &Config{URL: connStr})
	require.NoError(t, err)
	defer pool.Close()
	require.NoError(t, pool.Ping(context.Background()))
}

func TestNew_InvalidURL(t *testing.T) {
	_, err := New(context.Background(), &Config{
		URL:          "postgres://invalid:5432/nonexistent?sslmode=disable",
		RetryTimeout: 2 * time.Second,
	})
	require.Error(t, err)
}
