package grpc

import (
	"context"
	"reflect"
	"testing"

	"github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc/task_grpc"
	"github.com/fajarabdillahfn/todo-grpc/internal/usecase"
)

func Test_handler_DeleteTask(t *testing.T) {
	uMock := usecase.TaskUseCaseMock{
		DeleteFunc: func(ctx context.Context, id uint) error {
			return nil
		},
	}

	type fields struct {
		TaskUC usecase.TaskUseCase
	}
	type args struct {
		ctx context.Context
		id  *task_grpc.Id
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData *task_grpc.DeleteResponse
		wantErr  bool
	}{
		{
			name:   "normal",
			fields: fields{TaskUC: &uMock},
			args: args{
				ctx: context.Background(),
				id: &task_grpc.Id{
					Id: 1,
				},
			},
			wantData: &task_grpc.DeleteResponse{
				Status: "task 1 deleted",
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

			gotData, err := h.DeleteTask(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("handler.DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("handler.DeleteTask() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
