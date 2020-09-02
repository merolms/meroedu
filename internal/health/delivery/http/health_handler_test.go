package http_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	healthHTTP "github.com/meroedu/meroedu/internal/health/delivery/http"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	mockResponse = map[string]interface{}{
		"health": "UP",
	}
)

func TestGetHealth(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/", strings.NewReader(""))
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = healthHTTP.HealthCheck(c)
	require.NoError(t, err)
	assert.NoError(t, err)

}
