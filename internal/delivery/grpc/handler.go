package grpc

import (
	"time"

	"github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc/task_grpc"
	"github.com/fajarabdillahfn/todo-grpc/internal/model"
	"github.com/fajarabdillahfn/todo-grpc/internal/usecase"
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
)

type Handler struct {
	taskUseCase usecase.TaskUseCase
	task_grpc.UnimplementedTaskServiceServer
}

func NewTaskHandler(taskUseCase usecase.TaskUseCase) *Handler {
	return &Handler{
		taskUseCase: taskUseCase,
	}
}

func (h *Handler) transformTaskRPC(task *model.Task) *task_grpc.Task {
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
		IsCompleted: task.IsCompleted,
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
	}
}

func (h *Handler) transformTaskData(task *task_grpc.Task) *model.Task {
	created_at := time.Unix(task.GetCreatedAt().GetSeconds(), 0)
	updated_at := time.Unix(task.GetUpdatedAt().GetSeconds(), 0)

	return &model.Task{
		ID:          uint(task.GetId()),
		Title:       task.GetTitle(),
		Description: task.GetDescription(),
		IsCompleted: task.GetIsCompleted(),
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
	}
}

func (h *Handler) transformTaskUpdateData(task *task_grpc.TaskUpdate) *model.TaskUpdate {
	return &model.TaskUpdate{
		ID:          uint(task.GetId()),
		Title:       task.GetTitle(),
		Description: task.GetDescription(),
		IsCompleted: task.IsCompleted,
	}
}
