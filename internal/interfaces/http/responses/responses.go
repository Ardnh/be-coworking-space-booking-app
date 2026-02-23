package responses

import (
	"errors"

	errorsConst "github.com/Ardnh/be-coworking-space-booking-app/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success bool `json:"success"`
	Message any  `json:"message"`
	Data    any  `json:"data,omitempty"`
	Error   any  `json:"error,omitempty"`
}

type ResponseWithPagination struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
	Error      any    `json:"error,omitempty"`
	Pagination any    `json:"pagination,omitempty"`
}

type Pagination struct {
	CurrentPage int  `json:"current_page"`
	PageSize    int  `json:"page_size"`
	TotalItems  int  `json:"total_items"`
	TotalPages  int  `json:"total_pages"`
	HasNext     bool `json:"has_next"`
	HasPrevious bool `json:"has_previous"`
}

// Success response
func NewSuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func NewSuccessResponseWithPagination(c *fiber.Ctx, statusCode int, message string, data interface{}, pagination interface{}) error {

	return c.Status(statusCode).JSON(ResponseWithPagination{
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	})
}

// Error response
func NewErrorResponse(c *fiber.Ctx, statusCode int, message any, err any) error {

	return c.Status(statusCode).JSON(Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// Centralized error handler
func HandleError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, errorsConst.ErrNotFound):
		return NewErrorResponse(c, fiber.StatusNotFound, err.Error(), nil)
	case errors.Is(err, errorsConst.ErrConflict):
		return NewErrorResponse(c, fiber.StatusConflict, err.Error(), nil)
	case errors.Is(err, errorsConst.ErrInvalidInput):
		return NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	case errors.Is(err, errorsConst.ErrUnauthorized):
		return NewErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	case errors.Is(err, errorsConst.ErrForbidden):
		return NewErrorResponse(c, fiber.StatusForbidden, err.Error(), nil)
	default:
		return NewErrorResponse(c, fiber.StatusInternalServerError, "Internal server error", nil)
	}
}
