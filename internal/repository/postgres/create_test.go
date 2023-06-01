package postgres

import (
	"context"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fajarabdillahfn/todo-grpc/internal/model"
	"github.com/stretchr/testify/require"
)

func (s *Suite) Test_repository_Create() {
	newTask := model.Task{
		Title:       "Test",
		Description: "Test task",
		IsCompleted: false,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tasks" ("title","description","is_completed","created_at","updated_at") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
		WithArgs(newTask.Title, newTask.Description, newTask.IsCompleted, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	s.mock.ExpectCommit()

	id, err := s.repository.Create(context.Background(), &newTask)

	expectedData := 1

	require.NoError(s.T(), err, err)
	require.True(s.T(), id == uint(expectedData))
}
