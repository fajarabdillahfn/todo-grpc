package grpc

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc/task_grpc"
	"github.com/fajarabdillahfn/todo-grpc/internal/model"
	"github.com/fajarabdillahfn/todo-grpc/internal/usecase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_handler_UpdateTask(t *testing.T) {
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)

	title := "Hi"
	desc := "Hi test"
	isCompleted := false

	uMock := usecase.TaskUseCaseMock{
		UpdateFunc: func(ctx context.Context, task *model.TaskUpdate) (*model.Task, error) {
			return &model.Task{
				ID:          1,
				Title:       title,
				Description: desc,
				IsCompleted: isCompleted,
				CreatedAt:   theTime,
				UpdatedAt:   theTime,
			}, nil
		},
	}

	type fields struct {
		TaskUC usecase.TaskUseCase
	}
	type args struct {
		ctx  context.Context
		task *task_grpc.TaskUpdate
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData *task_grpc.Task
		wantErr  bool
	}{
		{
			name:   "normal",
			fields: fields{TaskUC: &uMock},
			args: args{
				ctx: context.Background(),
				task: &task_grpc.TaskUpdate{
					Id:          1,
					Title:       &title,
					Description: &desc,
					IsCompleted: &isCompleted,
				},
			},
			wantData: &task_grpc.Task{
				Id:          1,
				Title:       title,
				Description: desc,
				IsCompleted: isCompleted,
				UpdatedAt: &timestamppb.Timestamp{
					Seconds: timestamppb.New(theTime).Seconds,
				},
				CreatedAt: &timestamppb.Timestamp{
					Seconds: timestamppb.New(theTime).Seconds,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{
				taskUseCase:                    tt.fields.TaskUC,
				UnimplementedTaskServiceServer: task_grpc.UnimplementedTaskServiceServer{},
			}

			gotData, err := h.UpdateTask(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("handler.UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("handler.UpdateTask() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
