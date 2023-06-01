package grpc

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc/task_grpc"
	"github.com/fajarabdillahfn/todo-grpc/internal/model"
	"github.com/fajarabdillahfn/todo-grpc/internal/usecase"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_handler_GetTasks(t *testing.T) {
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)

	var res []model.Task = []model.Task{
		{
			ID:          1,
			Title:       "Test Task",
			Description: "The test task",
			IsCompleted: false,
			CreatedAt:   theTime,
			UpdatedAt:   theTime,
		},
		{
			ID:          2,
			Title:       "Test Task 2",
			Description: "The test task 2",
			IsCompleted: true,
			CreatedAt:   theTime,
			UpdatedAt:   theTime,
		},
	}

	uMock := &usecase.TaskUseCaseMock{
		GetAllFunc: func(ctx context.Context) (*[]model.Task, error) {
			return &res, nil
		},
	}

	type fields struct {
		TaskUC usecase.TaskUseCase
	}
	type args struct {
		ctx   context.Context
		empty emptypb.Empty
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData *task_grpc.TasksList
		wantErr  bool
	}{
		{
			name: "normal",
			fields: fields{
				TaskUC: uMock,
			},
			args: args{
				ctx: context.Background(),
			},
			wantData: &task_grpc.TasksList{
				Items: []*task_grpc.Task{
					{
						Id:          1,
						Title:       "Test Task",
						Description: "The test task",
						IsCompleted: false,
						UpdatedAt: &timestamppb.Timestamp{
							Seconds: timestamppb.New(theTime).GetSeconds(),
						},
						CreatedAt: &timestamppb.Timestamp{
							Seconds: timestamppb.New(theTime).GetSeconds(),
						},
					},
					{
						Id:          2,
						Title:       "Test Task 2",
						Description: "The test task 2",
						IsCompleted: true,
						UpdatedAt: &timestamppb.Timestamp{
							Seconds: timestamppb.New(theTime).GetSeconds(),
						},
						CreatedAt: &timestamppb.Timestamp{
							Seconds: timestamppb.New(theTime).GetSeconds(),
						},
					},
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

			gotData, err := h.GetTasks(tt.args.ctx, &tt.args.empty)
			if (err != nil) != tt.wantErr {
				t.Errorf("handler.GetTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("handler.GetTasks() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func Test_handler_GetTaskById(t *testing.T) {
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)

	uMock := &usecase.TaskUseCaseMock{
		GetByIDFunc: func(ctx context.Context, id uint) (*model.Task, error) {
			return &model.Task{
				ID:          1,
				Title:       "Test Task",
				Description: "The test task",
				IsCompleted: false,
				CreatedAt:   theTime,
				UpdatedAt:   theTime,
			}, nil
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
		wantData *task_grpc.Task
		wantErr  bool
	}{
		{
			name: "normal",
			fields: fields{
				TaskUC: uMock,
			},
			args: args{
				ctx: context.Background(),
				id: &task_grpc.Id{
					Id: 1,
				},
			},
			wantData: &task_grpc.Task{
				Id:          1,
				Title:       "Test Task",
				Description: "The test task",
				IsCompleted: false,
				UpdatedAt: &timestamppb.Timestamp{
					Seconds: timestamppb.New(theTime).GetSeconds(),
				},
				CreatedAt: &timestamppb.Timestamp{
					Seconds: timestamppb.New(theTime).GetSeconds(),
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

			gotData, err := h.GetTaskById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("handler.GetTaskById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("handler.GetTaskById() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
