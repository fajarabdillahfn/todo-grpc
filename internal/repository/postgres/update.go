package postgres

import (
	"context"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
)

// Update implements repository.TaskRepository
func (r *postgresTaskRepository) Update(ctx context.Context, task *model.Task) (*model.Task, error) {
	err := r.Conn.WithContext(ctx).Save(&task).Error
	if err != nil {
		return nil, err
	}

	return task, nil
}
