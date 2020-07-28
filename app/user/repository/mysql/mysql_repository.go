package mysql

import (
	"context"
	"database/sql"

	"github.com/meroedu/course-api/app/domain"
)

type mysqlUserRepo struct {
	DB *sql.DB
}

// NewMysqlUserRepository will create an implementation of User.Repository
func NewMysqlUserRepository(db *sql.DB) domain.UserRepository {
	return &mysqlUserRepo{
		DB: db,
	}
}

func (m *mysqlUserRepo) getOne(ctx context.Context, query string, args ...interface{}) (res domain.User, err error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return domain.User{}, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	res = domain.User{}

	err = row.Scan(
		&res.ID,
		&res.Name,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	return
}

func (m *mysqlUserRepo) GetByID(ctx context.Context, id int64) (domain.User, error) {
	query := `SELECT id, name, created_at, updated_at FROM user WHERE id=?`
	return m.getOne(ctx, query, id)
}
