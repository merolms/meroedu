package http

import (
	"net/http"
	"strconv"

	"strings"

	"github.com/meroedu/meroedu/pkg/log"

	"github.com/labstack/echo/v4"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/util"
)

// ResponseError represents the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// CategoryHandler ...
type CategoryHandler struct {
	CategoryUseCase domain.CategoryUseCase
}

// NewCategroyHandler ...
func NewCategroyHandler(e *echo.Echo, us domain.CategoryUseCase) {
	handler := &CategoryHandler{
		CategoryUseCase: us,
	}
	// Get Operation
	e.GET("/categories", handler.GetAll)
	e.GET("/categories/:id", handler.GetByID)
	e.GET("/categories/:id/", handler.GetByID)

	// Create/Add Operation
	e.POST("/categories", handler.CreateCategory)

	// Update Operation
	e.PUT("/categories/:id", handler.GetByID)
	e.PUT("/categories/actions", handler.GetByID)

	// Remove/Delete Operation
	e.DELETE("/categories/:id", handler.GetByID)
}

// GetAll ...
func (c *CategoryHandler) GetAll(echoContext echo.Context) error {
	log.Info("Calling GetAll Categories")
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

	list, err := c.CategoryUseCase.GetAll(ctx, start, limit)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, list)
}

// GetByID ...
func (c *CategoryHandler) GetByID(echoContext echo.Context) error {
	log.Info("Calling GetByID Categories")
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := echoContext.Request().Context()

	list, err := c.CategoryUseCase.GetByID(ctx, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, list)
}

// CreateCategory ...
func (c *CategoryHandler) CreateCategory(echoContext echo.Context) error {
	var category domain.Category
	err := echoContext.Bind(&category)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var ok bool
	if ok, err = util.IsRequestValid(&category); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.CategoryUseCase.CreateCategory(ctx, &category)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, category)

}

// UpdateCategory ...
func (c *CategoryHandler) UpdateCategory(echoContext echo.Context) error {
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	var category domain.Category
	err = echoContext.Bind(&category)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var ok bool
	if ok, err = util.IsRequestValid(&category); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.CategoryUseCase.UpdateCategory(ctx, &category, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, category)

}
