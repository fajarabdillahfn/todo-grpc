package grpc

import (
	"context"

	"github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc/task_grpc"
)

func (h *Handler) UpdateTask(ctx context.Context, task *task_grpc.TaskUpdate) (*task_grpc.Task, error) {
	taskData := h.transformTaskUpdateData(task)

	updatedTask, err := h.taskUseCase.Update(ctx, taskData)
	if err != nil {
		return nil, err
	}

	res := h.transformTaskRPC(updatedTask)

	return res, nil
}
