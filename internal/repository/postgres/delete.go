package postgres

import (
	"context"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
)

// Delete implements repository.TaskRepository
func (r *postgresTaskRepository) Delete(ctx context.Context, id uint) error {
	err := r.Conn.Delete(&model.Task{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
