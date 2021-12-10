package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lorenzo-mignola/review-manager/db"
	"github.com/lorenzo-mignola/review-manager/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	dbConnection := db.Connect()
	dbConnection.AutoMigrate(&models.Review{})

	r := gin.Default()
	r.GET("/reviews", func(c *gin.Context) {
		var reviews []models.Review
		dbConnection.Find(&reviews)
		c.JSON(200, reviews)
	})
	r.Run(":3030")
	log.Print("Server started")
}
