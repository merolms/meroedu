package mysql

import (
	"database/sql"

	"github.com/meroedu/meroedu/internal/domain"
)

type mysqlRepository struct {
	Conn *sql.DB
}

// Init will create an object that represent the attachment's Repository interface
func Init(db *sql.DB) domain.AttachmentRepository {
	return &mysqlRepository{
		Conn: db,
	}
}
