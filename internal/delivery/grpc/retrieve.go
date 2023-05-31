package grpc

import (
	"context"

	"github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc/task_grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) GetTasks(_ *emptypb.Empty, stream task_grpc.TaskService_GetTasksServer) error {
	ctx := context.Background()

	tasks, err := s.taskUseCase.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, t := range *tasks {
		taskRPC := s.transformTaskRPC(&t)

		if err := stream.Send(taskRPC); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) GetTaskById (ctx context.Context, id *task_grpc.Id) (*task_grpc.Task, error) {
	task, err := s.taskUseCase.GetByID(ctx, uint(id.GetId()))
	if err != nil {
		return nil, err
	}

	res := s.transformTaskRPC(task)

	return res, nil
}
