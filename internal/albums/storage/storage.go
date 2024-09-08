package storage

import "github.com/eugenius1/go-gin-rest/internal/albums/models"

type Storage interface {
	ListAlbums() ([]models.Album, error)
	GetAlbumByID(id string) (models.Album, error)
	CreateAlbum(album models.Album) error
}
