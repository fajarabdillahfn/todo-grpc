package postgres

import (
	"context"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
)

// GetAll implements repository.TaskRepository
func (r *postgresTaskRepository) GetAll(ctx context.Context) (*[]model.Task, error) {
	var tasksList *[]model.Task

	err := r.Conn.WithContext(ctx).Find(&tasksList).Error
	if err != nil {
		return nil, err
	}

	return tasksList, nil
}

// GetByID implements repository.TaskRepository
func (r *postgresTaskRepository) GetByID(ctx context.Context, id uint) (*model.Task, error) {
	var task *model.Task

	err := r.Conn.WithContext(ctx).First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return task, nil
}
