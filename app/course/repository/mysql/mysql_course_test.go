package mysql_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	courseMysqlRepo "github.com/meroedu/course-api/app/course/repository/mysql"
	"github.com/meroedu/course-api/app/domain"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockCourses := []domain.Course{
		domain.Course{
			ID: 1, Title: "title-1",
			Author: domain.User{ID: 1}, UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}
	rows := sqlmock.NewRows([]string{"id", "title", "author_id", "updated_at", "created_at"}).
		AddRow(mockCourses[0].ID, mockCourses[0].Title, mockCourses[0].Author.ID, mockCourses[0].UpdatedAt, mockCourses[0].CreatedAt)

	query := `SELECT id,title, author_id, updated_at, created_at FROM courses ORDER BY created_at DESC LIMIT \?,\?`
	mock.ExpectQuery(query).WillReturnRows(rows)
	c := courseMysqlRepo.InitMysqlRepository(db)
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
	row := sqlmock.NewRows([]string{"id", "title", "author_id", "updated_at", "created_at"}).
		AddRow("1", "testing-2", 0, time.Now(), time.Now())

	query := `SELECT id,title, author_id, updated_at, created_at FROM courses WHERE ID = \?`
	mock.ExpectQuery(query).WillReturnRows(row)
	c := courseMysqlRepo.InitMysqlRepository(db)
	course, err := c.GetByID(context.TODO(), 1)
	fmt.Println(course)
	assert.NoError(t, err)
	assert.NotNil(t, course)
}

func TestGetByTitle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	row := sqlmock.NewRows([]string{"id", "title", "author_id", "updated_at", "created_at"}).
		AddRow("1", "testing-2", 0, time.Now(), time.Now())

	query := `SELECT id,title, author_id, updated_at, created_at FROM courses WHERE title = \?`
	mock.ExpectQuery(query).WillReturnRows(row)
	c := courseMysqlRepo.InitMysqlRepository(db)
	course, err := c.GetByTitle(context.TODO(), "testing-2")
	fmt.Println(course)
	assert.NoError(t, err)
	assert.NotNil(t, course)
}
func TestCreateCourse(t *testing.T) {
	c := &domain.Course{
		Title:     "Java Programming",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Author: domain.User{
			ID:   1,
			Name: "Nepal kathmandu",
		},
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `INSERT  courses SET title=\?, author_id=\?, updated_at=\? , created_at=\?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(c.Title, c.Author.ID, c.UpdatedAt, c.CreatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := courseMysqlRepo.InitMysqlRepository(db)
	err = repo.CreateCourse(context.TODO(), c)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), c.ID)
}


func TestDeleteCourse(t *testing.T) {
	course_id:=12
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `DELETE FROM courses WHERE id = \?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(course_id).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := courseMysqlRepo.InitMysqlRepository(db)
	err = repo.DeleteCourse(context.TODO(), int64(course_id))
	assert.NoError(t, err)
}

func TestUpdateCourse(t *testing.T) {
	c := &domain.Course{
		ID: 	   12,
		Title:     "Java Programming",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Author: domain.User{
			ID:   1,
			Name: "Nepal kathmandu",
		},
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `UPDATE course set title=\?, author_id=\?, updated_at=\? WHERE ID = \?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(c.Title, c.Author.ID, c.UpdatedAt, c.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := courseMysqlRepo.InitMysqlRepository(db)
	err = repo.UpdateCourse(context.TODO(), c)
	assert.NoError(t, err)
}
