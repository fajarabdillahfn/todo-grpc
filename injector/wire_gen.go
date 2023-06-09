// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc"
	"github.com/fajarabdillahfn/todo-grpc/internal/repository/postgres"
	"github.com/fajarabdillahfn/todo-grpc/internal/usecase/v1"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *grpc.Handler {
	taskRepository := postgres.NewPostgresTaskRepository(db)
	taskUseCase := v1.NewTaskUseCaseV1(taskRepository)
	handler := grpc.NewTaskHandler(taskUseCase)
	return handler
}
