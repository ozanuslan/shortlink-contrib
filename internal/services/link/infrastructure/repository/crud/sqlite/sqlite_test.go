//go:build unit || (database && sqlite)

package sqlite

import (
	"context"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/db/sqlite"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/mock"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m,
		goleak.IgnoreTopFunction("github.com/golang/glog.(*fileSink).flushDaemon"),
		goleak.IgnoreTopFunction("database/sql.(*DB).connectionOpener"))
}

func TestSQLite(t *testing.T) {
	ctx := context.Background()

	err := os.Setenv("STORE_SQLITE_PATH", "/tmp/links-test.sqlite")
	require.NoError(t, err, "Cannot set ENV")

	// Create store
	st := &sqlite.Store{}
	err = st.Init(ctx)
	require.NoError(t, err)

	// Create repository
	store, err := New(ctx, &db.Store{
		Store: st,
	})

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(ctx, mock.AddLink)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(ctx, mock.GetLink.Hash)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(ctx, nil)
		require.NoError(t, err)
		assert.Equal(t, len(links.Link), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		require.NoError(t, store.Delete(ctx, mock.GetLink.Hash))
	})

	t.Run("Close", func(t *testing.T) {
		errDeleteFile := os.Remove(viper.GetString("STORE_SQLITE_PATH"))
		require.NoError(t, errDeleteFile)
	})

	t.Cleanup(func() {
		require.NoError(t, st.Close())
	})
}
