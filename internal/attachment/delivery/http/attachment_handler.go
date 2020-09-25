package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/util"
	"github.com/meroedu/meroedu/pkg/log"
)

// ResponseError represents the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// AttachmentHandler ...
type AttachmentHandler struct {
	AttachmentUseCase domain.AttachmentUseCase
}

// NewAttachmentHandler ...
func NewAttachmentHandler(e *echo.Echo, us domain.AttachmentUseCase) {
	handler := &AttachmentHandler{
		AttachmentUseCase: us,
	}

	// Create Attachment
	e.POST("attachments", handler.CreateAttachment)
	// Download attachment
	e.GET("attachments/download", handler.DownloadAttachment)

}

// CreateAttachment godoc
// @Summary Create an attachment.
// @Description Create an attachment..
// @Tags attachments
// @Accept */*
// @Param title formData string false  "Title"
// @Param description formData string false  "Description"
// @Param file formData file true  "Upload file"
// @Produce json
// @Success 200 {object} domain.Response
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /attachments [post]
func (a *AttachmentHandler) CreateAttachment(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	title := echoContext.FormValue("title")
	description := echoContext.FormValue("description")
	fileHeader, err := echoContext.FormFile("file")
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	courseID, err := strconv.Atoi(echoContext.FormValue("course_id"))
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	file, err := fileHeader.Open()
	defer file.Close()
	if err != nil {
		return err
	}
	sizer, ok := file.(util.Sizer)
	if !ok {
		return echoContext.JSON(http.StatusBadRequest, ResponseError{Message: "invalid size"})
	}
	attachment := domain.Attachment{
		Title:       title,
		Description: description,
		CourseID:    int64(courseID),
		File:        file,
		Filename:    fileHeader.Filename,
		Size:        sizer.Size(),
		Type:        fileHeader.Header.Get("Content-Type"),
	}
	response, err := a.AttachmentUseCase.CreateAttachment(ctx, attachment)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	res := domain.Response{
		Data:    response,
		Message: domain.Success,
	}
	return echoContext.JSON(http.StatusCreated, res)
}

// DownloadAttachment godoc
// @Summary Download an attachment.
// @Description Download an attachment.
// @Tags attachments
// @Accept */*
// @Param file query string true "uuid-encoded file name"
// @Produce json
// @Failure 500 {object} domain.APIResponseError "Internal Server Error"
// @Router /attachments/download [get]
func (a *AttachmentHandler) DownloadAttachment(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	fileName := echoContext.QueryParam("file")
	filePath, err := a.AttachmentUseCase.DownloadAttachment(ctx, fileName)
	if err != nil {
		log.Errorf("error while getting file path %v", err)
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.File(filePath)
}
