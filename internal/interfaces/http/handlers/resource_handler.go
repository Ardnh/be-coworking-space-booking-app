package handlers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	http "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/responses"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/utils/string_utils"
	validation_utils "github.com/Ardnh/be-coworking-space-booking-app/internal/utils/validator"
	"github.com/Ardnh/be-coworking-space-booking-app/pkg/constants"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ResourceHandlers struct {
	service   services.ResourceService
	validator *validator.Validate
	log       *logrus.Logger
}

func NewResourceHandler(vendorService services.ResourceService, validator *validator.Validate, log *logrus.Logger) *ResourceHandlers {
	return &ResourceHandlers{
		service:   vendorService,
		validator: validator,
		log:       log,
	}
}

func (h *ResourceHandlers) GetReviewsByResourceId(c *fiber.Ctx) error {

	// id := c.Params("resourceId", "")
	// if id == "" {
	// 	return http.NewErrorResponse(c, fiber.StatusBadRequest, "ID is required", nil)
	// }

	// resourceIdUUID, err := uuid.Parse(id)
	// if err != nil {
	// 	return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	// }

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get vendor", nil)
}

func (h *ResourceHandlers) GetResourceById(c *fiber.Ctx) error {
	// id := c.Params("resourceId", "")
	// if id == "" {
	// 	return http.NewErrorResponse(c, fiber.StatusBadRequest, "ID is required", nil)
	// }

	// resourceIdUUID, err := uuid.Parse(id)
	// if err != nil {
	// 	return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	// }

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get vendor", nil)
}

func (h *ResourceHandlers) GetResource(c *fiber.Ctx) error {

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get vendor", nil)
}

func (h *ResourceHandlers) CreateResource(c *fiber.Ctx) error {

	// ── 1. Parse field teks ──────────────────────────────────────────────────
	// String
	vendorId := strings.TrimSpace(c.FormValue("vendor_id"))
	resourceTypeId := strings.TrimSpace(c.FormValue("resource_type_id"))
	resourceName := strings.TrimSpace(c.FormValue("resource_name"))
	location := strings.TrimSpace(c.FormValue("location"))
	description := strings.TrimSpace(c.FormValue("description"))
	operationTimeFrom := strings.TrimSpace(c.FormValue("operation_time_from"))
	operationTimeTo := strings.TrimSpace(c.FormValue("operation_time_to"))
	endDate := strings.TrimSpace(c.FormValue("end_date"))

	// Number
	capacityStr := strings.TrimSpace(c.FormValue("capacity"))
	pricePerUnitStr := strings.TrimSpace(c.FormValue("price_per_unit"))

	pricePerUnit, err := strconv.ParseFloat(pricePerUnitStr, 64)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	capacity, err := strconv.Atoi(capacityStr)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	// Array of blocked Date
	blockedDateStr := strings.TrimSpace(c.FormValue("blocked_date"))
	var blockedDates []*dto.CreateBlockedDateRequest
	errParseBlockedDate := json.Unmarshal([]byte(blockedDateStr), &blockedDates)

	if errParseBlockedDate != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	// Validasi field wajib

	req := dto.CreateResourceRequestDto{
		VendorID:          vendorId,
		ResourceTypeID:    resourceTypeId,
		ResourceName:      resourceName,
		Location:          &location,
		Description:       &description,
		OperationTimeFrom: operationTimeFrom,
		OperationTimeTo:   operationTimeTo,
		EndDate:           endDate,
		Capacity:          capacity,
		PricePerUnit:      pricePerUnit,
		Status:            constants.ResourceStatusActive,
		BlockedDates:      blockedDates,
	}

	if err := h.validator.Struct(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "Failed to create vendor", validation_utils.FormatValidationErrors(err))
	}

	// ── 2. Parse multiple file gambar ────────────────────────────────────────
	form, err := c.MultipartForm()
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	files := form.File["images"]
	if len(files) == 0 {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "At least one image is required", nil)
	}

	result, errCreate := h.service.CreateResource(c.Context(), &req, files)
	if errCreate != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, errCreate.Error(), nil)
	}

	if errCreate != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, errCreate.Error(), nil)
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully create resource", result)
}

func (h *ResourceHandlers) UpdateResource(c *fiber.Ctx) error {

	resourceId := c.Params("resourceId", "")
	if resourceId == "" {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "Resource Id is required", nil)
	}

	// ── 1. Parse field teks ──────────────────────────────────────────────────
	// Number
	capacityStr := string_utils.ToStringPtr(c.FormValue("capacity"))
	pricePerUnitStr := string_utils.ToStringPtr(c.FormValue("price_per_unit"))

	capacity := 1
	if capacityStr != nil {
		capacityParsed, err := strconv.Atoi(*capacityStr)
		if err != nil {
			return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
		}

		capacity = capacityParsed
	}

	pricePerUnit := 0.0
	if pricePerUnitStr != nil {
		pricePerUnitParsed, err := strconv.ParseFloat(*pricePerUnitStr, 64)
		if err != nil {
			return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
		}

		pricePerUnit = pricePerUnitParsed
	}

	// Array of blocked Date
	blockedDateStr := strings.TrimSpace(c.FormValue("blocked_date"))
	var blockedDates []*dto.CreateBlockedDateRequest
	errParseBlockedDate := json.Unmarshal([]byte(blockedDateStr), &blockedDates)

	if errParseBlockedDate != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, errParseBlockedDate.Error(), nil)
	}

	req := dto.UpdateResourceRequestDto{
		ResourceName:      string_utils.ToStringPtr(c.FormValue("vendor_name")),
		Location:          string_utils.ToStringPtr(c.FormValue("location")),
		Description:       string_utils.ToStringPtr(c.FormValue("description")),
		OperationTimeFrom: string_utils.ToStringPtr(c.FormValue("operation_time_from")),
		OperationTimeTo:   string_utils.ToStringPtr(c.FormValue("operation_time_to")),
		EndDate:           string_utils.ToStringPtr(c.FormValue("end_date")),
		Capacity:          &capacity,
		PricePerUnit:      &pricePerUnit,
		BlockedDates:      blockedDates,
	}

	if err := h.validator.Struct(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "Failed to create vendor", validation_utils.FormatValidationErrors(err))
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully update resource", nil)
}

func (h *ResourceHandlers) DeleteResource(c *fiber.Ctx) error {

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully delete resource", nil)
}
