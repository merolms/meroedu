package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/pkg/log"
)

type mysqlRepository struct {
	conn *sql.DB
}

// Init will create an object that represent the tag's Repository interface
func Init(db *sql.DB) domain.TagRepository {
	return &mysqlRepository{
		conn: db,
	}
}
func (m *mysqlRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Tag, err error) {
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

func (m *mysqlRepository) GetAll(ctx context.Context, searchQuery string, start int, limit int) (res []domain.Tag, err error) {
	query := `SELECT id,name,updated_at,created_at FROM tags WHERE name like ? ORDER BY created_at DESC LIMIT ?,?`
	searchQuery = "%" + searchQuery + "%"
	res, err = m.fetch(ctx, query, searchQuery, start, limit)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (m *mysqlRepository) GetByID(ctx context.Context, id int64) (res *domain.Tag, err error) {
	query := `SELECT id,name,updated_at,created_at FROM tags WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}
	var tag domain.Tag
	if len(list) > 0 {
		tag = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return &tag, nil
}
func (m *mysqlRepository) GetByName(ctx context.Context, name string) (res *domain.Tag, err error) {
	query := `SELECT id,name,updated_at,created_at FROM tags WHERE name = ?`
	list, err := m.fetch(ctx, query, name)
	if err != nil {
		return nil, err
	}
	var tag domain.Tag
	if len(list) > 0 {
		tag = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return &tag, nil
}

func (m *mysqlRepository) CreateTag(ctx context.Context, a *domain.Tag) (err error) {
	query := `INSERT tags SET name=?,updated_at=?,created_at=?`
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

func (m *mysqlRepository) DeleteTag(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM tags WHERE id = ?"

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
func (m *mysqlRepository) UpdateTag(ctx context.Context, ar *domain.Tag) (err error) {
	query := `UPDATE tags set name=?,updated_at=? WHERE ID = ?`

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

func (m *mysqlRepository) CreateCourseTag(ctx context.Context, tagID int64, courseID int64) error {
	query := `INSERT courses_tags SET course_id=?,tag_id=?,created_at=?`
	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error("Error while preparing statement ", err)
		return err
	}
	res, err := stmt.ExecContext(ctx, courseID, tagID, time.Now().Unix())
	if err != nil {
		log.Error("Error while executing statement ", err)
		return err
	}
	_, err = res.LastInsertId()
	if err != nil {
		log.Error("Got Error from LastInsertId method: ", err)
		return err
	}
	return nil
}
func (m *mysqlRepository) DeleteCourseTag(ctx context.Context, tagID int64, courseID int64) error {
	query := "DELETE FROM courses_tags WHERE course_id = ? and tag_id= ?"

	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, courseID, tagID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", rowsAffected)
		return err
	}

	return nil
}
func (m *mysqlRepository) GetCourseTags(ctx context.Context, courseID int64) ([]domain.Tag, error) {
	query := `select t.id,t.name,t.updated_at,t.created_at from tags t, courses_tags ct where ct.course_id=?`
	tags, err := m.fetch(ctx, query, courseID)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (m *mysqlRepository) CreateLessonTag(ctx context.Context, tagID int64, lessonID int64) error {
	query := `INSERT lessons_tags SET lesson_id=?,tag_id=?,created_at=?`
	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error("Error while preparing statement ", err)
		return err
	}
	res, err := stmt.ExecContext(ctx, lessonID, tagID, time.Now().Unix())
	if err != nil {
		log.Error("Error while executing statement ", err)
		return err
	}
	_, err = res.LastInsertId()
	if err != nil {
		log.Error("Got Error from LastInsertId method: ", err)
		return err
	}
	return nil
}
func (m *mysqlRepository) DeleteLessonTag(ctx context.Context, tagID int64, lessonID int64) error {
	query := "DELETE FROM lessons_tags WHERE lesson_id = ? and tag_id= ?"

	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, lessonID, tagID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", rowsAffected)
		return err
	}

	return nil
}

func (m *mysqlRepository) GetLessonTags(ctx context.Context, lessonID int64) ([]domain.Tag, error) {
	query := `select t.id,t.name,t.updated_at,t.created_at from tags t, lessons_tags lt where lt.lesson_id=?`
	tags, err := m.fetch(ctx, query, lessonID)
	if err != nil {
		return nil, err
	}
	return tags, nil
}
