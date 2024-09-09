package albums

import (
	"github.com/eugenius1/go-gin-rest/internal/albums/storage"
	models "github.com/eugenius1/go-gin-rest/internal/models/albums"
)

type Service interface {
	ListAlbums() ([]models.Album, error)
	GetAlbumByID(id string) (models.Album, error)
	CreateAlbum(album models.Album) error
}

type service struct {
	storage storage.Storage
}

func NewService(storage storage.Storage) Service {
	return &service{
		storage: storage,
	}
}

func (s *service) ListAlbums() ([]models.Album, error) {
	return s.storage.ListAlbums()
}

func (s *service) GetAlbumByID(id string) (models.Album, error) {
	return s.storage.GetAlbumByID(id)
}

func (s *service) CreateAlbum(album models.Album) error {
	return s.storage.CreateAlbum(album)
}
