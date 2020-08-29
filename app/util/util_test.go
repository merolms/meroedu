package util_test

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/meroedu/meroedu/app/domain"
	"github.com/stretchr/testify/assert"

	"github.com/meroedu/meroedu/app/util"
)

func TestGetStatusCode(t *testing.T) {
	var response int

	response = util.GetStatusCode(nil)
	assert.Equal(t, response, http.StatusOK)

	response = util.GetStatusCode(domain.ErrInternalServerError)
	assert.Equal(t, response, http.StatusInternalServerError)

	response = util.GetStatusCode(domain.ErrNotFound)
	assert.Equal(t, response, http.StatusNotFound)

	response = util.GetStatusCode(domain.ErrConflict)
	assert.Equal(t, response, http.StatusConflict)

	response = util.GetStatusCode(domain.ErrBadParamInput)
	assert.Equal(t, response, http.StatusBadRequest)

	response = util.GetStatusCode(errors.New("unknown"))
	assert.Equal(t, response, http.StatusInternalServerError)

}

func TestIsRequestValid(t *testing.T) {
	mockCourse := domain.Course{}
	valid, err := util.IsRequestValid(mockCourse)

	assert.Error(t, err)
	assert.False(t, valid)

	mockCourse = domain.Course{
		Title:       "Title",
		Description: "Content",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	valid, err = util.IsRequestValid(mockCourse)
	assert.NoError(t, err)
	assert.True(t, valid)
}
