package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
	mysqlrepo "github.com/meroedu/meroedu/internal/tag/repository/mysql"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockTags := []domain.Tag{
		domain.Tag{
			ID: 1, Name: "IT", UpdatedAt: time.Now().Unix(), CreatedAt: time.Now().Unix(),
		},
	}
	rows := sqlmock.NewRows([]string{"id", "name", "updated_at", "created_at"}).
		AddRow(mockTags[0].ID, mockTags[0].Name, mockTags[0].UpdatedAt, mockTags[0].CreatedAt)

	query := `SELECT id,name,updated_at,created_at FROM tags WHERE name like \? ORDER BY created_at DESC LIMIT \?,\?`
	mock.ExpectQuery(query).WillReturnRows(rows)
	c := mysqlrepo.Init(db)
	start, limit := 0, 10
	list, err := c.GetAll(context.TODO(), "", start, limit)
	assert.NoError(t, err)
	assert.Len(t, list, 1)

}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	row := sqlmock.NewRows([]string{"id", "name", "updated_at", "created_at"}).
		AddRow("1", "testing-2", time.Now().Unix(), time.Now().Unix())

	query := `SELECT id,name,updated_at,created_at FROM tags WHERE ID = \?`
	mock.ExpectQuery(query).WillReturnRows(row)
	c := mysqlrepo.Init(db)
	tag, err := c.GetByID(context.TODO(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, tag)
}
func TestGetByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	row := sqlmock.NewRows([]string{"id", "name", "updated_at", "created_at"}).
		AddRow("1", "testing-2", time.Now().Unix(), time.Now().Unix())

	query := `SELECT id,name,updated_at,created_at FROM tags WHERE name = \?`
	mock.ExpectQuery(query).WillReturnRows(row)
	c := mysqlrepo.Init(db)
	category, err := c.GetByName(context.TODO(), "testing-2")
	assert.NoError(t, err)
	assert.NotNil(t, category)
}
func TestCreateTag(t *testing.T) {
	c := &domain.Tag{
		Name:      "Programming",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `INSERT tags SET name=\?,updated_at=\?,created_at=\?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(c.Name, c.UpdatedAt, c.CreatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := mysqlrepo.Init(db)
	err = repo.CreateTag(context.TODO(), c)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), c.ID)
}

func TestDeleteTag(t *testing.T) {
	tag_id := 12
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `DELETE FROM tags WHERE id = \?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(tag_id).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := mysqlrepo.Init(db)
	err = repo.DeleteTag(context.TODO(), int64(tag_id))
	assert.NoError(t, err)
}

func TestUpdateTag(t *testing.T) {
	c := &domain.Tag{
		ID:        12,
		Name:      "Programming",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `UPDATE tags set name=\?,updated_at=\? WHERE ID = \?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(c.Name, c.UpdatedAt, c.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := mysqlrepo.Init(db)
	err = repo.UpdateTag(context.TODO(), c)
	assert.NoError(t, err)
}

func TestDeleteCourseTag(t *testing.T) {
	tagID := 1
	courseID := 1
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `DELETE FROM courses_tags WHERE course_id = \? and tag_id= \?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(courseID, tagID).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := mysqlrepo.Init(db)
	err = repo.DeleteCourseTag(context.TODO(), int64(courseID), int64(tagID))
	assert.NoError(t, err)
}
func TestCreateCourseTag(t *testing.T) {
	var tagID int64 = 1
	var courseID int64 = 1
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `INSERT courses_tags SET course_id=\?,tag_id=\?,created_at=\?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(courseID, tagID, time.Now().Unix()).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := mysqlrepo.Init(db)
	err = repo.CreateCourseTag(context.TODO(), courseID, tagID)
	assert.NoError(t, err)
}
func TestGetCourseTags(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	row := sqlmock.NewRows([]string{"id", "name", "updated_at", "created_at"}).
		AddRow("1", "testing-2", time.Now().Unix(), time.Now().Unix())

	query := `select t.id,t.name,t.updated_at,t.created_at from tags t, courses_tags ct where ct.course_id=\?`
	mock.ExpectQuery(query).WillReturnRows(row)
	c := mysqlrepo.Init(db)
	tag, err := c.GetCourseTags(context.TODO(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, tag)
}
func TestCreateLessonTag(t *testing.T) {
	var tagID int64 = 1
	var lessonID int64 = 1
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `INSERT lessons_tags SET lesson_id=\?,tag_id=\?,created_at=\?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(lessonID, tagID, time.Now().Unix()).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := mysqlrepo.Init(db)
	err = repo.CreateLessonTag(context.TODO(), lessonID, tagID)
	assert.NoError(t, err)
}
func TestDeleteLessonTag(t *testing.T) {
	tagID := 1
	lessonID := 1
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `DELETE FROM lessons_tags WHERE lesson_id = \? and tag_id= \?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(lessonID, tagID).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := mysqlrepo.Init(db)
	err = repo.DeleteLessonTag(context.TODO(), int64(lessonID), int64(tagID))
	assert.NoError(t, err)
}

func TestGetLessonTags(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	row := sqlmock.NewRows([]string{"id", "name", "updated_at", "created_at"}).
		AddRow("1", "testing-2", time.Now().Unix(), time.Now().Unix())

	query := `select t.id,t.name,t.updated_at,t.created_at from tags t, lessons_tags lt where lt.lesson_id=\?`
	mock.ExpectQuery(query).WillReturnRows(row)
	c := mysqlrepo.Init(db)
	tag, err := c.GetLessonTags(context.TODO(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, tag)
}
