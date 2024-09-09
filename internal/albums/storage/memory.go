package storage

import models "github.com/eugenius1/go-gin-rest/internal/models/albums"

type memoryStorage struct {
	albums []models.Album
}

func NewMemoryStorage() Storage {
	return &memoryStorage{
		albums: []models.Album{
			{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
			{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		},
	}
}

func (s *memoryStorage) ListAlbums() ([]models.Album, error) {
	return s.albums, nil
}

func (s *memoryStorage) GetAlbumByID(id string) (models.Album, error) {
	for _, a := range s.albums {
		if a.ID == id {
			return a, nil
		}
	}
	return models.Album{}, nil
}

func (s *memoryStorage) CreateAlbum(album models.Album) error {
	s.albums = append(s.albums, album)
	return nil
}
