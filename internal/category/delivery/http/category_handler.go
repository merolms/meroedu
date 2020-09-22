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

// CategoryHandler ...
type CategoryHandler struct {
	CategoryUseCase domain.CategoryUseCase
}

// NewCategoryHandler ...
func NewCategoryHandler(e *echo.Echo, us domain.CategoryUseCase) {
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
	e.PUT("/categories/:id", handler.UpdateCategory)
	e.PUT("/categories/actions", handler.GetByID)

	// Remove/Delete Operation
	e.DELETE("/categories/:id", handler.DeleteCategory)
}

// GetAll godoc
// @Summary Get All Categories summaries.
// @Description Get All Categories summaries..
// @Tags categories
// @Accept */*
// @Produce json
// @Param start query int true "start"
// @Param limit query int true "limit"
// @Success 200 {object} domain.Summaries
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /categories [get]
func (c *CategoryHandler) GetAll(echoContext echo.Context) error {
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

// GetByID godoc
// @Summary Get category by ID.
// @Description Get Specific category details.
// @Tags categories
// @Accept */*
// @Produce json
// @Param id path int true "category Id"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError "We need ID!!"
// @Failure 404 {object} domain.APIResponseError "Can not find ID"
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /categories/{id} [get]
func (c *CategoryHandler) GetByID(echoContext echo.Context) error {
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

// CreateCategory godoc
// @Summary Create New Category
// @Description Create New Category
// @Tags categories
// @Accept */*
// @Produce json
// @Param category body domain.Category true "Category Data"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError
// @Failure 404 {object} domain.APIResponseError
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /categories [post]
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

// UpdateCategory godoc
// @Summary Update existing category
// @Description Update existing category
// @Tags categories
// @Accept */*
// @Produce json
// @Param id path int true "Category Id"
// @Param Category body domain.Category true "Category Data"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError
// @Failure 404 {object} domain.APIResponseError
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /categories/{id} [put]
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
	return echoContext.JSON(http.StatusOK, category)

}

// DeleteCategory godoc
// @Summary Delete existing category
// @Description delete category by given parameter id
// @Tags categories
// @Accept */*
// @Produce json
// @Param id path int true "Category Id"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError
// @Failure 404 {object} domain.APIResponseError
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /categories/{id} [delete]
func (c *CategoryHandler) DeleteCategory(echoContext echo.Context) error {
	idP, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := echoContext.Request().Context()

	err = c.CategoryUseCase.DeleteCategory(ctx, id)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.NoContent(http.StatusNoContent)
}
