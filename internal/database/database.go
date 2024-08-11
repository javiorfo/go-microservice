package database

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBDataConnection struct {
	Host        string
	Port        string
	DBName      string
	User        string
	Password    string
	ShowSQLInfo bool
}

var DBinstance *gorm.DB

func (db DBDataConnection) Connect() error {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		db.Host,
		db.Port,
		db.DBName,
		db.User,
		db.Password)

	loggerSQL := logger.Default.LogMode(logger.Info)
	if db.ShowSQLInfo {
		loggerSQL = logger.Discard
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: loggerSQL,
	})
	if err != nil {
		return err
	}

	log.Info("Connected to DB!")
	database.Logger = loggerSQL
	DBinstance = database
	return nil
}
