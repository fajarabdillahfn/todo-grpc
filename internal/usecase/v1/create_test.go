package v1

import (
	"context"
	"testing"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
	"github.com/fajarabdillahfn/todo-grpc/internal/repository"
)

func Test_usecase_Create(t *testing.T) {
	rMock := &repository.TaskRepositoryMock{
		CreateFunc: func(ctx context.Context, task *model.Task) (uint, error) {
			return 1, nil
		},
	}

	type fields struct {
		TaskRepo repository.TaskRepository
	}
	type args struct {
		ctx  context.Context
		task *model.Task
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData uint
		wantErr  bool
	}{
		{
			name: "normal",
			fields: fields{
				TaskRepo: rMock,
			},
			args: args{
				ctx: context.Background(),
				task: &model.Task{
					Title:       "Test",
					Description: "The test",
					IsCompleted: false,
				},
			},
			wantData: 1,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := taskUseCaseV1{
				taskRepo: tt.fields.TaskRepo,
			}

			gotData, err := u.Create(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotData != tt.wantData {
				t.Errorf("usecase.Create() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
