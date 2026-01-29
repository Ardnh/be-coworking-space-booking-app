package handlers

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	http "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type VendorHandlers struct {
	vendorService services.VendorService
	validator     *validator.Validate
	log           *logrus.Logger
}

func NewVendorHandlers(vendorService services.VendorService, validator *validator.Validate, log *logrus.Logger) *VendorHandlers {
	return &VendorHandlers{
		vendorService: vendorService,
		validator:     validator,
		log:           log,
	}
}

func (s *VendorHandlers) GetAllVendors(c *fiber.Ctx) error {
	return http.NewSuccessResponse(c, fiber.StatusCreated, "Successfully created user", nil)
}

func (s *VendorHandlers) GetVendorByID(c *fiber.Ctx) error {
	return http.NewSuccessResponse(c, fiber.StatusCreated, "Successfully created user", nil)
}

func (s *VendorHandlers) GetVendorsResourcesByVendorID(c *fiber.Ctx) error {
	return http.NewSuccessResponse(c, fiber.StatusCreated, "Successfully created user", nil)
}

func (s *VendorHandlers) CreateVendor(c *fiber.Ctx) error {
	return http.NewSuccessResponse(c, fiber.StatusCreated, "Successfully created user", nil)
}

func (s *VendorHandlers) UpdateVendor(c *fiber.Ctx) error {
	return http.NewSuccessResponse(c, fiber.StatusCreated, "Successfully created user", nil)
}

func (s *VendorHandlers) DeleteVendor(c *fiber.Ctx) error {
	return http.NewSuccessResponse(c, fiber.StatusCreated, "Successfully created user", nil)
}
