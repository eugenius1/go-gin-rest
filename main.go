package main

import (
	"github.com/eugenius1/go-gin-rest/internal"
	"github.com/eugenius1/go-gin-rest/internal/albums"
	storage "github.com/eugenius1/go-gin-rest/internal/albums/storage/db"
	"github.com/eugenius1/go-gin-rest/internal/models"
	"github.com/eugenius1/go-gin-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Database
	db, err := models.SetupDBFromEnv("public")
	if err != nil {
		panic(err)
	}

	services := internal.Services{
		Albums: albums.NewService(storage.NewStorage(db)),
	}

	// Router
	router := gin.Default()
	routes.RegisterRoutes(router, services)

	err = router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
