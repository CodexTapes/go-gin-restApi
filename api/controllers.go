package api

import (
	"net/http"

	models "github.com/CodexTapes/go-gin-restApi/models"
	"github.com/gin-gonic/gin"
)

// GetAllMovies GET /movies handler
func GetAllMovies(c *gin.Context) {
	var movies []models.Movie

	err := models.DB.Find(&movies)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to get movies"})
	}

	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// CreateMovie Movie Handler
func CreateMovie(c *gin.Context) {
	var movie models.Movie
	c.BindJSON(&movie)

	err := models.DB.Create(&movie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": "failed to Create Movie"})

		return
	}

	c.JSON(http.StatusOK, &movie)
}

// GetMovie Handler
// func GetMovie(c *gin.Context) {

// }

// UpdateMovie Handler
// func UpdateMovie(c *gin.Context) {

// }

// DeleteMovie Handler
// func DeleteMovie(c *gin.Context) {

// }
