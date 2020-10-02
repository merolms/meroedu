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
	e.GET("/contents/download", handler.DownloadContent)

	// Create/Add Operation
	e.POST("/contents", handler.CreateContent)

	// Update Operation
	e.PUT("/contents/:id", handler.GetByID)
	e.PUT("/contents/actions", handler.GetByID)

	// Remove/Delete Operation
	e.DELETE("/contents/:id", handler.GetByID)
}

// GetAll godoc
// @Summary Get All contents summaries.
// @Description Get All contents summaries..
// @Tags contents
// @Accept */*
// @Produce json
// @Param start query int true "start"
// @Param limit query int true "limit"
// @Success 200 {object} domain.Summaries
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /contents [get]
func (c *ContentHandler) GetAll(echoContext echo.Context) error {
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
	res := domain.Summaries{
		Response: domain.Response{
			Message: domain.Success,
			Data:    list,
		},
	}
	return echoContext.JSON(http.StatusOK, res)
}

// GetByID godoc
// @Summary Get Content by ID.
// @Description Get Specific Content details.
// @Tags contents
// @Accept */*
// @Produce json
// @Param id path int true "Content Id"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError "We need ID!!"
// @Failure 404 {object} domain.APIResponseError "Can not find ID"
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /contents/{id} [get]
func (c *ContentHandler) GetByID(echoContext echo.Context) error {
	idParam, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	ctx := echoContext.Request().Context()

	content, err := c.ContentUseCase.GetByID(ctx, int64(idParam))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	res := domain.Response{
		Data:    content,
		Message: domain.Success,
	}
	return echoContext.JSON(http.StatusOK, res)
}

// CreateContent godoc
// @Summary Create New Content
// @Description Create New Content
// @Tags contents
// @Accept */*
// @Produce json
// @Param Content body domain.Content true "Content Data"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError
// @Failure 404 {object} domain.APIResponseError
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /contents [post]
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

// UpdateContent godoc
// @Summary Update existing Content
// @Description Update existing Content
// @Tags contents
// @Accept */*
// @Produce json
// @Param id path int true "Content Id"
// @Param Content body domain.Content true "Content Data"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError
// @Failure 404 {object} domain.APIResponseError
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /contents/{id} [put]
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
	res := domain.Response{
		Data:    content,
		Message: domain.Success,
	}
	return echoContext.JSON(http.StatusOK, res)

}

// DeleteContent godoc
// @Summary Delete existing content
// @Description delete content by given parameter id
// @Tags contents
// @Accept */*
// @Produce json
// @Param id path int true "content Id"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.APIResponseError
// @Failure 404 {object} domain.APIResponseError
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /contents/{id} [delete]
func (c *ContentHandler) DeleteContent(echoContext echo.Context) error {
	idP, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := echoContext.Request().Context()

	err = c.ContentUseCase.DeleteContent(ctx, id)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.NoContent(http.StatusNoContent)
}

// DownloadContent godoc
// @Summary Download an attachment.
// @Description Download an attachment.
// @Tags contents
// @Accept */*
// @Param file query string true "uuid-encoded file name"
// @Produce json
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /contents/download [get]
func (c *ContentHandler) DownloadContent(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	fileName := echoContext.QueryParam("file")
	filePath, err := c.ContentUseCase.DownloadContent(ctx, fileName)
	if err != nil {
		log.Errorf("error while getting file path %v", err)
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.File(filePath)
}
