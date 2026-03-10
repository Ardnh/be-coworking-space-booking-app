package handlers

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	http "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type BookingHandler struct {
	service   services.BookingService
	validator *validator.Validate
	log       *logrus.Logger
}

func NewBookingHandler(service services.BookingService, validator *validator.Validate, log *logrus.Logger) *BookingHandler {
	return &BookingHandler{
		service:   service,
		validator: validator,
		log:       log,
	}
}

func (h *BookingHandler) CreateBooking(c *fiber.Ctx) error {

	return http.NewSuccessResponse(c, fiber.StatusOK, "OK", nil)
}

func (h *BookingHandler) CancelBooking(c *fiber.Ctx) error {

	bookingId := c.Params("bookingId", "")
	if bookingId == "" {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "bookingId is required", nil)
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Cancel OK", nil)
}
