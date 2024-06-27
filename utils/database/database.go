package database

import (
	"backend-api/config"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(c config.ProgrammingConfig) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New("\r]n"),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		c.DBHost,
		c.DBUser,
		c.DBPass,
		c.DBName,
		c.DBPort,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Error("Terjadi kesalahan pada database, error: ", err.Error())
		return nil, err
	}

	fmt.Printf("connected to database: %s", c.DBName)
	return db, nil
}
