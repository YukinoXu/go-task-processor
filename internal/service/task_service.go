package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/xuexiangxu/go-task-processor/internal/db"
	"github.com/xuexiangxu/go-task-processor/internal/model"
)

func CreateTask(ctx context.Context, task *model.Task) error {
	return db.DB.WithContext(ctx).Create(task).Error
}

func GetTaskByID(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	var task model.Task
	if err := db.DB.WithContext(ctx).First(&task, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func ListTasks(ctx context.Context, limit int, offset int) ([]model.Task, error) {
	var tasks []model.Task
	err := db.DB.WithContext(ctx).Limit(limit).Offset(offset).Order("created_at_desc").Find(&tasks).Error
	return tasks, err
}
