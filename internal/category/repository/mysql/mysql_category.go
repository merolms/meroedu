package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
)

type mysqlRepository struct {
	conn *sql.DB
}

// Init will create an object that represent the category's Repository interface
func Init(db *sql.DB) domain.CategoryRepository {
	return &mysqlRepository{
		conn: db,
	}
}
func (m *mysqlRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Category, err error) {
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

	result = make([]domain.Category, 0)
	for rows.Next() {
		t := domain.Category{}
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

func (m *mysqlRepository) GetAll(ctx context.Context, start int, limit int) (res []domain.Category, err error) {
	query := `SELECT id,name,updated_at,created_at FROM categories ORDER BY created_at DESC LIMIT ?,?`

	res, err = m.fetch(ctx, query, start, limit)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (m *mysqlRepository) GetByID(ctx context.Context, id int64) (res *domain.Category, err error) {
	query := `SELECT id,name,updated_at,created_at FROM categories WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}
	var category domain.Category
	if len(list) > 0 {
		category = list[0]
	} else {
		return &category, domain.ErrNotFound
	}

	return &category, nil
}
func (m *mysqlRepository) GetByName(ctx context.Context, name string) (res *domain.Category, err error) {
	query := `SELECT id,name,updated_at,created_at FROM categories WHERE name = ?`
	list, err := m.fetch(ctx, query, name)
	if err != nil {
		return nil, err
	}
	var category domain.Category
	if len(list) > 0 {
		category = list[0]
	} else {
		return &category, domain.ErrNotFound
	}

	return &category, nil
}

func (m *mysqlRepository) CreateCategory(ctx context.Context, a *domain.Category) (err error) {
	query := `INSERT categories SET name=?,updated_at=?,created_at=?`
	stmt, err := m.conn.PrepareContext(ctx, query)
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

func (m *mysqlRepository) DeleteCategory(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM categories WHERE id = ?"

	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAffected != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", rowsAffected)
		return
	}

	return
}
func (m *mysqlRepository) UpdateCategory(ctx context.Context, ar *domain.Category) (err error) {
	query := `UPDATE categories set name=?,updated_at=? WHERE ID = ?`

	stmt, err := m.conn.PrepareContext(ctx, query)
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
