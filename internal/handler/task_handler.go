package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xuexiangxu/go-task-processor/internal/service"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/tasks", service.CreateTask)
	r.GET("/tasks/:id", service.GetTask)
}
