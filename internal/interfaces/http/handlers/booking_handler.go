package handlers

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	http "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/responses"
	validation_utils "github.com/Ardnh/be-coworking-space-booking-app/internal/utils/validator"
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

	var req dto.CreateBookingRequestDto
	if err := c.BodyParser(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if err := h.validator.Struct(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, validation_utils.FormatValidationErrors(err), nil)
	}

	result, err := h.service.CreateBooking(c.Context(), &req)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully create booking", result)
}

// handler/booking_handler.go
func (h *BookingHandler) GetSlotAvailability(c *fiber.Ctx) error {
	var req dto.GetSlotAvailabilityRequestDto

	if err := c.ParamsParser(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid path params", nil)
	}
	if err := c.QueryParser(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid query params", nil)
	}
	if err := h.validator.Struct(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, validation_utils.FormatValidationErrors(err), nil)
	}

	// Langsung pakai req.ResourceID, req.Seats, dst.
	result, err := h.service.GetSlotAvailability(c.Context(), req)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get slot availability", result)
}

func (h *BookingHandler) CancelBooking(c *fiber.Ctx) error {

	bookingId := c.Params("bookingId", "")
	if bookingId == "" {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "bookingId is required", nil)
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Cancel OK", nil)
}
