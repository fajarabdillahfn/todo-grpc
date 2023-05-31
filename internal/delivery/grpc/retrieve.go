package grpc

import (
	"context"

	"github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc/task_grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) GetTasks(ctx context.Context, _ *emptypb.Empty) (*task_grpc.TasksList, error) {
	tasks, err := s.taskUseCase.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var tasksList = &task_grpc.TasksList{
		Items: []*task_grpc.Task{},
	}

	for _, t := range *tasks {
		taskRPC := s.transformTaskRPC(&t)

		tasksList.Items = append(tasksList.Items, taskRPC)
	}

	return tasksList, nil
}

func (s *server) GetTaskById (ctx context.Context, id *task_grpc.Id) (*task_grpc.Task, error) {
	task, err := s.taskUseCase.GetByID(ctx, uint(id.GetId()))
	if err != nil {
		return nil, err
	}

	res := s.transformTaskRPC(task)

	return res, nil
}
