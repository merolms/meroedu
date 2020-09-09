package http_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/labstack/echo/v4"
	attachmenthttp "github.com/meroedu/meroedu/internal/attachment/delivery/http"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpload(t *testing.T) {
	mockUCase := new(mocks.AttachmentUserCase)
	path := "/home/auzmor/Documents/product/mystuff/root/meroedu/internal/attachment/delivery/http/attachment_handler.go" //The path to upload the file
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	defer file.Close()
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		writer.Close()
		t.Error(err)
	}
	io.Copy(part, file)
	writer.Close()

	t.Run("success", func(t *testing.T) {
		mockUCase.On("CreateAttachment", mock.Anything, mock.AnythingOfType("domain.Attachment")).Return(nil, nil)
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "attachment/upload", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/attachment/upload")
		handler := attachmenthttp.AttachmentHandler{
			AttachmentUseCase: mockUCase,
		}

		err = handler.Upload(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		mockUCase.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		e := echo.New()
		mockUCase.On("CreateAttachment", mock.Anything, mock.AnythingOfType("domain.Attachment")).Return(nil, domain.ErrUnsupportedFileType)
		req, err := http.NewRequest(echo.POST, "attachment/upload", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/attachment/upload")
		handler := attachmenthttp.AttachmentHandler{
			AttachmentUseCase: mockUCase,
		}

		err = handler.Upload(c)
		require.Error(t, err)
		mockUCase.AssertExpectations(t)
	})
}
