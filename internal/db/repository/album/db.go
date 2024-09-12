package album

import (
	"time"

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

func (r *repo) ListAlbums() ([]models.Album, error) {
	var albums []models.Album
	if err := r.db.Find(&albums).Error; err != nil {
		return nil, err
	}
	return albums, nil
}

func (r *repo) GetAlbumByID(id string) (models.Album, error) {
	album := models.Album{ID: id}
	if err := r.db.First(&album).Error; err != nil {
		return models.Album{}, err
	}
	return album, nil
}

func (r *repo) CreateAlbum(album models.Album) error {
	return r.db.Create(&album).Error
}

func (r *repo) UpdateAlbum(album models.Album) error {
	// override input value
	album.UpdatedAt = time.Now()
	return r.db.Model(&album).Updates(&album).Error
}

func (r *repo) DeleteAlbum(id string) error {
	album := models.Album{ID: id}
	return r.db.Delete(&album).Error
}
