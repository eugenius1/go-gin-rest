package models

import (
	"github.com/eugenius1/go-gin-rest/internal/models/albums"
)

func All() []any {
	return albums.All
}
