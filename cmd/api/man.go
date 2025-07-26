package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xuexiangxu/go-task-processor/internal/config"
	"github.com/xuexiangxu/go-task-processor/internal/db"
	"github.com/xuexiangxu/go-task-processor/internal/handler"
)

func main() {
	config.LoadConfig()
	db.Connect()

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/tasks", handler.CreateTaskHandler)
		api.GET("/tasks/:id", handler.GetTaskHandler)
		api.GET("/tasks", handler.ListTaskHandler)
	}

	r.Run(":" + config.Cfg.Port)
}
