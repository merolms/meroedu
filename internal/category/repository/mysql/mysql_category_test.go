package mysql_test

import (
	"context"
	"testing"
	"time"

	mysqlrepo "github.com/meroedu/meroedu/internal/category/repository/mysql"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mockCategories := []domain.Category{
		domain.Category{
			ID: 1, Name: "IT", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}
	rows := sqlmock.NewRows([]string{"id", "name", "updated_at", "created_at"}).
		AddRow(mockCategories[0].ID, mockCategories[0].Name, mockCategories[0].UpdatedAt, mockCategories[0].CreatedAt)

	query := `SELECT id,name, updated_at, created_at FROM categories ORDER BY created_at DESC LIMIT \?,\?`
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
	row := sqlmock.NewRows([]string{"id", "name", "updated_at", "created_at"}).
		AddRow("1", "testing-2", time.Now(), time.Now())

	query := `SELECT id,name,updated_at,created_at FROM categories WHERE ID = \?`
	mock.ExpectQuery(query).WillReturnRows(row)
	c := mysqlrepo.Init(db)
	category, err := c.GetByID(context.TODO(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, category)
}

func TestGetByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	row := sqlmock.NewRows([]string{"id", "name", "updated_at", "created_at"}).
		AddRow("1", "testing-2", time.Now(), time.Now())

	query := `SELECT id,name,updated_at,created_at FROM categories WHERE name = \?`
	mock.ExpectQuery(query).WillReturnRows(row)
	c := mysqlrepo.Init(db)
	category, err := c.GetByName(context.TODO(), "testing-2")
	assert.NoError(t, err)
	assert.NotNil(t, category)
}

func TestCreateCategory(t *testing.T) {
	c := &domain.Category{
		Name:      "Programming",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `INSERT  categories SET name=\?,  updated_at=\? , created_at=\?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(c.Name, c.UpdatedAt, c.CreatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := mysqlrepo.Init(db)
	err = repo.CreateCategory(context.TODO(), c)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), c.ID)
}

func TestDeleteCategory(t *testing.T) {
	category_id := 12
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `DELETE FROM categories WHERE id = \?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(category_id).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := mysqlrepo.Init(db)
	err = repo.DeleteCategory(context.TODO(), int64(category_id))
	assert.NoError(t, err)
}

func TestUpdateCategory(t *testing.T) {
	c := &domain.Category{
		ID:        12,
		Name:      "Programming",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `UPDATE categories set name=\?, updated_at=\? WHERE ID = \?`
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(c.Name, c.UpdatedAt, c.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := mysqlrepo.Init(db)
	err = repo.UpdateCategory(context.TODO(), c)
	assert.NoError(t, err)
}
