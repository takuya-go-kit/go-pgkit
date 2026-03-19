package migrate

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRun_EmptyParams(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	dir := t.TempDir()
	absDir, err := filepath.Abs(dir)
	require.NoError(t, err)

	tests := []struct {
		name   string
		connURL string
		path  string
		want  string
	}{
		{"empty connURL", "", absDir, "connection URL is empty"},
		{"empty path", "postgres://localhost/db", "", "migrations path is empty"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Run(ctx, tt.connURL, tt.path)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.want)
		})
	}
}

func TestRun_CancelledContext(t *testing.T) {
	t.Parallel()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	dir := t.TempDir()
	absDir, err := filepath.Abs(dir)
	require.NoError(t, err)

	err = Run(ctx, "postgres://user:pass@localhost:5432/db?sslmode=disable", absDir)
	require.Error(t, err)
	assert.ErrorIs(t, err, context.Canceled)
}
