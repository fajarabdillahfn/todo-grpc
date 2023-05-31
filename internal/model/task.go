package model

import (
	"time"
)

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"uniqueIndex" json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_complete"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
