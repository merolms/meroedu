package mysql_test

import (
	"context"
	"testing"
	"time"

	mysqlrepo "github.com/meroedu/meroedu/internal/course/repository/mysql"
	"github.com/meroedu/meroedu/internal/domain"
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
			ID:          1,
			Title:       "title-1",
			Description: "tell me nah",
			Author:      domain.User{ID: 1},
			Category:    domain.Category{ID: 1},
			UpdatedAt:   time.Now(), CreatedAt: time.Now(),
		},
	}
	rows := sqlmock.NewRows([]string{"id", "title", "description", "duration", "image_url", "status", "author_id", "category_id", "updated_at", "created_at"}).
		AddRow(mockCourses[0].ID, mockCourses[0].Title, mockCourses[0].Description, 20, "https://", domain.CourseInDraft, mockCourses[0].Author.ID, mockCourses[0].Category.ID, mockCourses[0].UpdatedAt, mockCourses[0].CreatedAt)

	query := `SELECT id,title, description, duration, image_url, status, author_id, category_id, updated_at, created_at FROM courses ORDER BY created_at DESC LIMIT \?,\?`
	mock.ExpectQuery(query).WillReturnRows(rows)
	c := mysqlrepo.Init(db)
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
	row := sqlmock.NewRows([]string{"id", "title", "description", "duration", "image_url", "status", "author_id", "category_id", "updated_at", "created_at"}).
		AddRow("1", "testing-2", "description", 20, "https://gogole.com/3432.jpg", domain.CourseInDraft, 0, 0, time.Now(), time.Now())

	query := `SELECT id,title, description, duration, image_url, status, author_id, category_id,updated_at, created_at FROM courses WHERE ID = \?`
	mock.ExpectQuery(query).WillReturnRows(row)
	c := mysqlrepo.Init(db)
	course, err := c.GetByID(context.TODO(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, *course)
}

func TestGetByTitle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	row := sqlmock.NewRows([]string{"id", "title", "description", "duration", "image_url", "status", "author_id", "category_id", "updated_at", "created_at"}).
		AddRow("1", "testing-2", "description", 20, "https://", domain.CourseArchived, 0, 0, time.Now(), time.Now())

	query := `SELECT id,title, description, duration, image_url, status, author_id, category_id,updated_at, created_at FROM courses WHERE title = \?`
	mock.ExpectQuery(query).WillReturnRows(row)
	c := mysqlrepo.Init(db)
	course, err := c.GetByTitle(context.TODO(), "testing-2")
	assert.NoError(t, err)
	assert.NotNil(t, course)
}
func TestCreateCourse(t *testing.T) {
	date := time.Now()
	c := &domain.Course{
		Title:       "Java Programming",
		Description: "Testing",
		Category: domain.Category{
			ID: 1,
		},
		Author: domain.User{
			ID:   1,
			Name: "Nepal kathmandu",
		},
		UpdatedAt: date,
		CreatedAt: date,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `INSERT courses SET title=\?, description=\?, duration=\?, status=\?, image_url=\?, author_id=\?, category_id=\?, updated_at=\?, created_at=\?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(c.Title, c.Description, c.Duration, c.Status, c.ImageURL, c.Author.ID, c.Category.ID, c.UpdatedAt, c.CreatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := mysqlrepo.Init(db)
	err = repo.CreateCourse(context.TODO(), c)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), c.ID)
}

func TestDeleteCourse(t *testing.T) {
	course_id := 12
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `DELETE FROM courses WHERE id = \?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(course_id).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := mysqlrepo.Init(db)
	err = repo.DeleteCourse(context.TODO(), int64(course_id))
	assert.NoError(t, err)
}

func TestUpdateCourse(t *testing.T) {
	c := &domain.Course{
		ID:        12,
		Title:     "Java Programming",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Category: domain.Category{
			ID: 1,
		},
		Author: domain.User{
			ID:   1,
			Name: "Nepal kathmandu",
		},
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `UPDATE courses set title=\?, description=\?, updated_at=\? WHERE ID = \?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(c.Title, c.Description, c.UpdatedAt, c.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := mysqlrepo.Init(db)
	err = repo.UpdateCourse(context.TODO(), c)
	assert.NoError(t, err)
}
