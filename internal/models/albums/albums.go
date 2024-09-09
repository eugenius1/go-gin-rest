package albums

import (
	"gorm.io/gorm"
)

var All = []any{
	&Album{},
}

// Album represents data about a record Album.
type Album struct {
	gorm.Model
	ID     string  `json:"id"` // this json tag overrides the default one in gorm.Model
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
