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

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/domain/mocks"
	lessonHTTP "github.com/meroedu/meroedu/internal/lesson/delivery/http"
)

func TestGetAll(t *testing.T) {
	var mockLesson domain.Lesson
	err := faker.FakeData(&mockLesson)
	assert.NoError(t, err)
	mockUCase := new(mocks.LessonUseCase)
	mockList := make([]domain.Lesson, 0)
	mockList = append(mockList, mockLesson)
	limit := "10"
	mockUCase.On("GetAll", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(mockList, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/lessons?start=0&limit="+limit, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := lessonHTTP.LessonHandler{
		LessonUseCase: mockUCase,
	}
	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetAllError(t *testing.T) {
	mockUCase := new(mocks.LessonUseCase)
	limit := "10"
	mockUCase.On("GetAll", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, domain.ErrInternalServerError)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/lessons?start=0&limit="+limit, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := lessonHTTP.LessonHandler{
		LessonUseCase: mockUCase,
	}
	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	var mockLesson domain.Lesson
	err := faker.FakeData(&mockLesson)
	assert.NoError(t, err)

	mockUCase := new(mocks.LessonUseCase)

	num := int(mockLesson.ID)

	mockUCase.On("GetByID", mock.Anything, int64(num)).Return(mockLesson, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/lessons/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("lessons/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := lessonHTTP.LessonHandler{
		LessonUseCase: mockUCase,
	}
	err = handler.GetByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestCreateLesson(t *testing.T) {
	mockLesson := domain.Lesson{
		Title:     "Title",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tempmockLesson := mockLesson
	tempmockLesson.ID = 0
	mockUCase := new(mocks.LessonUseCase)

	j, err := json.Marshal(tempmockLesson)
	assert.NoError(t, err)

	mockUCase.On("CreateLesson", mock.Anything, mock.AnythingOfType("*domain.Lesson")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/lessons", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/lessons")

	handler := lessonHTTP.LessonHandler{
		LessonUseCase: mockUCase,
	}
	err = handler.CreateLesson(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)
}
