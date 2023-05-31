package grpc

import (
	"time"

	"github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc/task_grpc"
	"github.com/fajarabdillahfn/todo-grpc/internal/model"
	"github.com/fajarabdillahfn/todo-grpc/internal/usecase"
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	taskUseCase usecase.TaskUseCase
	task_grpc.UnimplementedTaskServiceServer
}

func NewTaskServerGrpc(gServer *grpc.Server, taskUseCase usecase.TaskUseCase) {
	taskServer := &server{
		taskUseCase: taskUseCase,
	}

	task_grpc.RegisterTaskServiceServer(gServer, taskServer)
	reflection.Register(gServer)
}

func (s *server) transformTaskRPC(task *model.Task) *task_grpc.Task {
	if task == nil {
		return nil
	}

	created_at := &google_protobuf.Timestamp{
		Seconds: task.CreatedAt.Unix(),
	}

	updated_at := &google_protobuf.Timestamp{
		Seconds: task.UpdatedAt.Unix(),
	}

	return &task_grpc.Task{
		Id:          uint64(task.ID),
		Title:       task.Title,
		Description: task.Description,
		IsComplete:  task.IsCompleted,
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
	}
}

func (s *server) transformTaskData(task *task_grpc.Task) *model.Task {
	created_at := time.Unix(task.GetCreatedAt().GetSeconds(), 0)
	updated_at := time.Unix(task.GetUpdatedAt().GetSeconds(), 0)

	return &model.Task{
		ID:          uint(task.GetId()),
		Title:       task.GetTitle(),
		Description: task.GetDescription(),
		IsCompleted: task.GetIsComplete(),
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
	}
}

func (s *server) transformTaskUpdateData(task *task_grpc.TaskUpdate) *model.TaskUpdate {
	return &model.TaskUpdate{
		ID:          uint(task.GetId()),
		Title:       task.GetTitle(),
		Description: task.GetDescription(),
		IsCompleted: task.IsComplete,
	}
}
