package main

import (
	"fmt"
	"log"
	"net"
	"os"

	grpcTaskDelivery "github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc"
	pgTaskRepo "github.com/fajarabdillahfn/todo-grpc/internal/repository/postgres"
	v1TaskUseCase "github.com/fajarabdillahfn/todo-grpc/internal/usecase/v1"
	"github.com/fajarabdillahfn/todo-grpc/pkg/db/postgres"
	"google.golang.org/grpc"
)

func main() {
	db := postgres.OpenDB()

	taskRepo := pgTaskRepo.NewPostgresTaskRepository(db)
	taskUseCase := v1TaskUseCase.NewTaskUseCaseV1(taskRepo)

	netListen, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err.Error())
	}

	server := grpc.NewServer()
	grpcTaskDelivery.NewTaskServerGrpc(server, taskUseCase)

	log.Printf("Server start at %v", netListen.Addr())
	err = server.Serve(netListen)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err.Error())
	}
}