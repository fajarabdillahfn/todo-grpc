package postgres

import (
	"context"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fajarabdillahfn/todo-grpc/internal/model"
	"github.com/stretchr/testify/require"
)

func (s *Suite) Test_repository_Update() {
	updatedTask := model.Task{
		ID:          1,
		Title:       "Hi",
		Description: "Hi description",
		IsCompleted: false,
	}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(`UPDATE "tasks" SET "title"=$1,"description"=$2,"is_completed"=$3,"created_at"=$4,"updated_at"=$5 WHERE "id" = $6`)).
		WithArgs(updatedTask.Title, updatedTask.Description, updatedTask.IsCompleted, sqlmock.AnyArg(), sqlmock.AnyArg(), updatedTask.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	task, err := s.repository.Update(context.Background(), &updatedTask)

	require.NoError(s.T(), err, err)

	require.True(s.T(), updatedTask.ID == task.ID)
	require.True(s.T(), updatedTask.Title == task.Title)
	require.True(s.T(), updatedTask.Description == task.Description)
	require.True(s.T(), updatedTask.IsCompleted == task.IsCompleted)
}
