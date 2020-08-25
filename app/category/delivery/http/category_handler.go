package http

import (
	"fmt"
	"net/http"
	"strconv"

	"strings"

	"github.com/labstack/echo/v4"
	"github.com/meroedu/meroedu/app/domain"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
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
	fmt.Println("Calling GetAll Categories")
	ctx := echoContext.Request().Context()
	start, limit := 0, 10
	var err error
	for k, v := range echoContext.QueryParams() {
		switch k {
		case "start":
			val := strings.TrimSpace(v[0])
			if start, err = strconv.Atoi(val); err != nil {
				return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
			}
		case "limit":
			val := strings.TrimSpace(v[0])
			if limit, err = strconv.Atoi(val); err != nil {
				return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
			}
		}
	}

	list, err := c.CategoryUseCase.GetAll(ctx, start, limit)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, list)
}

// GetByID ...
func (c *CategoryHandler) GetByID(echoContext echo.Context) error {
	fmt.Println("Calling GetByID Categories")
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := echoContext.Request().Context()

	list, err := c.CategoryUseCase.GetByID(ctx, int64(idParam))
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
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
	if ok, err = isRequestValid(&category); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.CategoryUseCase.CreateCategory(ctx, &category)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
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
	if ok, err = isRequestValid(&category); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.CategoryUseCase.UpdateCategory(ctx, &category, int64(idParam))
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, category)

}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func isRequestValid(m *domain.Category) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
