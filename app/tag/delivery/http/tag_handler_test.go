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

	"github.com/meroedu/meroedu/app/domain"
	"github.com/meroedu/meroedu/app/domain/mocks"
	tagHTTP "github.com/meroedu/meroedu/app/tag/delivery/http"
)

func TestGetAll(t *testing.T) {
	var mockTag domain.Tag
	err := faker.FakeData(&mockTag)
	assert.NoError(t, err)
	mockUCase := new(mocks.TagUseCase)
	mockList := make([]domain.Tag, 0)
	mockList = append(mockList, mockTag)
	limit := "10"
	mockUCase.On("GetAll", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(mockList, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/tags?start=0&limit="+limit, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetAllError(t *testing.T) {
	mockUCase := new(mocks.TagUseCase)
	limit := "10"
	mockUCase.On("GetAll", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, domain.ErrInternalServerError)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/tags?start=0&limit="+limit, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	var mockTag domain.Tag
	err := faker.FakeData(&mockTag)
	assert.NoError(t, err)

	mockUCase := new(mocks.TagUseCase)

	num := int(mockTag.ID)

	mockUCase.On("GetByID", mock.Anything, int64(num)).Return(mockTag, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/tags/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("tags/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	err = handler.GetByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestCreateTag(t *testing.T) {
	mockTag := domain.Tag{
		Name:      "Title",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tempmockTag := mockTag
	tempmockTag.ID = 0
	mockUCase := new(mocks.TagUseCase)

	j, err := json.Marshal(tempmockTag)
	assert.NoError(t, err)

	mockUCase.On("CreateTag", mock.Anything, mock.AnythingOfType("*domain.Tag")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/tags", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/tags")

	handler := tagHTTP.TagHandler{
		TagUseCase: mockUCase,
	}
	err = handler.CreateTag(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)
}
