//go:build unit || (database && badger)

package badger

import (
	"context"
	"testing"

	"github.com/dgraph-io/badger/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	db "github.com/shortlink-org/shortlink/internal/pkg/db/badger"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/mock"
)

// func TestMain(m *testing.M) {
//	goleak.VerifyTestMain(m)
// }

func TestBadger(t *testing.T) {
	ctx := context.Background()

	st := db.Store{}
	st.Init(ctx)

	store := Store{
		client: st.GetConn().(*badger.DB),
	}

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(ctx, mock.AddLink)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(ctx, mock.GetLink.Hash)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(ctx, nil)
		require.NoError(t, err)
		assert.Equal(t, len(links.Link), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		require.NoError(t, store.Delete(ctx, mock.GetLink.Hash))
	})
}
