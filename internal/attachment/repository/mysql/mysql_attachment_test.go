package mysql_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	mysqlrepo "github.com/meroedu/meroedu/internal/attachment/repository/mysql"
	"github.com/meroedu/meroedu/internal/domain"
)

func TestCreateAttachment(t *testing.T) {
	a := domain.Attachment{
		ID:   12,
		Name: "3ddba0fa-28a2-4d99-be66-85acdd2c9a80.png",
		Size: 30,
		Type: "image/png",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `INSERT attachments SET title=\?, description=\?, name=\?, size=\?, type=\?`

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(a.Title, a.Description, a.Name, a.Size, a.Type).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := mysqlrepo.Init(db)
	err = repo.CreateAttachment(context.TODO(), a)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), a.ID)
}
