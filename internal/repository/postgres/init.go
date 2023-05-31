package postgres

import (
	"github.com/fajarabdillahfn/todo-grpc/internal/repository"
	"gorm.io/gorm"
)

type postgresTaskRepository struct {
	Conn *gorm.DB
}

func NewPostgresTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &postgresTaskRepository{
		Conn: conn,
	}
}
