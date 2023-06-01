package grpc

import (
	"context"
	"reflect"
	"testing"

	"github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc/task_grpc"
	"github.com/fajarabdillahfn/todo-grpc/internal/model"
	"github.com/fajarabdillahfn/todo-grpc/internal/usecase"
)

func Test_handler_CreateTask(t *testing.T) {
	uMock := usecase.TaskUseCaseMock{
		CreateFunc: func(ctx context.Context, task *model.Task) (uint, error) {
			return 1, nil
		},
	}

	type fields struct {
		TaskUC usecase.TaskUseCase
	}
	type args struct {
		ctx  context.Context
		task *task_grpc.TaskInput
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData *task_grpc.Id
		wantErr  bool
	}{
		{
			name:   "normal",
			fields: fields{TaskUC: &uMock},
			args: args{
				ctx: context.Background(),
				task: &task_grpc.TaskInput{
					Title:       "Test",
					Description: "Test desc",
				},
			},
			wantData: &task_grpc.Id{
				Id: 1,
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

			gotData, err := h.CreateTask(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("handler.CreateTask error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("handler.CreateTask = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
