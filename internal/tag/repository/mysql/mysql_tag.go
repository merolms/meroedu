package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
)

type mysqlRepository struct {
	Conn *sql.DB
}

// Init will create an object that represent the tag's Repository interface
func Init(db *sql.DB) domain.TagRepository {
	return &mysqlRepository{
		Conn: db,
	}
}
func (m *mysqlRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Tag, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
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

	result = make([]domain.Tag, 0)
	for rows.Next() {
		t := domain.Tag{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
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

func (m *mysqlRepository) GetAll(ctx context.Context, start int, limit int) (res []domain.Tag, err error) {
	query := `SELECT id,name, updated_at, created_at FROM tags ORDER BY created_at DESC LIMIT ?,?`

	res, err = m.fetch(ctx, query, start, limit)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (m *mysqlRepository) GetByID(ctx context.Context, id int64) (res domain.Tag, err error) {
	query := `SELECT id,name,updated_at,created_at FROM tags WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Tag{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
func (m *mysqlRepository) GetByName(ctx context.Context, name string) (res domain.Tag, err error) {
	query := `SELECT id,name,updated_at,created_at FROM tags WHERE name = ?`
	list, err := m.fetch(ctx, query, name)
	if err != nil {
		return domain.Tag{}, err
	}
	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlRepository) CreateTag(ctx context.Context, a *domain.Tag) (err error) {
	query := `INSERT  tags SET name=?, updated_at=? , created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error("Error while preparing statement ", err)
		return
	}
	res, err := stmt.ExecContext(ctx, a.Name, a.UpdatedAt, a.CreatedAt)
	if err != nil {
		log.Error("Error while executing statement ", err)
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Error("Got Error from LastInsertId method: ", err)
		return
	}
	a.ID = lastID
	return
}

func (m *mysqlRepository) DeleteTag(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM tags WHERE id = ?"

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
func (m *mysqlRepository) UpdateTag(ctx context.Context, ar *domain.Tag) (err error) {
	query := `UPDATE tags set name=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, ar.Name, ar.UpdatedAt, ar.ID)
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
