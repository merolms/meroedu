package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/meroedu/meroedu/app/domain"
	lessonMysqlRepo "github.com/meroedu/meroedu/app/lesson/repository/mysql"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockLessons := []domain.Lesson{
		domain.Lesson{
			ID: 1, Title: "IT", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}
	rows := sqlmock.NewRows([]string{"id", "title", "updated_at", "created_at"}).
		AddRow(mockLessons[0].ID, mockLessons[0].Title, mockLessons[0].UpdatedAt, mockLessons[0].CreatedAt)

	query := `SELECT id,name, updated_at, created_at FROM lessons ORDER BY created_at DESC LIMIT \?,\?`
	mock.ExpectQuery(query).WillReturnRows(rows)
	c := lessonMysqlRepo.InitMysqlRepository(db)
	start, limit := 0, 10
	list, err := c.GetAll(context.TODO(), start, limit)
	assert.NoError(t, err)
	assert.Len(t, list, 1)

}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	row := sqlmock.NewRows([]string{"id", "name", "updated_at", "created_at"}).
		AddRow("1", "testing-2", time.Now(), time.Now())

	query := `SELECT id,name,updated_at,created_at FROM lessons WHERE ID = \?`
	mock.ExpectQuery(query).WillReturnRows(row)
	c := lessonMysqlRepo.InitMysqlRepository(db)
	lesson, err := c.GetByID(context.TODO(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, lesson)
}

func TestCreateLesson(t *testing.T) {
	c := &domain.Lesson{
		Title:     "Programming",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `INSERT  lessons SET title=\?,  updated_at=\? , created_at=\?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(c.Title, c.UpdatedAt, c.CreatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := lessonMysqlRepo.InitMysqlRepository(db)
	err = repo.CreateLesson(context.TODO(), c)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), c.ID)
}

func TestDeleteLesson(t *testing.T) {
	lesson_id := 12
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `DELETE FROM lessons WHERE id = \?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(lesson_id).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := lessonMysqlRepo.InitMysqlRepository(db)
	err = repo.DeleteLesson(context.TODO(), int64(lesson_id))
	assert.NoError(t, err)
}

func TestUpdateLesson(t *testing.T) {
	c := &domain.Lesson{
		ID:        12,
		Title:     "Programming",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `UPDATE lessons set title=\?, updated_at=\? WHERE ID = \?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(c.Title, c.UpdatedAt, c.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := lessonMysqlRepo.InitMysqlRepository(db)
	err = repo.UpdateLesson(context.TODO(), c)
	assert.NoError(t, err)
}
