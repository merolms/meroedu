package mysql

import (
	"context"
	"database/sql"

	"github.com/meroedu/meroedu/app/domain"
	"github.com/sirupsen/logrus"
)

type mysqlRepository struct {
	Conn *sql.DB
}

// InitMysqlRepository will create an object that represent the course's Repository interface
func InitMysqlRepository(db *sql.DB) domain.AttachmentRepository {
	return &mysqlRepository{
		Conn: db,
	}
}

// func (m *mysqlRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Course, err error) {
// 	rows, err := m.Conn.QueryContext(ctx, query, args...)
// 	if err != nil {
// 		logrus.Error(err)
// 		return nil, err
// 	}

// 	defer func() {
// 		errRow := rows.Close()
// 		if errRow != nil {
// 			logrus.Error(errRow)
// 		}
// 	}()

// 	result = make([]domain.Course, 0)
// 	for rows.Next() {
// 		t := domain.Course{}
// 		authorID := int64(0)
// 		err = rows.Scan(
// 			&t.ID,
// 			&t.Title,
// 			&t.Author.ID,
// 			&t.UpdatedAt,
// 			&t.CreatedAt,
// 		)

// 		if err != nil {
// 			logrus.Error(err)
// 			return nil, err
// 		}
// 		t.Author = domain.User{
// 			ID: authorID,
// 		}
// 		result = append(result, t)
// 	}

// 	return result, nil
// }

// func (m *mysqlRepository) GetAll(ctx context.Context, start int, limit int) (res []domain.Attachment, err error) {
// 	query := `SELECT id,title, author_id, updated_at, created_at FROM courses ORDER BY created_at DESC LIMIT ?,? `

// 	res, err = m.fetch(ctx, query, start, limit)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

func (m *mysqlRepository) Upload(ctx context.Context, a *domain.Attachment) {
	query := `INSERT courses SET id=?, file=?, updated_at=? , created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error("Error while preparing statement ", err)
		return
	}
	res, err := stmt.ExecContext(ctx, a.ID, a.File, a.UpdatedAt, a.CreatedAt)
	if err != nil {
		logrus.Error("Error while executing statement ", err)
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		logrus.Error("Got Error from LastInsertId method: ", err)
		return
	}
	a.ID = lastID
	return
}
