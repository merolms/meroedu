package http_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	contentHTTP "github.com/meroedu/meroedu/internal/content/delivery/http"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/domain/mocks"
)

func TestGetAll(t *testing.T) {
	var mockContent domain.Content
	err := faker.FakeData(&mockContent)
	assert.NoError(t, err)
	mockUCase := new(mocks.ContentUseCase)
	mockList := make([]domain.Content, 0)
	mockList = append(mockList, mockContent)
	limit := "10"
	mockUCase.On("GetAll", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(mockList, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/contents?start=0&limit="+limit, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := contentHTTP.ContentHandler{
		ContentUseCase: mockUCase,
	}
	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetAllError(t *testing.T) {
	mockUCase := new(mocks.ContentUseCase)
	limit := "10"
	mockUCase.On("GetAll", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, domain.ErrInternalServerError)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/contents?start=0&limit="+limit, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := contentHTTP.ContentHandler{
		ContentUseCase: mockUCase,
	}
	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	var mockContent domain.Content
	err := faker.FakeData(&mockContent)
	assert.NoError(t, err)

	mockUCase := new(mocks.ContentUseCase)

	num := int(mockContent.ID)

	mockUCase.On("GetByID", mock.Anything, int64(num)).Return(&mockContent, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/contents/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("contents/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := contentHTTP.ContentHandler{
		ContentUseCase: mockUCase,
	}
	err = handler.GetByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestCreateContent(t *testing.T) {
	f := make(url.Values)
	f.Set("title", "simple title")
	f.Set("description", "description")
	f.Set("lesson_id", "1")
	f.Set("content", "Here we go")
	f.Set("content_type", "formatted-text")

	mockContent := domain.Content{
		Title:       "Title",
		Description: "description",
		ContentType: domain.ContentIsFormattedText,
		Content:     "Here we go",
		LessonID:    1,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	mockUCase := new(mocks.ContentUseCase)

	mockUCase.On("CreateContent", mock.Anything, mock.AnythingOfType("*domain.Content")).Return(&mockContent, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/contents", strings.NewReader(f.Encode()))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/contents")

	handler := contentHTTP.ContentHandler{
		ContentUseCase: mockUCase,
	}
	err = handler.CreateContent(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestDeleteContent(t *testing.T) {
	var mockContent domain.Content
	err := faker.FakeData(&mockContent)
	assert.NoError(t, err)

	mockUCase := new(mocks.ContentUseCase)

	num := int(mockContent.ID)

	mockUCase.On("DeleteContent", mock.Anything, int64(num)).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/contents/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("contents/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := contentHTTP.ContentHandler{
		ContentUseCase: mockUCase,
	}
	err = handler.DeleteContent(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUCase.AssertExpectations(t)

}
