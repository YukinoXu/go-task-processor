package db

import (
	"log"

	"github.com/xuexiangxu/go-task-processor/internal/config"
	"github.com/xuexiangxu/go-task-processor/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(postgres.Open(config.Cfg.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to PostgreSQL")

	err = DB.AutoMigrate(&model.Task{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	log.Println("Database migrated")
}
