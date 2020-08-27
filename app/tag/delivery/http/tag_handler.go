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

// TagHandler ...
type TagHandler struct {
	TagUseCase domain.TagUseCase
}

// NewTagHandler ...
func NewTagHandler(e *echo.Echo, us domain.TagUseCase) {
	handler := &TagHandler{
		TagUseCase: us,
	}
	// Get Operation
	e.GET("/tags", handler.GetAll)
	e.GET("/tags/:id", handler.GetByID)
	e.GET("/tags/:id/", handler.GetByID)

	// Create/Add Operation
	e.POST("/tags", handler.CreateTag)

	// Update Operation
	e.PUT("/tags/:id", handler.GetByID)
	e.PUT("/tags/actions", handler.GetByID)

	// Remove/Delete Operation
	e.DELETE("/tags/:id", handler.GetByID)
}

// GetAll ...
func (c *TagHandler) GetAll(echoContext echo.Context) error {
	fmt.Println("Calling GetAll Tags")
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

	list, err := c.TagUseCase.GetAll(ctx, start, limit)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, list)
}

// GetByID ...
func (c *TagHandler) GetByID(echoContext echo.Context) error {
	fmt.Println("Calling GetByID Tags")
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := echoContext.Request().Context()

	list, err := c.TagUseCase.GetByID(ctx, int64(idParam))
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, list)
}

// CreateTag ...
func (c *TagHandler) CreateTag(echoContext echo.Context) error {
	var tag domain.Tag
	err := echoContext.Bind(&tag)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var ok bool
	if ok, err = isRequestValid(&tag); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.TagUseCase.CreateTag(ctx, &tag)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, tag)

}

// UpdateTag ...
func (c *TagHandler) UpdateTag(echoContext echo.Context) error {
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	var tag domain.Tag
	err = echoContext.Bind(&tag)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var ok bool
	if ok, err = isRequestValid(&tag); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.TagUseCase.UpdateTag(ctx, &tag, int64(idParam))
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, tag)

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

func isRequestValid(m *domain.Tag) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
