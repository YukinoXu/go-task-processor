package db

import (
	"log"

	"github.com/xuexiangxu/go-task-processor/internal/config"
	"github.com/xuexiangxu/go-task-processor/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := config.Cfg.DBUrl
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&model.Task{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	DB = db
	log.Println("Database migrated")
}
