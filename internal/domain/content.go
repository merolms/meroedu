package domain

import (
	"context"
	"errors"
	"mime/multipart"
)

// type ContentType string

// ContentType ...
// type ContentType interface {
// 	fmt.Stringer
// 	CalExpr() string
// }

// ContentType ...
type ContentType struct {
	Type string `json:"type" validator:"required"`
}

// Content Type
var (
	ContentIsImage         = ContentType{"image"}
	ContentIsFile          = ContentType{"file"}
	ContentIsFormattedText = ContentType{"formatted-text"}
	ContentIsEmbeddedURL   = ContentType{"embed-url"}
)

func (u ContentType) String() (string, error) {
	switch u {
	case ContentIsImage:
		return ContentIsImage.Type, nil
	case ContentIsFile:
		return ContentIsFile.Type, nil
	case ContentIsFormattedText:
		return ContentIsFormattedText.Type, nil
	case ContentIsEmbeddedURL:
		return ContentIsEmbeddedURL.Type, nil
	}
	return "", errors.New("unsupported content type")
}

// Content ...
type Content struct {
	ID          int64          `json:"id,omitempty"`
	LessonID    int64          `json:"lesson_id,omitempty"`
	Title       string         `json:"title" validator:"required"`
	Description string         `json:"description,omitempty"`
	FileHeader  string         `json:"file_header,omitempty"`
	ContentType ContentType    `json:"content_type" validator:"required"`
	Content     string         `json:"content" validator:"required"`
	Name        string         `json:"name,omitempty"`
	Size        int64          `json:"file_size,omitempty"`
	File        multipart.File `json:"-" faker:"-"`
	EmbedURL    string         `json:"embed_url,omitempty"`
	Caption     string         `json:"caption,omitempty"`
	UpdatedAt   int64          `json:"updated_at,omitempty"`
	CreatedAt   int64          `json:"created_at,omitempty"`
}

// ContentUseCase represent the Content's repository contract
type ContentUseCase interface {
	GetAll(ctx context.Context, start int, limit int) ([]Content, error)
	GetByID(ctx context.Context, id int64) (*Content, error)
	UpdateContent(ctx context.Context, Content *Content, id int64) (*Content, error)
	CreateContent(ctx context.Context, Content *Content) (*Content, error)
	DeleteContent(ctx context.Context, id int64) error
	GetContentByLesson(ctx context.Context, lessonID int64) ([]Content, error)
	DownloadContent(ctx context.Context, fileName string) (string, error)
}

// ContentRepository represent the Content's repository
type ContentRepository interface {
	GetAll(ctx context.Context, start int, limit int) ([]Content, error)
	GetByID(ctx context.Context, id int64) (*Content, error)
	UpdateContent(ctx context.Context, Content *Content) error
	CreateContent(ctx context.Context, Content *Content) error
	DeleteContent(ctx context.Context, id int64) error
	GetContentCountByLesson(ctx context.Context, lessonID int64) (int, error)
	GetContentByLesson(ctx context.Context, lessonID int64) ([]Content, error)
}

// ContentStorage represent the content's storage contract
type ContentStorage interface {
	CreateContent(ctx context.Context, content Content) error
	DownloadContent(ctx context.Context, fileName string) (string, error)
}
