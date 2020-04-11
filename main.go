package main

import (
	"os"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

	"github.com/CodexTapes/go-gin-restApi/models"
)

func main() {
	router := gin.Default()

	router.GET("/movies", getAllMovies)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.initDB(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	router.Run()
}