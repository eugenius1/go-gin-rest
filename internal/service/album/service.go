package album

import (
	"github.com/eugenius1/go-gin-rest/internal/models"
)

type Service interface {
	ListAlbums() ([]models.Album, error)
	GetAlbumByID(id string) (models.Album, error)
	CreateAlbum(album models.Album) error
}

type service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) ListAlbums() ([]models.Album, error) {
	return s.repo.ListAlbums()
}

func (s *service) GetAlbumByID(id string) (models.Album, error) {
	return s.repo.GetAlbumByID(id)
}

func (s *service) CreateAlbum(album models.Album) error {
	return s.repo.CreateAlbum(album)
}
