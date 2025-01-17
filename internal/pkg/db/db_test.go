//go:build unit

package db

import (
	"context"
	"testing"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/config"

	"github.com/stretchr/testify/require"
)

// TestLink ...
func TestLink(t *testing.T) {
	ctx := context.Background()

	// Init logger
	conf := config.Configuration{}
	log, err := logger.NewLogger(logger.Zap, conf)
	require.NoError(t, err, "Error init a logger")

	// Init db
	s, err := New(ctx, log, nil, nil)
	require.NoError(t, err, "Error init a db")

	// Init db
	require.NoError(t, s.Init(ctx), "Error  create a new DB connection")
}
