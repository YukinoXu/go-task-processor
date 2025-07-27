package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xuexiangxu/go-task-processor/internal/db"
	"github.com/xuexiangxu/go-task-processor/internal/model"
	"github.com/xuexiangxu/go-task-processor/internal/mq"
)

type TaskRequest struct {
	Type    string `json:"type" binding:"required"`
	Payload string `json:"payload" binding:"required"`
}

// POST /tasks
func CreateTask(c *gin.Context) {
	var req TaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := model.Task{
		ID:        uuid.New(),
		Type:      req.Type,
		Payload:   req.Payload,
		Status:    model.StatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save task"})
		return
	}

	taskJson, _ := json.Marshal(task)
	if err := mq.PublishTask(string(taskJson)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enqueue task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// GET /tasks/:id
func GetTask(c *gin.Context) {
	id := c.Param("id")
	uuidVal, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task model.Task
	if err := db.DB.First(&task, "id = ?", uuidVal).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTaskStatus(id uuid.UUID, status model.TaskStatus) error {
	return db.DB.Model(&model.Task{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     status,
			"updated_at": time.Now(),
		}).Error
}

func FinishTask(id uuid.UUID, result string, status model.TaskStatus) error {
	return db.DB.Model(&model.Task{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     status,
			"result":     result,
			"updated_at": time.Now(),
		}).Error
}
