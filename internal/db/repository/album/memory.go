package album

import (
	"github.com/eugenius1/go-gin-rest/internal/models"
)

type memoryRepo struct {
	albums []models.Album
}

// Deprecated.
func NewMemoryRepo() *memoryRepo {
	return &memoryRepo{
		albums: []models.Album{
			{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
			{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		},
	}
}

func (r *memoryRepo) ListAlbums() ([]models.Album, error) {
	return r.albums, nil
}

func (r *memoryRepo) GetAlbumByID(id string) (models.Album, error) {
	for _, a := range r.albums {
		if a.ID == id {
			return a, nil
		}
	}
	return models.Album{}, nil
}

func (r *memoryRepo) CreateAlbum(album models.Album) error {
	r.albums = append(r.albums, album)
	return nil
}
