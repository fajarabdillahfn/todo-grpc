package postgres

import (
	"context"
	"time"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
)

// Update implements repository.TaskRepository
func (r *postgresTaskRepository) Update(ctx context.Context, task *model.Task) (*model.Task, error) {
	task.UpdatedAt = time.Now()
	
	err := r.Conn.WithContext(ctx).Save(&task).Error
	if err != nil {
		return nil, err
	}

	return task, nil
}
