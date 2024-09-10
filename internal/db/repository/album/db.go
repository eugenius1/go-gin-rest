package album

import (
	"gorm.io/gorm"

	"github.com/eugenius1/go-gin-rest/internal/models"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}

func (s *repo) ListAlbums() ([]models.Album, error) {
	var albums []models.Album
	if err := s.db.Find(&albums).Error; err != nil {
		return nil, err
	}
	return albums, nil
}

func (s *repo) GetAlbumByID(id string) (models.Album, error) {
	album := models.Album{ID: id}
	if err := s.db.First(&album).Error; err != nil {
		return models.Album{}, err
	}
	return album, nil
}

func (s *repo) CreateAlbum(album models.Album) error {
	return s.db.Create(&album).Error
}
