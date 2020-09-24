package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	repository "github.com/meroedu/meroedu/internal/user/repository/mysql"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "updated_at", "created_at"}).
		AddRow(1, "Dinesh Katwal", time.Now().Unix(), time.Now().Unix())

	query := "SELECT id, name, created_at, updated_at FROM user WHERE id=\\?"

	prep := mock.ExpectPrepare(query)
	userID := int64(1)
	prep.ExpectQuery().WithArgs(userID).WillReturnRows(rows)

	a := repository.NewMysqlUserRepository(db)

	course, err := a.GetByID(context.TODO(), userID)
	assert.NoError(t, err)
	assert.NotNil(t, course)
}
