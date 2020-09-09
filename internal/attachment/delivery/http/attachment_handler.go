package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/meroedu/meroedu/internal/domain"
	"github.com/meroedu/meroedu/internal/util"
)

// ResponseError represents the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// AttachmentHandler ...
type AttachmentHandler struct {
	AttachmentUseCase domain.AttachmentUserCase
}

// NewAttachmentHandler ...
func NewAttachmentHandler(e *echo.Echo, us domain.AttachmentUserCase) {
	handler := &AttachmentHandler{
		AttachmentUseCase: us,
	}
	// Create Attachment
	e.POST("attachments", handler.CreateAttachment)
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
		return err
	}
	file, err := fileHeader.Open()
	defer file.Close()
	if err != nil {
		return err
	}

	attachment := domain.Attachment{
		Title:       title,
		Description: description,
		File:        file,
		Filename:    fileHeader.Filename,
		Size:        file.(util.Sizer).Size(),
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
