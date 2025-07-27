package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xuexiangxu/go-task-processor/internal/config"
	"github.com/xuexiangxu/go-task-processor/internal/db"
	"github.com/xuexiangxu/go-task-processor/internal/handler"
	"github.com/xuexiangxu/go-task-processor/internal/mq"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	mq.InitRabbitMQ()

	r := gin.Default()
	handler.RegisterRoutes(r)

	r.Run(":" + config.Cfg.Port)
}
