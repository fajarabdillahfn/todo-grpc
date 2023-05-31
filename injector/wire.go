//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	grpcTaskDelivery "github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc"
	pgTaskRepo "github.com/fajarabdillahfn/todo-grpc/internal/repository/postgres"
	v1TaskUseCase "github.com/fajarabdillahfn/todo-grpc/internal/usecase/v1"
)

func InitializedService(db *gorm.DB) *grpcTaskDelivery.Handler {
	wire.Build(
		grpcTaskDelivery.NewTaskHandler,
		pgTaskRepo.NewPostgresTaskRepository,
		v1TaskUseCase.NewTaskUseCaseV1,
	)

	return &grpcTaskDelivery.Handler{}
}
