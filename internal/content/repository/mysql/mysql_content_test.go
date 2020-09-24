package mysql_test

import (
	"context"
	"testing"
	"time"

	mysqlrepo "github.com/meroedu/meroedu/internal/content/repository/mysql"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockContents := []domain.Content{
		domain.Content{
			ID: 1, Title: "IT", UpdatedAt: time.Now().Unix(), CreatedAt: time.Now().Unix(),
		},
	}
	rows := sqlmock.NewRows([]string{"id", "title", "updated_at", "created_at"}).
		AddRow(mockContents[0].ID, mockContents[0].Title, mockContents[0].UpdatedAt, mockContents[0].CreatedAt)

	query := `SELECT id,title, updated_at, created_at FROM contents ORDER BY created_at DESC LIMIT \?,\?`
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
	row := sqlmock.NewRows([]string{"id", "title", "updated_at", "created_at"}).
		AddRow("1", "testing-2", time.Now().Unix(), time.Now().Unix())

	query := `SELECT id,title,updated_at,created_at FROM contents WHERE ID = \?`
	mock.ExpectQuery(query).WillReturnRows(row)
	c := mysqlrepo.Init(db)
	tag, err := c.GetByID(context.TODO(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, tag)
}

func TestCreateContent(t *testing.T) {
	c := &domain.Content{
		Title:     "Programming",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `INSERT  contents SET title=\?,  updated_at=\? , created_at=\?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(c.Title, c.UpdatedAt, c.CreatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := mysqlrepo.Init(db)
	err = repo.CreateContent(context.TODO(), c)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), c.ID)
}

func TestDeleteContent(t *testing.T) {
	tag_id := 12
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `DELETE FROM contents WHERE id = \?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(tag_id).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := mysqlrepo.Init(db)
	err = repo.DeleteContent(context.TODO(), int64(tag_id))
	assert.NoError(t, err)
}

func TestUpdateContent(t *testing.T) {
	c := &domain.Content{
		ID:        12,
		Title:     "Programming",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `UPDATE contents set title=\?, updated_at=\? WHERE ID = \?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(c.Title, c.UpdatedAt, c.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := mysqlrepo.Init(db)
	err = repo.UpdateContent(context.TODO(), c)
	assert.NoError(t, err)
}
