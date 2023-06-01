package postgres

import (
	"context"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fajarabdillahfn/todo-grpc/internal/model"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
)

func (s *Suite) Test_repository_GetAll() {
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "is_completed", "created_at", "updated_at"}).
			AddRow(1, "Test Task", "The test task", false, theTime, theTime).
			AddRow(2, "Test Task 2", "The test task 2", true, theTime, theTime))

	ctx := context.Background()
	res, err := s.repository.GetAll(ctx)

	expectedData := []model.Task{
		{
			ID:          1,
			Title:       "Test Task",
			Description: "The test task",
			IsCompleted: false,
			CreatedAt:   theTime,
			UpdatedAt:   theTime,
		},
		{
			ID:          2,
			Title:       "Test Task 2",
			Description: "The test task 2",
			IsCompleted: true,
			CreatedAt:   theTime,
			UpdatedAt:   theTime,
		},
	}

	require.NoError(s.T(), err, err)
	require.Nil(s.T(), deep.Equal(&expectedData, res))
}

func (s *Suite) Test_repository_GetByID() {
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)

	id := 1

	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks"`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "is_completed", "created_at", "updated_at"}).
			AddRow(1, "Test Task", "The test task", false, theTime, theTime).
			AddRow(2, "Test Task 2", "The test task 2", true, theTime, theTime))

	ctx := context.Background()
	res, err := s.repository.GetByID(ctx, uint(id))

	expectedData := model.Task{
		ID:          1,
		Title:       "Test Task",
		Description: "The test task",
		IsCompleted: false,
		CreatedAt:   theTime,
		UpdatedAt:   theTime,
	}

	require.NoError(s.T(), err, err)
	require.Nil(s.T(), deep.Equal(&expectedData, res))
}
