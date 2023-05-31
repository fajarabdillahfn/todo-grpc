package v1

import (
	"github.com/fajarabdillahfn/todo-grpc/internal/repository"
)

type taskUseCaseV1 struct {
	taskRepo repository.TaskRepository
}

func NewTaskUseCaseV1(rTask repository.TaskRepository) repository.TaskRepository {
	return &taskUseCaseV1{
		taskRepo: rTask,
	}
}
