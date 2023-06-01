package v1

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
	"github.com/fajarabdillahfn/todo-grpc/internal/repository"
)

func Test_usecase_GetAll(t *testing.T) {
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

	rMock := &repository.TaskRepositoryMock{
		GetAllFunc: func(ctx context.Context) (*[]model.Task, error) {
			return &res, nil
		},
	}

	type fields struct {
		TaskRepo repository.TaskRepository
	}
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData *[]model.Task
		wantErr  bool
	}{
		{
			name: "normal",
			fields: fields{
				TaskRepo: rMock,
			},
			args: args{
				ctx: context.Background(),
			},
			wantData: &res,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := taskUseCaseV1{
				taskRepo: tt.fields.TaskRepo,
			}

			gotData, err := u.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("usecase.GetAll() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func Test_usecase_GetByID(t *testing.T) {
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)

	rMock := &repository.TaskRepositoryMock{
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
		TaskRepo repository.TaskRepository
	}
	type args struct {
		ctx context.Context
		id  uint
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData *model.Task
		wantErr  bool
	}{
		{
			name: "normal",
			fields: fields{
				TaskRepo: rMock,
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantData: &model.Task{
				ID:          1,
				Title:       "Test Task",
				Description: "The test task",
				IsCompleted: false,
				CreatedAt:   theTime,
				UpdatedAt:   theTime,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := taskUseCaseV1{
				taskRepo: tt.fields.TaskRepo,
			}

			gotData, err := u.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("usecase.GetByID() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
