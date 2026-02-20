package handlers

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	http "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func (h *ResourceHandlers) GetReviewsByResourceId(c *fiber.Ctx) error {

	id := c.Params("resourceId", "")
	if id == "" {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "ID is required", nil)
	}

	resourceIdUUID, err := uuid.Parse(id)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get vendor", nil)
}

func (h *ResourceHandlers) GetResourceById(c *fiber.Ctx) error {
	id := c.Params("resourceId", "")
	if id == "" {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "ID is required", nil)
	}

	resourceIdUUID, err := uuid.Parse(id)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get vendor", nil)
}

func (h *ResourceHandlers) GetResource(c *fiber.Ctx) error {

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get vendor", nil)
}

func (h *ResourceHandlers) CreateResource(c *fiber.Ctx) error {

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get vendor", nil)
}

func (h *ResourceHandlers) UpdateResource(c *fiber.Ctx) error {

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get vendor", nil)
}

func (h *ResourceHandlers) DeleteResource(c *fiber.Ctx) error {

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get vendor", nil)
}
