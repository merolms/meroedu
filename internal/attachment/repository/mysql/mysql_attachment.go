package mysql

import (
	"context"
	"database/sql"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
)

type mysqlRepository struct {
	conn *sql.DB
}

// Init will create an object that represent the attachment's Repository interface
func Init(db *sql.DB) domain.AttachmentRepository {
	return &mysqlRepository{
		conn: db,
	}
}

func (r mysqlRepository) CreateAttachment(ctx context.Context, a domain.Attachment) error {
	query := `INSERT attachments SET title=?, description=?, name=?, size=?, type=?`
	stmt, err := r.conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error("error while preparing statement ", err)
		return err
	}

	res, err := stmt.ExecContext(ctx, a.Title, a.Description, a.Name, a.Size, a.Type)
	if err != nil {
		log.Error("error while executing statement ", err)
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Error("got an error from LastInsertId method: ", err)
		return err
	}
	a.ID = lastID
	return nil
}
