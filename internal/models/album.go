package models

import (
	"time"

	"gorm.io/gorm"
)

// Album represents data about a record Album.
type Album struct {
	ID        string         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-" gorm:"<-:create"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
