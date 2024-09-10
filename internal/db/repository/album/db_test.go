package album

import (
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/eugenius1/go-gin-rest/internal/db"
	"github.com/eugenius1/go-gin-rest/internal/models"
)

func Test_repo(t *testing.T) {
	t.Parallel()

	db, err := db.SetupFromEnv("test_albums")
	require.NoError(t, err)

	s := NewRepo(db)

	album := models.Album{
		ID:     ulid.Make().String(),
		Title:  "Test Title",
		Artist: "Test Artist",
		Price:  9.99,
	}

	creation := time.Now()
	err = s.CreateAlbum(album)
	require.NoError(t, err)

	got, err := s.GetAlbumByID(album.ID)
	require.NoError(t, err)
	assert.Equal(t, album.ID, got.ID)
	assert.Equal(t, album.Title, got.Title)
	assert.Equal(t, album.Artist, got.Artist)
	assert.Equal(t, album.Price, got.Price)
	assert.True(t, got.CreatedAt.After(creation))
	assert.True(t, got.UpdatedAt.After(creation))
	assert.False(t, got.DeletedAt.Valid)

	albums, err := s.ListAlbums()
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(albums), 1)
}
