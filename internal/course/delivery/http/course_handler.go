package http

import (
	"net/http"
	"strconv"

	"strings"

	"github.com/labstack/echo/v4"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/util"
)

// ResponseError represents the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// CourseHandler ...
type CourseHandler struct {
	CourseUseCase domain.CourseUseCase
}

// NewCourseHandler ...
func NewCourseHandler(e *echo.Echo, us domain.CourseUseCase) {
	handler := &CourseHandler{
		CourseUseCase: us,
	}
	// Get Operation
	e.GET("/courses", handler.GetAll)
	e.GET("/courses/:id", handler.GetByID)
	e.GET("/courses/:id/stats", handler.GetByID)
	e.GET("/courses/:id/lessons", handler.GetByID)
	e.GET("/courses/:id/users", handler.GetByID)
	e.GET("/courses/:id/teams", handler.GetByID)

	// Create/Add Operation
	e.POST("/courses", handler.CreateCourse)
	e.POST("/courses/import", handler.GetByID)
	e.POST("/courses/:id/lessons", handler.GetByID)
	e.POST("/courses/:id/users", handler.GetByID)
	e.POST("/courses/:id/teams", handler.GetByID)

	// Update Operation
	e.PUT("/courses/:id", handler.UpdateCourse)
	e.PUT("/courses/:id/lessons/:id", handler.GetByID)
	e.PUT("/courses/actions", handler.GetByID)

	// Remove/Delete Operation
	e.DELETE("/courses/:id", handler.DeleteCourse)
	e.DELETE("/courses/:id/lessons/:id", handler.GetByID)
}

// GetAll godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags courses
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /courses [get]
func (c *CourseHandler) GetAll(echoContext echo.Context) error {
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

	listCourse, err := c.CourseUseCase.GetAll(ctx, start, limit)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, listCourse)
}

// GetByID ...
func (c *CourseHandler) GetByID(echoContext echo.Context) error {
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := echoContext.Request().Context()

	listCourse, err := c.CourseUseCase.GetByID(ctx, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, *listCourse)
}

// CreateCourse ...
func (c *CourseHandler) CreateCourse(echoContext echo.Context) error {
	var course domain.Course
	err := echoContext.Bind(&course)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var ok bool
	if ok, err = util.IsRequestValid(&course); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.CourseUseCase.CreateCourse(ctx, &course)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, course)

}

// UpdateCourse ...
func (c *CourseHandler) UpdateCourse(echoContext echo.Context) error {
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, err)
	}
	var course domain.Course
	err = echoContext.Bind(&course)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var ok bool
	if ok, err = util.IsRequestValid(&course); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.CourseUseCase.UpdateCourse(ctx, &course, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, course)

}

// DeleteCourse will delete course by given param
func (c *CourseHandler) DeleteCourse(echoContext echo.Context) error {
	idP, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := echoContext.Request().Context()

	err = c.CourseUseCase.DeleteCourse(ctx, id)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.NoContent(http.StatusNoContent)
}
