package models

import (
	"github.com/jinzhu/gorm"
)

// Movie Model for the Movie table
type Movie struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primary_key"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Genre    string `json:"genre"`
}

// MovieCollection - A []Movie
type MovieCollection struct {
	Movies []Movie
}

// Archive - Collection Pointer
var Archive *MovieCollection
