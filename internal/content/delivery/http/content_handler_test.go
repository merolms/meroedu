package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

	mockUCase.On("GetByID", mock.Anything, int64(num)).Return(mockContent, nil)

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
	mockContent := domain.Content{
		Title:     "Title",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tempmockContent := mockContent
	tempmockContent.ID = 0
	mockUCase := new(mocks.ContentUseCase)

	j, err := json.Marshal(tempmockContent)
	assert.NoError(t, err)

	mockUCase.On("CreateContent", mock.Anything, mock.AnythingOfType("*domain.Content")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/contents", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

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
