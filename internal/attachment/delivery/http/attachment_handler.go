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
	AttachmentUseCase domain.AttachmentUseCase
}

// NewAttachmentHandler ...
func NewAttachmentHandler(e *echo.Echo, us domain.AttachmentUseCase) {
	handler := &AttachmentHandler{
		AttachmentUseCase: us,
	}
	// Get Operation
	e.POST("attachment/upload", handler.Upload)
	e.GET("attachment/download", handler.Download)
}

// Upload ...
func (a *AttachmentHandler) Upload(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
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
		File:     file,
		Filename: fileHeader.Filename,
		Size:     file.(util.Sizer).Size(),
		Type:     fileHeader.Header.Get("Content-Type"),
	}
	res, err := a.AttachmentUseCase.CreateAttachment(ctx, attachment)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, res)
}

// Download serves file
func (a *AttachmentHandler) Download(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()
	fileName := echoContext.QueryParam("fileName")
	filePath, err := a.AttachmentUseCase.DownloadAttachment(ctx, fileName)
	if err != nil {
		return echoContext.JSON(util.GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.File(filePath)
}
