package goose

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
		name    string
		connStr string
		path    string
		want    string
	}{
		{"empty connStr", "", absDir, "connection string is empty"},
		{"empty path", "postgres://localhost/db", "", "migrations path is empty"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Run(ctx, tt.connStr, tt.path)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.want)
		})
	}
}
