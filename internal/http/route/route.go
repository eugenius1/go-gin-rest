package route

import (
	"github.com/gin-gonic/gin"

	"github.com/eugenius1/go-gin-rest/internal/http/controller/album"
	"github.com/eugenius1/go-gin-rest/internal/service"
)

func RegisterRoutes(router gin.IRouter, services service.Services) {
	v1 := router.Group("/v1")
	album.NewController(services.Albums).RegisterRoutes(v1)
}
