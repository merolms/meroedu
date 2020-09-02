package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// NewHealthHandler ...
func NewHealthHandler(e *echo.Echo) {
	// Init HealthCheck
	e.GET("/", HealthCheck)
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"health": "UP",
	})
}
