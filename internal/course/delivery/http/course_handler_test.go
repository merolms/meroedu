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

	courseHTTP "github.com/meroedu/meroedu/internal/course/delivery/http"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/domain/mocks"
)

func TestGetAll(t *testing.T) {
	var mockCourse domain.Course
	err := faker.FakeData(&mockCourse)
	assert.NoError(t, err)
	mockUCase := new(mocks.CourseUseCase)
	mockListCourse := make([]domain.Course, 0)
	mockListCourse = append(mockListCourse, mockCourse)
	limit := "10"
	mockUCase.On("GetAll", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(mockListCourse, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/courses?start=0&limit="+limit, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := courseHTTP.CourseHandler{
		CourseUseCase: mockUCase,
	}
	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetAllError(t *testing.T) {
	mockUCase := new(mocks.CourseUseCase)
	limit := "10"
	mockUCase.On("GetAll", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, domain.ErrInternalServerError)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/courses?start=0&limit="+limit, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := courseHTTP.CourseHandler{
		CourseUseCase: mockUCase,
	}
	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	var mockCourse domain.Course
	err := faker.FakeData(&mockCourse)
	assert.NoError(t, err)

	mockUCase := new(mocks.CourseUseCase)

	num := int(mockCourse.ID)

	mockUCase.On("GetByID", mock.Anything, int64(num)).Return(mockCourse, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/courses/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("courses/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := courseHTTP.CourseHandler{
		CourseUseCase: mockUCase,
	}
	err = handler.GetByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestCreateCourse(t *testing.T) {
	mockCourse := domain.Course{
		Title:       "Title",
		Description: "Content",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tempmockCourse := mockCourse
	tempmockCourse.ID = 0
	mockUCase := new(mocks.CourseUseCase)

	j, err := json.Marshal(tempmockCourse)
	assert.NoError(t, err)

	mockUCase.On("CreateCourse", mock.Anything, mock.AnythingOfType("*domain.Course")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/courses", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/courses")

	handler := courseHTTP.CourseHandler{
		CourseUseCase: mockUCase,
	}
	err = handler.CreateCourse(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)
}
