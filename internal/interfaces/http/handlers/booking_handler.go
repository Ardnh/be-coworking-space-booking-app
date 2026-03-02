package handlers

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	"github.com/go-playground/validator/v10"
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
