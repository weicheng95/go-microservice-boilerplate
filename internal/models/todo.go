package models

import (
	"context"
	"time"

	"gitlab.com/kiplexlab/go-microservice-boilerplate/internal/db"
	"gorm.io/gorm"
)

type Todo struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	Title     string
	Content   string
	Active    bool `gorm:"default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type TodoModel struct{}

func (todoModel *TodoModel) Create(ctx context.Context, todo *Todo) (int, error) {
	result := db.GetDB().Create(todo)
	if result.Error != nil {
		return -1, result.Error
	}
	return 1, result.Error
}
