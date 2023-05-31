package v1

import (
	"context"
	"time"
)

// Delete implements repository.TaskRepository
func (u *taskUseCaseV1) Delete(ctx context.Context, id uint) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	_, err := u.taskRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return u.taskRepo.Delete(ctx, id)
}
