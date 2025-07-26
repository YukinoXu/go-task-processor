package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Type      string    `json:"type"`
	Payload   string    `json:"payload"`
	Status    string    `json:"status"`
	Result    string    `json:"result"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
