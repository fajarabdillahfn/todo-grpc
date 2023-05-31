package usecase

import (
	"context"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
)

type TaskUseCase interface {
	GetAll(ctx context.Context) (*[]model.Task, error)
	GetByID(ctx context.Context, id uint) (*model.Task, error)
	Create(ctx context.Context, task *model.Task) (uint, error)
	Update(ctx context.Context, task *model.Task) (*model.Task, error)
	Delete(ctx context.Context, id uint) error
}
