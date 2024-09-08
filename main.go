package main

import (
	"github.com/eugenius1/go-gin-rest/internal"
	"github.com/eugenius1/go-gin-rest/internal/albums"
	"github.com/eugenius1/go-gin-rest/internal/albums/storage"
	"github.com/eugenius1/go-gin-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	services := internal.Services{
		Albums: albums.NewService(storage.NewMemoryStorage()),
	}

	routes.RegisterRoutes(router, services)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
