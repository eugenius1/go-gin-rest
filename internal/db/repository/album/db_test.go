package album

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/eugenius1/go-gin-rest/internal/db"
	"github.com/eugenius1/go-gin-rest/internal/models"
	"github.com/eugenius1/go-gin-rest/pkg/id"
)

func Test_repo(t *testing.T) {
	t.Parallel()

	db, err := db.SetupFromEnv("test_album")
	require.NoError(t, err)

	r := NewRepo(db)

	album := models.Album{
		ID:     id.New(),
		Title:  "Test Title",
		Artist: "Test Artist",
		Price:  9.99,
	}

	creation := time.Now()
	err = r.CreateAlbum(album)
	require.NoError(t, err)

	got, err := r.GetAlbumByID(album.ID)
	require.NoError(t, err)
	assert.Equal(t, album.ID, got.ID)
	assert.Equal(t, album.Title, got.Title)
	assert.Equal(t, album.Artist, got.Artist)
	assert.Equal(t, album.Price, got.Price)
	assert.True(t, got.CreatedAt.After(creation))
	assert.True(t, got.UpdatedAt.After(creation))
	assert.False(t, got.DeletedAt.Valid)

	albums, err := r.ListAlbums()
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(albums), 1)
}
