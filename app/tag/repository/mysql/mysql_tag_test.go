package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/meroedu/meroedu/app/domain"
	tagMysqlRepo "github.com/meroedu/meroedu/app/tag/repository/mysql"
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
			ID: 1, Name: "IT", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}
	rows := sqlmock.NewRows([]string{"id", "name", "updated_at", "created_at"}).
		AddRow(mockTags[0].ID, mockTags[0].Name, mockTags[0].UpdatedAt, mockTags[0].CreatedAt)

	query := `SELECT id,name, updated_at, created_at FROM tags ORDER BY created_at DESC LIMIT \?,\?`
	mock.ExpectQuery(query).WillReturnRows(rows)
	c := tagMysqlRepo.InitMysqlRepository(db)
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

	query := `SELECT id,name,updated_at,created_at FROM tags WHERE ID = \?`
	mock.ExpectQuery(query).WillReturnRows(row)
	c := tagMysqlRepo.InitMysqlRepository(db)
	tag, err := c.GetByID(context.TODO(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, tag)
}

func TestCreateTag(t *testing.T) {
	c := &domain.Tag{
		Name:      "Programming",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `INSERT  tags SET name=\?,  updated_at=\? , created_at=\?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(c.Name, c.UpdatedAt, c.CreatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := tagMysqlRepo.InitMysqlRepository(db)
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

	repo := tagMysqlRepo.InitMysqlRepository(db)
	err = repo.DeleteTag(context.TODO(), int64(tag_id))
	assert.NoError(t, err)
}

func TestUpdateTag(t *testing.T) {
	c := &domain.Tag{
		ID:        12,
		Name:      "Programming",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `UPDATE tags set name=\?, updated_at=\? WHERE ID = \?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(c.Name, c.UpdatedAt, c.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := tagMysqlRepo.InitMysqlRepository(db)
	err = repo.UpdateTag(context.TODO(), c)
	assert.NoError(t, err)
}
