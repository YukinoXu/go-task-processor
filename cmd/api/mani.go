package main

import (
	"github.com/xuexiangxu/go-task-processor/internal/config"
	"github.com/xuexiangxu/go-task-processor/internal/db"
)

func main() {
	config.LoadConfig()
	db.Connect()
}
