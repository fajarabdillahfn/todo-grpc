package v1

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
	"github.com/fajarabdillahfn/todo-grpc/internal/repository"
)

func Test_usecase_Update(t *testing.T) {
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	trueBool := true

	rMock := &repository.TaskRepositoryMock{
		GetByIDFunc: func(ctx context.Context, id uint) (*model.Task, error) {
			return &model.Task{
				ID:          1,
				Title:       "Hi",
				Description: "Hi description",
				IsCompleted: false,
				CreatedAt:   theTime,
				UpdatedAt:   theTime,
			}, nil
		},
		UpdateFunc: func(ctx context.Context, task *model.Task) (*model.Task, error) {
			return &model.Task{
				ID:          1,
				Title:       "Halo",
				Description: "Halo description",
				IsCompleted: true,
				CreatedAt:   theTime,
				UpdatedAt:   theTime,
			}, nil
		},
	}

	type fields struct {
		TaskRepo repository.TaskRepository
	}
	type args struct {
		ctx  context.Context
		task *model.TaskUpdate
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
				task: &model.TaskUpdate{
					Title:       "Halo",
					Description: "Halo description",
					IsCompleted: &trueBool,
				},
			},
			wantData: &model.Task{
				ID:          1,
				Title:       "Halo",
				Description: "Halo description",
				IsCompleted: true,
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

			gotData, err := u.Update(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("usecase.Create() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
