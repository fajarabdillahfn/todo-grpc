package grpc

import (
	"context"
	"fmt"

	"github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc/task_grpc"
)

func (h *Handler) DeleteTask(ctx context.Context, id *task_grpc.Id) (*task_grpc.DeleteResponse, error) {
	err := h.taskUseCase.Delete(ctx, uint(id.Id))
	if err != nil {
		return &task_grpc.DeleteResponse{
			Status: "delete error",
		}, err
	}

	return &task_grpc.DeleteResponse{
		Status: fmt.Sprintf("task %d deleted", id.GetId()),
	}, nil
}
