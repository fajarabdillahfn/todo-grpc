package grpc

import (
	"context"

	"github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc/task_grpc"
)

func (s *server) UpdateTask(ctx context.Context, task *task_grpc.Task) (*task_grpc.Task, error) {
	taskData := s.transformTaskData(task)

	updatedTask, err := s.taskUseCase.Update(ctx, taskData)
	if err != nil {
		return nil, err
	}

	res := s.transformTaskRPC(updatedTask)

	return res, nil
}
