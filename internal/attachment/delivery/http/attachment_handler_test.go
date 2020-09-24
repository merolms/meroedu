package http_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
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
		mockUCase := new(mocks.AttachmentUseCase)
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		write, err := writer.CreateFormField("course_id")
		if err != nil {
			writer.Close()
			t.Error(err)
		}
		write.Write([]byte(strconv.Itoa(1)))
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
		mockUCase := new(mocks.AttachmentUseCase)
		e := echo.New()
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		write, err := writer.CreateFormField("course_id")
		if err != nil {
			writer.Close()
			t.Error(err)
		}
		write.Write([]byte(strconv.Itoa(1)))
		req, err := http.NewRequest(echo.POST, "attachments", body)
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
		mockUCase := new(mocks.AttachmentUseCase)
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		write, err := writer.CreateFormField("course_id")
		if err != nil {
			writer.Close()
			t.Error(err)
		}
		write.Write([]byte(strconv.Itoa(1)))
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

func TestDownloadAttachment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUCase := new(mocks.AttachmentUseCase)
		rootDirectory, err := os.Getwd()
		if err != nil {
			t.Error(err)
		}
		path := rootDirectory + "/" + "attachment_handler.go"
		mockUCase.On("DownloadAttachment", mock.Anything, mock.AnythingOfType("string")).Return(path, nil)
		e := echo.New()
		req, err := http.NewRequest(echo.GET, "/attachments/download?file=attachment_handler.go", strings.NewReader(""))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler := attachmenthttp.AttachmentHandler{
			AttachmentUseCase: mockUCase,
		}
		err = handler.DownloadAttachment(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockUCase.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		mockUCase := new(mocks.AttachmentUseCase)
		mockUCase.On("DownloadAttachment", mock.Anything, mock.AnythingOfType("string")).Return("", domain.ErrNotFound)
		e := echo.New()
		req, err := http.NewRequest(echo.GET, "/attachments/download?file=abc.txt", strings.NewReader(""))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler := attachmenthttp.AttachmentHandler{
			AttachmentUseCase: mockUCase,
		}
		err = handler.DownloadAttachment(c)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)
		mockUCase.AssertExpectations(t)
	})
}
