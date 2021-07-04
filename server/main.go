package main

import (
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
	log.Print("Server started")
}
