package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/meroedu/course-api/app/domain"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
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
	e.GET("/courses", handler.GetAll)
	e.POST("/courses", handler.CreateCourse)
	e.GET("/courses/:id", handler.GetByID)

}

// GetAll ...
func (c *CourseHandler) GetAll(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	skip := 0
	limit := 10
	var err error
	fmt.Println(echoContext.QueryParams())
	if len(echoContext.QueryParams()) > 0 {
		if skip, err = strconv.Atoi(echoContext.QueryParam("skip")); err != nil {
			return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		}
		if limit, err = strconv.Atoi(echoContext.QueryParam("limit")); err != nil {
			return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		}
	}
	listCourse, err := c.CourseUseCase.GetAll(ctx, skip, limit)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	// echoContext.Response().Header().Set(`X-Cursor`, nextCursor)
	return echoContext.JSON(http.StatusOK, listCourse)
}

// GetByID ...
func (c *CourseHandler) GetByID(echoContext echo.Context) error {
	fmt.Println("Caling course handler get by id")
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := echoContext.Request().Context()

	listCourse, err := c.CourseUseCase.GetByID(ctx, int64(idParam))
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusOK, listCourse)
}

// CreateCourse ...
func (c *CourseHandler) CreateCourse(echoContext echo.Context) error {
	var course domain.Course
	err := echoContext.Bind(&course)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	var ok bool
	if ok, err = isRequestValid(&course); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := echoContext.Request().Context()
	err = c.CourseUseCase.CreateCourse(ctx, &course)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, course)

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

func isRequestValid(m *domain.Course) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
