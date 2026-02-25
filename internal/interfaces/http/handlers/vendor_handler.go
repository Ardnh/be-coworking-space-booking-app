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

func (h *VendorHandlers) GetVendorResourcesByVendorId(c *fiber.Ctx) error {

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get vendor resources", nil)
}

func (h *VendorHandlers) GetAllVendors(c *fiber.Ctx) error {

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
	pageInt := 1
	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			pageInt = 1
		} else {
			pageInt = page
		}
	}

	sortBy := c.Query("sort_by", "created_at")
	sortOrder := c.Query("sort_order", "DESC")

	// Validate sortOrder
	sortOrder = strings.ToUpper(sortOrder)
	if sortOrder != "ASC" && sortOrder != "DESC" {
		sortOrder = "DESC"
	}

	// 3. Optional: Get additional filters
	name := c.Query("name", "")
	city := c.Query("city", "")
	minRatingStr := c.Query("min_rating", "")
	searchQuery := c.Query("search", "")

	var minRating float64
	if minRatingStr != "" {
		parsed, err := strconv.ParseFloat(minRatingStr, 64)
		if err != nil {
			return http.NewErrorResponse(c, fiber.StatusBadRequest, "min_rating must be a number", nil)
		}
		minRating = parsed
	}

	params := dto.VendorFilterRequest{
		Name:          name,
		City:          city,
		MinRating:     &minRating,
		SearchQuery:   searchQuery,
		Page:          pageInt,
		PageSize:      limit,
		SortBy:        sortBy,
		SortDirection: sortOrder,
	}

	result, totalItems, err := h.vendorService.GetAllVendors(c.Context(), params)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusNotFound, err.Error(), nil)
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

	return http.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Successfully retrieved vendor", result, pagination)
}

func (h *VendorHandlers) GetVendorByID(c *fiber.Ctx) error {

	id := c.Params("id", "")
	if id == "" {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "ID is required", nil)
	}

	vendorIdUUID, err := uuid.Parse(id)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	result, err := h.vendorService.GetVendorByID(c.Context(), vendorIdUUID)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get vendor", result)
}

func (h *VendorHandlers) GetVendorsResourcesByVendorID(c *fiber.Ctx) error {

	id := c.Params("vendor_id", "")
	if id == "" {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "ID is required", nil)
	}

	vendorIdUUID, err := uuid.Parse(id)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	result, err := h.vendorService.GetVendorsResourcesByVendorID(c.Context(), vendorIdUUID)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get vendor resources", result)
}

func (h *VendorHandlers) CreateVendor(c *fiber.Ctx) error {

	// ── 1. Parse field teks ──────────────────────────────────────────────────
	// String
	userOwnerId := strings.TrimSpace(c.FormValue("owner_user_id"))
	vendorName := strings.TrimSpace(c.FormValue("vendor_name"))
	address := strings.TrimSpace(c.FormValue("address"))
	city := strings.TrimSpace(c.FormValue("city"))
	phoneNumber := strings.TrimSpace(c.FormValue("phone_number"))
	email := strings.TrimSpace(c.FormValue("email"))
	description := strings.TrimSpace(c.FormValue("description"))
	profileImage, err := c.FormFile("image")

	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	req := dto.CreateVendorRequestDto{
		OwnerUserID: userOwnerId,
		VendorName:  vendorName,
		Address:     address,
		City:        city,
		PhoneNumber: phoneNumber,
		Email:       email,
		Description: description,
	}

	if err := h.validator.Struct(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "Failed to create vendor", validation_utils.FormatValidationErrors(err))
	}

	result, err := h.vendorService.CreateVendor(c.Context(), profileImage, &req)
	if err != nil {
		h.log.Error(err)
		return http.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to create vendor", nil)
	}

	return http.NewSuccessResponse(c, fiber.StatusCreated, "Successfully created vendor", result)
}

func (h *VendorHandlers) UpdateVendor(c *fiber.Ctx) error {

	id := c.Params("vendorId", "")
	if id == "" {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "ID is required", nil)
	}

	// ── 1. Parse field teks ──────────────────────────────────────────────────
	// String
	vendorName := strings.TrimSpace(c.FormValue("vendor_name"))
	address := strings.TrimSpace(c.FormValue("address"))
	city := strings.TrimSpace(c.FormValue("city"))
	phoneNumber := strings.TrimSpace(c.FormValue("phone_number"))
	email := strings.TrimSpace(c.FormValue("email"))
	description := strings.TrimSpace(c.FormValue("description"))
	profileImage, err := c.FormFile("image")

	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	req := dto.UpdateVendorRequestDto{
		VendorName:  &vendorName,
		Address:     &address,
		City:        &city,
		PhoneNumber: &phoneNumber,
		Email:       &email,
		Description: &description,
	}

	if err := h.validator.Struct(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, validation_utils.FormatValidationErrors(err), nil)
	}

	vendorIdUUID, err := uuid.Parse(id)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	result, err := h.vendorService.UpdateVendor(c.Context(), vendorIdUUID, profileImage, &req)
	if err != nil {
		return http.HandleError(c, err)
	}

	return http.NewSuccessResponse(c, fiber.StatusCreated, "Successfully update vendor", result)
}

func (h *VendorHandlers) DeleteVendor(c *fiber.Ctx) error {

	id := c.Params("vendorId", "")
	if id == "" {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "ID is required", nil)
	}

	vendorIdUUID, err := uuid.Parse(id)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	errDelete := h.vendorService.DeleteVendor(c.Context(), vendorIdUUID)
	if errDelete != nil {
		return http.HandleError(c, errDelete)
	}

	return http.NewSuccessResponse(c, fiber.StatusCreated, "Successfully delete vendor", nil)
}
