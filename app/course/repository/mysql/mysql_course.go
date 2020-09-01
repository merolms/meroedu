package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/meroedu/meroedu/app/domain"
	"github.com/sirupsen/logrus"
)

type mysqlRepository struct {
	Conn *sql.DB
}

// InitMysqlRepository will create an object that represent the course's Repository interface
func InitMysqlRepository(db *sql.DB) domain.CourseRepository {
	return &mysqlRepository{
		Conn: db,
	}
}
func (m *mysqlRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Course, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.Course, 0)
	for rows.Next() {
		t := domain.Course{}
		authorID := int64(0)
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Author.ID,
			&t.Category.ID,
			&t.UpdatedAt,
			&t.CreatedAt,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		t.Author = domain.User{
			ID: authorID,
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlRepository) GetAll(ctx context.Context, start int, limit int) (res []domain.Course, err error) {
	query := `SELECT id,title, description, author_id, category_id, updated_at, created_at FROM courses ORDER BY created_at DESC LIMIT ?,? `

	res, err = m.fetch(ctx, query, start, limit)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (m *mysqlRepository) GetByID(ctx context.Context, id int64) (res domain.Course, err error) {
	query := `SELECT id,title, author_id, updated_at, created_at
  						FROM courses WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Course{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlRepository) GetByTitle(ctx context.Context, title string) (res domain.Course, err error) {
	query := `SELECT id,title, author_id, updated_at, created_at
  						FROM courses WHERE title = ?`

	list, err := m.fetch(ctx, query, title)
	if err != nil {
		return
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return
}

func (m *mysqlRepository) CreateCourse(ctx context.Context, a *domain.Course) (err error) {
	query := `INSERT  courses SET title=?, description=?, author_id=?, category_id=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error("Error while preparing statement ", err)
		return
	}
	var fields []interface{}
	fields = append(fields, a.Title)
	fields = append(fields, a.Description)
	if a.Author.ID == 0 {
		fields = append(fields, nil)
	} else {
		fields = append(fields, a.Author.ID)
	}
	if a.Category.ID == 0 {
		fields = append(fields, nil)
	} else {
		fields = append(fields, a.Category.ID)
	}
	fmt.Println(fields...)
	res, err := stmt.ExecContext(ctx, fields...)
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

func (m *mysqlRepository) DeleteCourse(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM courses WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", rowsAfected)
		return
	}

	return
}
func (m *mysqlRepository) UpdateCourse(ctx context.Context, ar *domain.Course) (err error) {
	query := `UPDATE course set title=?, author_id=?, category_id=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, ar.Title, ar.Author.ID, ar.Category.ID, ar.UpdatedAt, ar.ID)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}
