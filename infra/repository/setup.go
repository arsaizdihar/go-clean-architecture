package repository

import (
	"go-ci/domain/entity"
	"go-ci/infra/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Setup(config *config.Config) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open(config.DBFile), &gorm.Config{
		Logger: newLogger,
	})
	db.Exec("PRAGMA foreign_keys = ON", nil)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Product{})

	return db
}
