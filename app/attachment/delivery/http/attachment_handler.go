package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/meroedu/meroedu/app/domain"
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
	// Get Operation
	e.POST("attachment/upload", handler.Upload)
}

// Upload ...
func (a *AttachmentHandler) Upload(echoContext echo.Context) error {
	fmt.Println("Calling Upload file")
	err := a.AttachmentUseCase.Upload(echoContext)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, "attachment saved")
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
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
