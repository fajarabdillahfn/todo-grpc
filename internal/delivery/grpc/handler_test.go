package grpc

import (
	"testing"

	"github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc/task_grpc"
	"github.com/fajarabdillahfn/todo-grpc/internal/model"
	"github.com/fajarabdillahfn/todo-grpc/internal/usecase"
)

func Test_transformTaskData(t *testing.T) {
	uMock := usecase.TaskUseCaseMock{}
	h := Handler{
		taskUseCase:                    &uMock,
		UnimplementedTaskServiceServer: task_grpc.UnimplementedTaskServiceServer{},
	}

	taskProto := task_grpc.Task{
		Id:          1,
		Title:       "Hi",
		Description: "Hi Test",
		IsCompleted: false,
	}

	taskData := model.Task{
		ID:          1,
		Title:       "Hi",
		Description: "Hi Test",
		IsCompleted: false,
	}

	t.Run("normal", func(t *testing.T) {
		gotData := h.transformTaskData(&taskProto)
		if gotData.ID != taskData.ID && gotData.Title != taskData.Title && gotData.Description != taskData.Description && gotData.IsCompleted != taskData.IsCompleted {
			t.Errorf("handler.transformTaskData = %v, want %v", gotData, &taskData)
		}
	})
}

func Test_transformTaskUpdateData(t *testing.T) {
	uMock := usecase.TaskUseCaseMock{}
	h := Handler{
		taskUseCase:                    &uMock,
		UnimplementedTaskServiceServer: task_grpc.UnimplementedTaskServiceServer{},
	}

	title := "Hi"
	desc := "Hi Test"
	isCompleted := false

	taskProto := task_grpc.TaskUpdate{
		Id:          1,
		Title:       &title,
		Description: &desc,
		IsCompleted: &isCompleted,
	}

	taskData := model.TaskUpdate{
		ID:          1,
		Title:       title,
		Description: desc,
		IsCompleted: &isCompleted,
	}

	t.Run("normal", func(t *testing.T) {
		gotData := h.transformTaskUpdateData(&taskProto)
		if gotData.ID != taskData.ID && gotData.Title != taskData.Title && gotData.Description != taskData.Description && gotData.IsCompleted != taskData.IsCompleted {
			t.Errorf("handler.transformTaskUpdateData = %v, want %v", gotData, &taskData)
		}
	})
}

func Test_transformTaskRPC(t *testing.T) {
	uMock := usecase.TaskUseCaseMock{}
	h := Handler{
		taskUseCase:                    &uMock,
		UnimplementedTaskServiceServer: task_grpc.UnimplementedTaskServiceServer{},
	}

	taskProto := task_grpc.Task{
		Id:          1,
		Title:       "Hi",
		Description: "Hi Test",
		IsCompleted: false,
	}

	taskData := model.Task{
		ID:          1,
		Title:       "Hi",
		Description: "Hi Test",
		IsCompleted: false,
	}

	t.Run("normal", func(t *testing.T) {
		gotData := h.transformTaskRPC(&taskData)
		if gotData.Id != taskProto.Id && gotData.Title != taskProto.Title && gotData.Description != taskProto.Description && gotData.IsCompleted != taskProto.IsCompleted {
			t.Errorf("handler.transformTaskData = %v, want %v", gotData, &taskData)
		}
	})
}