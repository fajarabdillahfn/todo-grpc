package grpc

import (
	"context"

	"github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc/task_grpc"
)

func (s *server) CreateTask(ctx context.Context, task *task_grpc.TaskInput) (*task_grpc.Id, error) {
	taskRPC := task_grpc.Task{
		Title:       task.GetTitle(),
		Description: task.GetDescription(),
		IsCompleted:  false,
	}
	taskData := s.transformTaskData(&taskRPC)

	id, err := s.taskUseCase.Create(ctx, taskData)
	if err != nil {
		return nil, err
	}

	return &task_grpc.Id{
		Id: uint64(id),
	}, err
}
