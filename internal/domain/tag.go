package domain

import (
	"context"
)

// Tag ...
type Tag struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty" validate:"required"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
}

// TagUseCase represent the Tag's repository contract
type TagUseCase interface {
	GetAll(ctx context.Context, searchQuery string, start int, limit int) ([]Tag, error)
	GetByID(ctx context.Context, id int64) (*Tag, error)
	GetByName(ctx context.Context, name string) (*Tag, error)
	UpdateTag(ctx context.Context, Tag *Tag, id int64) error
	CreateTag(ctx context.Context, Tag *Tag) error
	DeleteTag(ctx context.Context, id int64) error
	CreateCourseTag(ctx context.Context, tagID int64, courseID int64) error
	DeleteCourseTag(ctx context.Context, tagID int64, courseID int64) error
	GetCourseTags(ctx context.Context, courseID int64) ([]Tag, error)
	CreateLessonTag(ctx context.Context, tagID int64, lessonID int64) error
	DeleteLessonTag(ctx context.Context, tagID int64, lessonID int64) error
	GetLessonTags(ctx context.Context, lessonID int64) ([]Tag, error)
}

// TagRepository represent the Tag's repository
type TagRepository interface {
	GetAll(ctx context.Context, searchQuery string, start int, limit int) ([]Tag, error)
	GetByID(ctx context.Context, id int64) (*Tag, error)
	GetByName(ctx context.Context, name string) (*Tag, error)
	UpdateTag(ctx context.Context, Tag *Tag) error
	CreateTag(ctx context.Context, Tag *Tag) error
	DeleteTag(ctx context.Context, id int64) error
	CreateCourseTag(ctx context.Context, tagID int64, courseID int64) error
	DeleteCourseTag(ctx context.Context, tagID int64, courseID int64) error
	GetCourseTags(ctx context.Context, courseID int64) ([]Tag, error)
	CreateLessonTag(ctx context.Context, tagID int64, lessonID int64) error
	DeleteLessonTag(ctx context.Context, tagID int64, lessonID int64) error
	GetLessonTags(ctx context.Context, lessonID int64) ([]Tag, error)
}
