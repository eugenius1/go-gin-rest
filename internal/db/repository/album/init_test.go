package album

import (
	"os"
	"testing"

	"gorm.io/gorm"

	dbpkg "github.com/eugenius1/go-gin-rest/internal/db"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var err error
	db, err = dbpkg.SetupFromEnv("test_album")
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
