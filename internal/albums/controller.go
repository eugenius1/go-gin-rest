package albums

import (
	"net/http"

	"github.com/eugenius1/go-gin-rest/internal/albums/models"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service Service
}

func NewController(service Service) *Controller {
	return &Controller{
		Service: service,
	}
}

func (c *Controller) RegisterRoutes(router gin.IRouter) {
	router = router.Group("/albums")
	router.GET("/", c.getAlbums)
	router.POST("/", c.postAlbums)
	router.GET("/:id", c.getAlbumByID)
}

func (c *Controller) getAlbums(ctx *gin.Context) {
	albums, err := c.Service.ListAlbums()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, albums)
}

func (c *Controller) postAlbums(ctx *gin.Context) {
	var newAlbum models.Album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.Service.CreateAlbum(newAlbum); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newAlbum)
}

func (c *Controller) getAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	album, err := c.Service.GetAlbumByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if album.ID == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	ctx.JSON(http.StatusOK, album)
}
