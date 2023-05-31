package main

import (
	"log"
	"net"
	"os"

	"github.com/fajarabdillahfn/todo-grpc/injector"
	"github.com/fajarabdillahfn/todo-grpc/internal/delivery/grpc/task_grpc"
	"github.com/fajarabdillahfn/todo-grpc/pkg/db/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db := postgres.OpenDB()

	netListen, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err.Error())
	}

	server := grpc.NewServer()

	taskHandler := injector.InitializedService(db)

	task_grpc.RegisterTaskServiceServer(server, taskHandler)
	reflection.Register(server)

	log.Printf("Server start at %v", netListen.Addr())
	err = server.Serve(netListen)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err.Error())
	}
}
