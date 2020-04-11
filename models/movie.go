package models

import (
	"github.com/jinzhu/gorm"
)

type Movie struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Director string `json:"director"`
	Genre string `json:"genre"`		
}