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
	e.PUT("/lessons/:id", handler.UpdateLesson)
	e.PUT("/lessons/actions", handler.GetByID)

	// Remove/Delete Operation
	e.DELETE("/lessons/:id", handler.DeleteLesson)
}

// GetAll godoc
// @Summary Get All lessons summaries.
// @Description Get All lessons summaries..
// @Tags lessons
// @Accept */*
// @Produce json
// @Param start query int true "start"
// @Param limit query int true "limit"
// @Success 200 {object} domain.Summaries
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /lessons [get]
func (c *LessonHandler) GetAll(echoContext echo.Context) error {
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
	res := domain.Summaries{
		Response: domain.Response{
			Message: domain.Success,
			Data:    list,
		},
	}
	return echoContext.JSON(http.StatusOK, res)
}

// GetByID godoc
// @Summary Get Lesson by ID.
// @Description Get Specific Lesson details.
// @Tags lessons
// @Accept */*
// @Produce json
// @Param id path int true "Lesson Id"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError "We need ID!!"
// @Failure 404 {object} domain.APIResponseError "Can not find ID"
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /lessons/{id} [get]
func (c *LessonHandler) GetByID(echoContext echo.Context) error {
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := echoContext.Request().Context()

	lesson, err := c.LessonUseCase.GetByID(ctx, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	res := domain.Response{
		Data:    lesson,
		Message: domain.Success,
	}
	return echoContext.JSON(http.StatusOK, res)
}

// CreateLesson godoc
// @Summary Create New Lesson
// @Description Create New Lesson
// @Tags lessons
// @Accept */*
// @Produce json
// @Param Lesson body domain.Lesson true "Lesson Data"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError
// @Failure 404 {object} domain.APIResponseError
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /lessons [post]
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
	res := domain.Response{
		Data:    lesson,
		Message: domain.Success,
	}
	return echoContext.JSON(http.StatusCreated, res)

}

// UpdateLesson godoc
// @Summary Update existing Lesson
// @Description Update existing Lesson
// @Tags lessons
// @Accept */*
// @Produce json
// @Param id path int true "Lesson Id"
// @Param Lesson body domain.Lesson true "Lesson Data"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError
// @Failure 404 {object} domain.APIResponseError
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /lessons/{id} [put]
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
	res := domain.Response{
		Data:    lesson,
		Message: domain.Success,
	}
	return echoContext.JSON(http.StatusOK, res)

}

// DeleteLesson godoc
// @Summary Delete existing Lesson
// @Description delete Lesson by given parameter id
// @Tags lessons
// @Accept */*
// @Produce json
// @Param id path int true "Lesson Id"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError
// @Failure 404 {object} domain.APIResponseError
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /lessons/{id} [delete]
func (c *LessonHandler) DeleteLesson(echoContext echo.Context) error {
	idP, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := echoContext.Request().Context()

	err = c.LessonUseCase.DeleteLesson(ctx, id)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.NoContent(http.StatusNoContent)
}
