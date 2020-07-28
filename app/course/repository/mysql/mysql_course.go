package mysql

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/meroedu/course-api/app/domain"
	"github.com/sirupsen/logrus"
)

type mysqlRepository struct {
	Conn *gorm.DB
}

// InitMysqlRepository will create an object that represent the course's Repository interface
func InitMysqlRepository(Conn *gorm.DB) domain.CourseRepository {
	return &mysqlRepository{Conn}
}

func (mysqlRepo *mysqlRepository) GetAll(ctx context.Context, skip int, limit int) (res []domain.Course, err error) {
	var courses []domain.Course
	fmt.Println("Start and limit", skip, limit)
	if err := mysqlRepo.Conn.Limit(limit).Offset(skip).Find(&courses).Error; err != nil {
		logrus.Error(err)
		return []domain.Course{}, err
	}
	return courses, nil
}

func (mysqlRepo *mysqlRepository) GetByID(ctx context.Context, id int64) (res domain.Course, err error) {
	course := domain.Course{}

	if err := mysqlRepo.Conn.Where("id=?", id).Find(&course).Error; err != nil {
		logrus.Error(err)
	}
	if course.ID <= 0 {
		return domain.Course{}, domain.ErrNotFound
	}
	return course, nil
}

func (mysqlRepo *mysqlRepository) GetByTitle(ctx context.Context, title string) (res domain.Course, err error) {
	course := domain.Course{}
	if err := mysqlRepo.Conn.Find(&course).Where("title=?", title); err != nil {
		logrus.Error(err)
	}
	if course.ID <= 0 {
		return course, domain.ErrNotFound
	}
	return course, nil
}
func (mysqlRepo *mysqlRepository) CreateCourse(ctx context.Context, course *domain.Course) (int64, error) {
	if err := mysqlRepo.Conn.Save(&course).Error; err != nil {
		logrus.Error(err)
		return 0, err
	}
	return course.ID, nil
}
