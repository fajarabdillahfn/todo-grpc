package v1

import (
	"context"
	"testing"

	"github.com/fajarabdillahfn/todo-grpc/internal/model"
	"github.com/fajarabdillahfn/todo-grpc/internal/repository"
)

func Test_usecase_Delete(t *testing.T) {
	rMock := &repository.TaskRepositoryMock{
		GetByIDFunc: func(ctx context.Context, id uint) (*model.Task, error) {
			return &model.Task{
				ID: 1,
			}, nil
		},
		DeleteFunc: func(ctx context.Context, id uint) error {
			return nil
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
		name    string
		fields  fields
		args    args
		wantErr bool
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
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := taskUseCaseV1{
				taskRepo: tt.fields.TaskRepo,
			}

			err := u.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
