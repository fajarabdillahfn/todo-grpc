package v1

import (
	"github.com/fajarabdillahfn/todo-grpc/internal/repository"
	"github.com/fajarabdillahfn/todo-grpc/internal/usecase"
)

type taskUseCaseV1 struct {
	taskRepo repository.TaskRepository
}

func NewTaskUseCaseV1(rTask repository.TaskRepository) usecase.TaskUseCase {
	return &taskUseCaseV1{
		taskRepo: rTask,
	}
}
