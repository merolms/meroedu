package http

import (
	"fmt"
	"net/http"
	"strconv"

	"strings"

	"github.com/labstack/echo/v4"
	"github.com/meroedu/meroedu/app/domain"
	"github.com/meroedu/meroedu/app/util"
)

// ResponseError represents the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// LessonHandler ...
type LessonHandler struct {
	LessonUseCase domain.LessonUseCase
}

// NewLessonHandler ...
func NewLessonHandler(e *echo.Echo, us domain.LessonUseCase) {
	handler := &LessonHandler{
		LessonUseCase: us,
	}
	// Get Operation
	e.GET("/lessons", handler.GetAll)
	e.GET("/lessons/:id", handler.GetByID)
	e.GET("/lessons/:id/", handler.GetByID)

	// Create/Add Operation
	e.POST("/lessons", handler.CreateLesson)

	// Update Operation
	e.PUT("/lessons/:id", handler.GetByID)
	e.PUT("/lessons/actions", handler.GetByID)

	// Remove/Delete Operation
	e.DELETE("/lessons/:id", handler.GetByID)
}

// GetAll ...
func (c *LessonHandler) GetAll(echoContext echo.Context) error {
	fmt.Println("Calling GetAll Lessons")
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

	list, err := c.LessonUseCase.GetAll(ctx, start, limit)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, list)
}

// GetByID ...
func (c *LessonHandler) GetByID(echoContext echo.Context) error {
	fmt.Println("Calling GetByID Lessons")
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := echoContext.Request().Context()

	list, err := c.LessonUseCase.GetByID(ctx, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, list)
}

// CreateLesson ...
func (c *LessonHandler) CreateLesson(echoContext echo.Context) error {
	var lesson domain.Lesson
	err := echoContext.Bind(&lesson)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var ok bool
	if ok, err = util.IsRequestValid(&lesson); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.LessonUseCase.CreateLesson(ctx, &lesson)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, lesson)

}

// UpdateLesson ...
func (c *LessonHandler) UpdateLesson(echoContext echo.Context) error {
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	var lesson domain.Lesson
	err = echoContext.Bind(&lesson)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var ok bool
	if ok, err = util.IsRequestValid(&lesson); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.LessonUseCase.UpdateLesson(ctx, &lesson, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, lesson)

}
