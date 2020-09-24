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

	categoryHTTP "github.com/meroedu/meroedu/internal/category/delivery/http"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/domain/mocks"
)

func TestGetAll(t *testing.T) {
	var mockCategory domain.Category
	err := faker.FakeData(&mockCategory)
	assert.NoError(t, err)
	mockUCase := new(mocks.CategoryUseCase)
	mockList := make([]domain.Category, 0)
	mockList = append(mockList, mockCategory)
	limit := "10"
	mockUCase.On("GetAll", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(mockList, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/categories?start=0&limit="+limit, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := categoryHTTP.CategoryHandler{
		CategoryUseCase: mockUCase,
	}
	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetAllError(t *testing.T) {
	mockUCase := new(mocks.CategoryUseCase)
	limit := "10"
	mockUCase.On("GetAll", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, domain.ErrInternalServerError)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/categories?start=0&limit="+limit, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := categoryHTTP.CategoryHandler{
		CategoryUseCase: mockUCase,
	}
	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	var mockCategory domain.Category
	err := faker.FakeData(&mockCategory)
	assert.NoError(t, err)

	mockUCase := new(mocks.CategoryUseCase)

	num := int(mockCategory.ID)

	mockUCase.On("GetByID", mock.Anything, int64(num)).Return(&mockCategory, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/categories/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("categories/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := categoryHTTP.CategoryHandler{
		CategoryUseCase: mockUCase,
	}
	err = handler.GetByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestCreateCategory(t *testing.T) {
	mockCategory := domain.Category{
		Name:      "Title",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	tempmockCategory := mockCategory
	tempmockCategory.ID = 0
	mockUCase := new(mocks.CategoryUseCase)

	j, err := json.Marshal(tempmockCategory)
	assert.NoError(t, err)

	mockUCase.On("CreateCategory", mock.Anything, mock.AnythingOfType("*domain.Category")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/categories", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/categories")

	handler := categoryHTTP.CategoryHandler{
		CategoryUseCase: mockUCase,
	}
	err = handler.CreateCategory(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestUpdateCategory(t *testing.T) {
	mockTag := domain.Category{
		ID:        124,
		Name:      "category",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	tempmockCategory := mockTag
	mockUCase := new(mocks.CategoryUseCase)
	j, err := json.Marshal(tempmockCategory)
	assert.NoError(t, err)
	mockUCase.On("UpdateCategory", mock.Anything, mock.AnythingOfType("*domain.Category"), mock.AnythingOfType("int64")).Return(nil)
	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/categories/124", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/categories/:id")
	c.SetParamNames("id")
	c.SetParamValues("124")

	handler := categoryHTTP.CategoryHandler{
		CategoryUseCase: mockUCase,
	}
	err = handler.UpdateCategory(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestDeleteCategory(t *testing.T) {
	var mockTag domain.Category
	err := faker.FakeData(&mockTag)
	assert.NoError(t, err)

	mockUCase := new(mocks.CategoryUseCase)

	num := int(mockTag.ID)

	mockUCase.On("DeleteCategory", mock.Anything, int64(num)).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/categories/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("categories/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := categoryHTTP.CategoryHandler{
		CategoryUseCase: mockUCase,
	}
	err = handler.DeleteCategory(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUCase.AssertExpectations(t)
}
