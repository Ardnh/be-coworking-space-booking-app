package handlers

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ResourceHandlers struct {
	vendorService services.ResourceService
	validator     *validator.Validate
	log           *logrus.Logger
}

func NewResourceHandler(vendorService services.ResourceService, validator *validator.Validate, log *logrus.Logger) *ResourceHandlers {
	return &ResourceHandlers{
		vendorService: vendorService,
		validator:     validator,
		log:           log,
	}
}
