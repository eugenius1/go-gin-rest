package storage

import models "github.com/eugenius1/go-gin-rest/internal/models/albums"

type Storage interface {
	ListAlbums() ([]models.Album, error)
	GetAlbumByID(id string) (models.Album, error)
	CreateAlbum(album models.Album) error
}
