package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

// DB - Postgres Db Handle
var DB gorm.DB

// InitDB establishes postgres connection
func InitDB(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *gorm.DB {
	DbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(DbDriver, DbURL)

	// moved to main.go
	// defer db.Close()

	if err != nil {
		log.Fatal("Failed to connect to DB")
	}

	db.AutoMigrate(&Movie{})

	return db
}
