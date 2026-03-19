package postgres

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew_Validation(t *testing.T) {
	t.Parallel()
	validURL := "postgres://user:pass@localhost:5432/dbname?sslmode=disable"

	tests := []struct {
		name string
		ctx  context.Context
		cfg  *Config
		want string
	}{
		{
			name: "nil context",
			ctx:  nil,
			cfg:  &Config{URL: validURL},
			want: "context is nil",
		},
		{
			name: "nil config",
			ctx:  context.Background(),
			cfg:  nil,
			want: "config is nil",
		},
		{
			name: "invalid DSN format",
			ctx:  context.Background(),
			cfg:  &Config{URL: "invalid://bad"},
			want: "invalid DSN format",
		},
		{
			name: "MaxConns negative",
			ctx:  context.Background(),
			cfg:  &Config{URL: validURL, MaxConns: -1},
			want: "MaxConns must be 0 (default) or 1..10000",
		},
		{
			name: "MinConns negative",
			ctx:  context.Background(),
			cfg:  &Config{URL: validURL, MinConns: -1},
			want: "MinConns must be 0..10000",
		},
		{
			name: "MinConns greater than MaxConns",
			ctx:  context.Background(),
			cfg:  &Config{URL: validURL, MaxConns: 1, MinConns: 5},
			want: "MinConns (5) must be <= MaxConns (1)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.ctx, tt.cfg)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.want)
		})
	}
}
