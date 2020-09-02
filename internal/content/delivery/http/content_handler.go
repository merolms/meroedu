package http

import (
	"net/http"
	"strconv"

	"strings"

	"github.com/labstack/echo/v4"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/util"
	"github.com/meroedu/meroedu/pkg/log"
)

// ResponseError represents the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ContentHandler ...
type ContentHandler struct {
	ContentUseCase domain.ContentUseCase
}

// NewContentHandler ...
func NewContentHandler(e *echo.Echo, us domain.ContentUseCase) {
	handler := &ContentHandler{
		ContentUseCase: us,
	}
	// Get Operation
	e.GET("/contents", handler.GetAll)
	e.GET("/contents/:id", handler.GetByID)
	e.GET("/contents/:id/", handler.GetByID)

	// Create/Add Operation
	e.POST("/contents", handler.CreateContent)

	// Update Operation
	e.PUT("/contents/:id", handler.GetByID)
	e.PUT("/contents/actions", handler.GetByID)

	// Remove/Delete Operation
	e.DELETE("/contents/:id", handler.GetByID)
}

// GetAll ...
func (c *ContentHandler) GetAll(echoContext echo.Context) error {
	log.Info("Calling GetAll Contents")
	ctx := echoContext.Request().Context()
	start, limit := 0, 10
	var err error
	for k, v := range echoContext.QueryParams() {
		switch k {
		case "start":
			val := strings.TrimSpace(v[0])
			if start, err = strconv.Atoi(val); err != nil {
				return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
			}
		case "limit":
			val := strings.TrimSpace(v[0])
			if limit, err = strconv.Atoi(val); err != nil {
				return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
			}
		}
	}

	list, err := c.ContentUseCase.GetAll(ctx, start, limit)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, list)
}

// GetByID ...
func (c *ContentHandler) GetByID(echoContext echo.Context) error {
	log.Info("Calling GetByID Contents")
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := echoContext.Request().Context()

	list, err := c.ContentUseCase.GetByID(ctx, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, list)
}

// CreateContent ...
func (c *ContentHandler) CreateContent(echoContext echo.Context) error {
	var content domain.Content
	err := echoContext.Bind(&content)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var ok bool
	if ok, err = util.IsRequestValid(&content); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.ContentUseCase.CreateContent(ctx, &content)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, content)

}

// UpdateContent ...
func (c *ContentHandler) UpdateContent(echoContext echo.Context) error {
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	var content domain.Content
	err = echoContext.Bind(&content)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var ok bool
	if ok, err = util.IsRequestValid(&content); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.ContentUseCase.UpdateContent(ctx, &content, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, content)

}
