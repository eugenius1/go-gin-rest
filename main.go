package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/eugenius1/go-gin-rest/internal/db"
	albumrepo "github.com/eugenius1/go-gin-rest/internal/db/repository/album"
	"github.com/eugenius1/go-gin-rest/internal/http/route"
	"github.com/eugenius1/go-gin-rest/internal/service"
	"github.com/eugenius1/go-gin-rest/internal/service/album"
)

func main() {
	// Database
	db, err := db.SetupFromEnv("public")
	if err != nil {
		panic(err)
	}

	services := service.Services{
		Albums: album.NewService(albumrepo.NewRepo(db)),
	}

	// Router
	router := gin.Default()
	route.RegisterRoutes(router, services)

	addr := os.Getenv("HTTP_ADDR")
	err = router.Run(addr)
	if err != nil {
		panic(err)
	}
}
