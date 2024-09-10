package album

import "github.com/eugenius1/go-gin-rest/internal/models"

type Repo interface {
	ListAlbums() ([]models.Album, error)
	GetAlbumByID(id string) (models.Album, error)
	CreateAlbum(album models.Album) error
}