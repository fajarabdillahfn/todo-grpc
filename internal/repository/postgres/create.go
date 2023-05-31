package postgres

import (
	"context"
	"time"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
)

// Create implements repository.TaskRepository
func (r *postgresTaskRepository) Create(ctx context.Context, task *model.Task) (uint, error) {
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	res := r.Conn.WithContext(ctx).Create(task)
	if res.Error != nil {
		return 0, res.Error
	}

	return task.ID, nil
}
