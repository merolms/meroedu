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
	e.GET("/tags/course/:id", handler.GetCourseTags)
	e.GET("/tags/lesson/:id", handler.GetLessonTags)

	// Create/Add Operation
	e.POST("/tags", handler.CreateTag)
	e.POST("/tags/course/:course_id/:tag_id", handler.CreateCourseTag)
	e.POST("/tags/lesson/:lesson_id/:tag_id", handler.CreateLessonTag)

	// Update Operation
	e.PUT("/tags/:id", handler.UpdateTag)
	e.PUT("/tags/actions", handler.GetByID)

	// Remove/Delete Operation
	e.DELETE("/tags/:id", handler.DeleteTag)
	e.DELETE("/tags/course/:course_id/:tag_id", handler.DeleteCourseTag)
	e.DELETE("/tags/lesson/:lesson_id/:tag_id", handler.DeleteLessonTag)

}

// GetAll godoc
// @Summary Get All Tags summaries.
// @Description Get All Tags summaries..
// @Tags tags
// @Accept */*
// @Produce json
// @Param start query int true "start"
// @Param limit query int true "limit"
// @Success 200 {object} domain.Summaries
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /tags [get]
func (c *TagHandler) GetAll(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	start, limit := 0, 10
	searchQuery := echoContext.QueryParam("q")
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

	list, err := c.TagUseCase.GetAll(ctx, searchQuery, start, limit)
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
// @Summary Get tag by ID.
// @Description Get Specific tag details.
// @Tags tags
// @Accept */*
// @Produce json
// @Param id path int true "tag Id"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError "We need ID!!"
// @Failure 404 {object} domain.APIResponseError "Can not find ID"
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /tags/{id} [get]
func (c *TagHandler) GetByID(echoContext echo.Context) error {
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := echoContext.Request().Context()

	tag, err := c.TagUseCase.GetByID(ctx, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	res := domain.Response{
		Data:    tag,
		Message: domain.Success,
	}
	return echoContext.JSON(http.StatusOK, res)
}

// CreateTag godoc
// @Summary Create New tag
// @Description Create New tag
// @Tags tags
// @Accept */*
// @Produce json
// @Param tag body domain.Tag true "tag Data"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError
// @Failure 404 {object} domain.APIResponseError
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /tags [post]
func (c *TagHandler) CreateTag(echoContext echo.Context) error {
	var tag domain.Tag
	err := echoContext.Bind(&tag)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var ok bool
	if ok, err = util.IsRequestValid(&tag); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.TagUseCase.CreateTag(ctx, &tag)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	res := domain.Response{
		Data:    tag,
		Message: domain.Success,
	}
	return echoContext.JSON(http.StatusCreated, res)

}

// UpdateTag godoc
// @Summary Update existing tag
// @Description Update existing tag
// @Tags tags
// @Accept */*
// @Produce json
// @Param id path int true "tag Id"
// @Param tag body domain.Tag true "tag Data"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError
// @Failure 404 {object} domain.APIResponseError
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /tags/{id} [put]
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
	if ok, err = util.IsRequestValid(&tag); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.TagUseCase.UpdateTag(ctx, &tag, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	res := domain.Response{
		Data:    tag,
		Message: domain.Success,
	}
	return echoContext.JSON(http.StatusOK, res)
}

// DeleteTag godoc
// @Summary Delete existing tag
// @Description delete tag by given parameter id
// @Tags tags
// @Accept */*
// @Produce json
// @Param id path int true "Tag Id"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError
// @Failure 404 {object} domain.APIResponseError
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /tags/{id} [delete]
func (c *TagHandler) DeleteTag(echoContext echo.Context) error {
	idP, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := echoContext.Request().Context()

	err = c.TagUseCase.DeleteTag(ctx, id)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.NoContent(http.StatusNoContent)
}

// GetCourseTags godoc
// @Summary Get tags by CourseID.
// @Description Get all tags specific to course
// @Tags tags
// @Accept */*
// @Produce json
// @Param id path int true "course id"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError "We need ID!!"
// @Failure 404 {object} domain.APIResponseError "Can not find ID"
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /tags/course/{id} [get]
func (c *TagHandler) GetCourseTags(echoContext echo.Context) error {
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := echoContext.Request().Context()

	tags, err := c.TagUseCase.GetCourseTags(ctx, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	res := domain.Summaries{
		Response: domain.Response{
			Message: domain.Success,
			Data:    tags,
		},
	}
	return echoContext.JSON(http.StatusOK, res)
}

// CreateCourseTag godoc
// @Summary Create course tags
// @Description Create course tags
// @Tags tags
// @Accept */*
// @Produce json
// @Param course_id path int true "course id"
// @Param tag_id path int true "tag id"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError "We need ID!!"
// @Failure 404 {object} domain.APIResponseError "Can not find ID"
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /tags/course/{course_id}/{tag_id} [post]
func (c *TagHandler) CreateCourseTag(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	courseID, err := strconv.Atoi(echoContext.Param("course_id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	tagID, err := strconv.Atoi(echoContext.Param("tag_id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	err = c.TagUseCase.CreateCourseTag(ctx, int64(tagID), int64(courseID))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.NoContent(http.StatusCreated)
}

// DeleteCourseTag godoc
// @Summary Delete course tag by CourseID and tagID.
// @Description Delete course tag by CourseID and tagID.
// @Tags tags
// @Accept */*
// @Produce json
// @Param course_id path int true "course id"
// @Param tag_id path int true "tag id"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError "We need ID!!"
// @Failure 404 {object} domain.APIResponseError "Can not find ID"
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /tags/course/{course_id}/{tag_id} [delete]
func (c *TagHandler) DeleteCourseTag(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	courseID, err := strconv.Atoi(echoContext.Param("course_id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	tagID, err := strconv.Atoi(echoContext.Param("tag_id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	err = c.TagUseCase.DeleteCourseTag(ctx, int64(tagID), int64(courseID))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.NoContent(http.StatusOK)
}

// CreateLessonTag godoc
// @Summary Create Lesson Tag
// @Description Create Lesson Tag
// @Tags tags
// @Accept */*
// @Produce json
// @Param lesson_id path int true  "lesson id"
// @Param tag_id path int true  "tag id"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError "We need ID!!"
// @Failure 404 {object} domain.APIResponseError "Can not find ID"
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /tags/lesson/{lesson_id}/{tag_id} [post]
func (c *TagHandler) CreateLessonTag(echoContext echo.Context) error {
	lessonID, err := strconv.Atoi(echoContext.Param("lesson_id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	tagID, err := strconv.Atoi(echoContext.Param("tag_id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := echoContext.Request().Context()

	err = c.TagUseCase.CreateLessonTag(ctx, int64(tagID), int64(lessonID))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.NoContent(http.StatusCreated)
}

// DeleteLessonTag godoc
// @Summary Delete lesson tags by lessonID and tagID
// @Description Delete lesson tags by lessonID and tagID
// @Tags tags
// @Accept */*
// @Produce json
// @Param lesson_id path int true  "lesson id"
// @Param tag_id path int true  "tag id"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError "We need ID!!"
// @Failure 404 {object} domain.APIResponseError "Can not find ID"
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /tags/lesson/{lesson_id}/{tag_id} [delete]
func (c *TagHandler) DeleteLessonTag(echoContext echo.Context) error {
	lessonID, err := strconv.Atoi(echoContext.Param("lesson_id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	tagID, err := strconv.Atoi(echoContext.Param("tag_id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := echoContext.Request().Context()

	err = c.TagUseCase.DeleteLessonTag(ctx, int64(tagID), int64(lessonID))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.NoContent(http.StatusOK)
}

// GetLessonTags godoc
// @Summary Get tags by LessonID.
// @Description Get all tags specific to lesson
// @Tags tags
// @Accept */*
// @Produce json
// @Param id path int true "Lesson Id"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError "We need ID!!"
// @Failure 404 {object} domain.APIResponseError "Can not find ID"
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /tags/lesson/{id} [get]
func (c *TagHandler) GetLessonTags(echoContext echo.Context) error {
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := echoContext.Request().Context()

	tags, err := c.TagUseCase.GetLessonTags(ctx, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	res := domain.Summaries{
		Response: domain.Response{
			Message: domain.Success,
			Data:    tags,
		},
	}
	return echoContext.JSON(http.StatusOK, res)
}
