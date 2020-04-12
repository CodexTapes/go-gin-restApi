package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	api "github.com/CodexTapes/go-gin-restApi/api"
	models "github.com/CodexTapes/go-gin-restApi/models"
)

func main() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := models.InitDB(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	defer db.Close()

	router.GET("/movies", api.GetAllMovies)
	router.POST("/movies", api.CreateMovie)
	router.GET("/movies/:id", api.GetMovie)
	router.PUT("/movies/:id", api.UpdateMovie)
	router.DELETE("/movies/:id", api.DeleteMovie)

	router.Run(":3000")
}
