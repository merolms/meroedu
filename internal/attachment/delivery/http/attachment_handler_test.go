package http_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	attachmenthttp "github.com/meroedu/meroedu/internal/attachment/delivery/http"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/domain/mocks"
)

func TestCreateAttachment(t *testing.T) {
	rootDirectory, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	path := rootDirectory + "/" + "attachment_handler.go" //The path to upload the file
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	defer file.Close()

	t.Run("success", func(t *testing.T) {
		mockUCase := new(mocks.AttachmentUserCase)
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", filepath.Base(path))
		if err != nil {
			writer.Close()
			t.Error(err)
		}
		io.Copy(part, file)
		writer.Close()
		mockUCase.On("CreateAttachment", mock.Anything, mock.AnythingOfType("domain.Attachment")).Return(nil, nil)
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/attachments", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/attachments")
		handler := attachmenthttp.AttachmentHandler{
			AttachmentUseCase: mockUCase,
		}
		err = handler.CreateAttachment(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		mockUCase.AssertExpectations(t)
	})
	t.Run("error with no file", func(t *testing.T) {
		mockUCase := new(mocks.AttachmentUserCase)
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "attachments", strings.NewReader(""))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/attachments")
		handler := attachmenthttp.AttachmentHandler{
			AttachmentUseCase: mockUCase,
		}

		err = handler.CreateAttachment(c)
		require.NoError(t, err)
	})
	t.Run("unsupported file", func(t *testing.T) {
		mockUCase := new(mocks.AttachmentUserCase)
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", filepath.Base(path))
		if err != nil {
			writer.Close()
			t.Error(err)
		}
		io.Copy(part, file)
		writer.Close()
		mockUCase.On("CreateAttachment", mock.Anything, mock.AnythingOfType("domain.Attachment")).Return(nil, domain.ErrUnsupportedFileType)
		e := echo.New()
		req, err := http.NewRequest(echo.POST, "attachments", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/attachments")
		handler := attachmenthttp.AttachmentHandler{
			AttachmentUseCase: mockUCase,
		}
		err = handler.CreateAttachment(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		mockUCase.AssertExpectations(t)
	})
}
