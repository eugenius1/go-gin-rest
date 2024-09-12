package album

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/eugenius1/go-gin-rest/internal/models"
	"github.com/eugenius1/go-gin-rest/pkg/id"
)

func Test_repo(t *testing.T) {
	t.Parallel()

	r := NewRepo(db)

	album := models.Album{
		ID:     id.New(),
		Title:  "Test Title",
		Artist: "Test Artist",
		Price:  9.99,
	}

	// Create
	creation := time.Now()
	err := r.CreateAlbum(album)
	require.NoError(t, err)

	// Get
	got, err := r.GetAlbumByID(album.ID)
	require.NoError(t, err)
	assert.Equal(t, album.ID, got.ID)
	assert.Equal(t, album.Title, got.Title)
	assert.Equal(t, album.Artist, got.Artist)
	assert.Equal(t, album.Price, got.Price)
	assert.True(t, got.CreatedAt.After(creation), "got %s, want after %s", got.CreatedAt, creation)
	assert.True(t, got.UpdatedAt.After(creation), "got %s, want after %s", got.UpdatedAt, creation)
	assert.False(t, got.DeletedAt.Valid)
	creation = got.CreatedAt

	// List
	albums, err := r.ListAlbums()
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(albums), 1)

	updatedAlbum := models.Album{
		ID:     album.ID,
		Title:  "Updated Title",
		Artist: "Updated Artist",
		Price:  19.99,
	}

	// Update
	update := time.Now()
	err = r.UpdateAlbum(updatedAlbum)
	require.NoError(t, err)

	got, err = r.GetAlbumByID(album.ID)
	require.NoError(t, err)
	assert.Equal(t, updatedAlbum.ID, got.ID)
	assert.Equal(t, updatedAlbum.Title, got.Title)
	assert.Equal(t, updatedAlbum.Artist, got.Artist)
	assert.Equal(t, updatedAlbum.Price, got.Price)
	assert.True(t, got.CreatedAt.Equal(creation), "got %s, want %s", got.CreatedAt, creation)
	assert.True(t, got.UpdatedAt.After(update), "got %s, want after %s", got.UpdatedAt, update)
	assert.False(t, got.DeletedAt.Valid)

	// Delete
	err = r.DeleteAlbum(album.ID)
	require.NoError(t, err)

	got, err = r.GetAlbumByID(album.ID)
	require.Error(t, err)
}
