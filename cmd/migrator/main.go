package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/javiorfo/go-microservice/config"
	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/javiorfo/go-microservice/internal/database"
)

func main() {
	if err := config.DBDataConnection.Connect(); err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}
	db := database.DBinstance

	log.Info("Running migrations...")
	err := db.AutoMigrate(&model.Dummy{})
	if err != nil {
		log.Errorf("failed to migrate database: %v", err)
		return
	}

	log.Info("Migration completed succesfully!")
}
