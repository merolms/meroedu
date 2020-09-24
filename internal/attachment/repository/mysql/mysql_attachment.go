package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
)

type mysqlRepository struct {
	conn *sql.DB
}

func (m *mysqlRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Attachment, err error) {
	rows, err := m.conn.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
		}
	}()

	result = make([]domain.Attachment, 0)
	for rows.Next() {
		t := domain.Attachment{}
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Name,
			&t.Size,
			&t.Type,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

// Init will create an object that represent the attachment's Repository interface
func Init(db *sql.DB) domain.AttachmentRepository {
	return &mysqlRepository{
		conn: db,
	}
}

// CreateAttachment ...
func (r mysqlRepository) CreateAttachment(ctx context.Context, a domain.Attachment) error {
	query := `INSERT attachments SET title=?,description=?,name=?,size=?,type=?,course_id=?,updated_at=?,created_at=?`
	stmt, err := r.conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error("error while preparing statement ", err)
		return err
	}
	timestamp := time.Now().Unix()
	res, err := stmt.ExecContext(ctx, a.Title, a.Description, a.Name, a.Size, a.Type, a.CourseID, timestamp, timestamp)
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

func (m *mysqlRepository) GetAttachmentByCourse(ctx context.Context, courseID int64) ([]domain.Attachment, error) {
	query := `SELECT id,title,description,name,size,type,updated_at,created_at FROM attachments WHERE course_id = ?`
	list, err := m.fetch(ctx, query, courseID)
	if err != nil {
		return nil, err
	}
	return list, nil
}
