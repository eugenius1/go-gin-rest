package db

import (
	"gorm.io/gorm"

	models "github.com/eugenius1/go-gin-rest/internal/models/albums"
)

type storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *storage {
	return &storage{
		db: db,
	}
}

func (s *storage) ListAlbums() ([]models.Album, error) {
	var albums []models.Album
	if err := s.db.Find(&albums).Error; err != nil {
		return nil, err
	}
	return albums, nil
}

func (s *storage) GetAlbumByID(id string) (models.Album, error) {
	album := models.Album{ID: id}
	if err := s.db.First(&album).Error; err != nil {
		return models.Album{}, err
	}
	return album, nil
}

func (s *storage) CreateAlbum(album models.Album) error {
	return s.db.Create(&album).Error
}
