package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/eugenius1/go-gin-rest/internal"
	"github.com/eugenius1/go-gin-rest/internal/albums"
)

func RegisterRoutes(router gin.IRouter, services internal.Services) {
	v1 := router.Group("/v1")
	albums.NewController(services.Albums).RegisterRoutes(v1)
}
