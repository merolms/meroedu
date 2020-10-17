package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	mysqlrepo "github.com/meroedu/meroedu/internal/attachment/repository/mysql"
	"github.com/meroedu/meroedu/internal/domain"
)

func TestCreateAttachment(t *testing.T) {
	a := domain.Attachment{
		ID:        12,
		Name:      "3ddba0fa-28a2-4d99-be66-85acdd2c9a80.png",
		Size:      30,
		Type:      "image/png",
		CourseID:  1,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening stub database connection", err)
	}
	query := `INSERT attachments SET title=\?,description=\?,name=\?,size=\?,type=\?,course_id=\?,updated_at=\?,created_at=\?`

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(a.Title, a.Description, a.Name, a.Size, a.Type, a.CourseID, a.UpdatedAt, a.CreatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	repo := mysqlrepo.Init(db)
	err = repo.CreateAttachment(context.TODO(), a)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), a.ID)
}

func TestGetAttachmentByCourse(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	row := sqlmock.NewRows([]string{"id", "title", "description", "name", "size", "type", "updated_at", "created_at"}).
		AddRow("1", "testing-2", "description", "name", 240, "application/pdf", time.Now().Unix(), time.Now().Unix())

	query := `SELECT id,title,description,name,size,type,updated_at,created_at FROM attachments WHERE course_id = \?`
	mock.ExpectQuery(query).WillReturnRows(row)
	c := mysqlrepo.Init(db)
	attachments, err := c.GetAttachmentByCourse(context.TODO(), 1)

	assert.NoError(t, err)
	assert.Equal(t, len(attachments), 1)
}
