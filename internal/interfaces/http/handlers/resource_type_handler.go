package handlers

import (
	"strconv"
	"strings"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	http "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/responses"
	validation_utils "github.com/Ardnh/be-coworking-space-booking-app/internal/utils/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ResourceTypeHandlers struct {
	resourceTypeService services.ResourceTypeService
	validator           *validator.Validate
	log                 *logrus.Logger
}

func NewResourceTypeHandlers(resourceTypeService services.ResourceTypeService, validator *validator.Validate, log *logrus.Logger) *ResourceTypeHandlers {
	return &ResourceTypeHandlers{
		resourceTypeService: resourceTypeService,
		validator:           validator,
		log:                 log,
	}
}

func (h *ResourceTypeHandlers) GetAllResourceType(c *fiber.Ctx) error {

	// 1. Parse query parameters dengan default values
	limit, err := strconv.Atoi(c.Query("pageSize", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}
	if limit > 50 {
		limit = 50
	}

	// 2. Parse query parameters page
	pageStr := c.Query("page", "1")
	pageInt := 0
	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			pageInt = 1
		} else {
			pageInt = page
		}
	}

	offset := (pageInt - 1) * limit
	sortBy := c.Query("sort_by", "created_at")
	sortOrder := c.Query("sort_order", "DESC")

	// Validate sortOrder
	sortOrder = strings.ToUpper(sortOrder)
	if sortOrder != "ASC" && sortOrder != "DESC" {
		sortOrder = "DESC"
	}

	// 3. Optional: Get additional filters
	searchQuery := c.Query("name", "")

	result, totalItems, err := h.resourceTypeService.GetAllResourceType(c.Context(), searchQuery, limit, offset, sortBy, sortOrder)
	if err != nil {
		return http.HandleError(c, err)
	}

	totalPages := 0
	if totalItems > 0 {
		totalPages = (totalItems + limit - 1) / limit
	}

	pagination := http.Pagination{
		CurrentPage: pageInt,
		PageSize:    limit,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		HasNext:     pageInt < totalPages,
		HasPrevious: pageInt > 1,
	}

	return http.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Successfully retrieved get all resources", result, pagination)
}

func (h *ResourceTypeHandlers) GetResourceTypeById(c *fiber.Ctx) error {

	id := c.Params("resourceTypeId", "")
	if id == "" {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "ID is required", nil)
	}

	resourceTypeIdUUID, err := uuid.Parse(id)
	if err != nil {
		return http.HandleError(c, err)
	}
	result, err := h.resourceTypeService.GetResourceTypeById(c.Context(), resourceTypeIdUUID)
	if err != nil {
		return http.HandleError(c, err)
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get resource type by id", result)
}

func (h *ResourceTypeHandlers) CreateResourceType(c *fiber.Ctx) error {

	var req dto.CreateResourceTypeRequestDto
	if err := c.BodyParser(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "Failed to create resource type", err.Error())
	}

	if err := h.validator.Struct(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "Failed to create resource type", validation_utils.FormatValidationErrors(err))
	}

	result, err := h.resourceTypeService.CreateResourceType(c.Context(), &req)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully created resource type", result)
}

func (h *ResourceTypeHandlers) UpdateResourceType(c *fiber.Ctx) error {

	id := c.Params("resourceTypeId", "")
	if id == "" {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "ID is required", nil)
	}

	var req dto.UpdateResourceTypeRequestDto
	if err := c.BodyParser(&req); err != nil {
		return http.HandleError(c, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, validation_utils.FormatValidationErrors(err), nil)
	}

	resourceTypeIdUUID, err := uuid.Parse(id)
	if err != nil {
		return http.HandleError(c, err)
	}

	result, err := h.resourceTypeService.UpdateResourceType(c.Context(), resourceTypeIdUUID, &req)
	if err != nil {
		return http.HandleError(c, err)
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully updated resource type", result)
}

func (h *ResourceTypeHandlers) DeleteResourceType(c *fiber.Ctx) error {

	id := c.Params("resourceTypeId", "")
	if id == "" {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "ID is required", nil)
	}

	resourceTypeIdUUID, err := uuid.Parse(id)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	errDelete := h.resourceTypeService.DeleteResourceType(c.Context(), resourceTypeIdUUID)
	if errDelete != nil {
		return http.HandleError(c, errDelete)
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully deleted resource type", nil)
}
