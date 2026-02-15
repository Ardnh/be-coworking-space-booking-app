package handlers

import (
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	http "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/responses"
	validation_utils "github.com/Ardnh/be-coworking-space-booking-app/internal/utils/validator"
)

type UserHandlers struct {
	userService services.UserService
	validator   *validator.Validate
	log         *logrus.Logger
}

func NewUserHandlers(userService services.UserService, validator *validator.Validate, log *logrus.Logger) *UserHandlers {
	return &UserHandlers{
		userService: userService,
		validator:   validator,
		log:         log,
	}
}

func (h *UserHandlers) GetUserByToken(c *fiber.Ctx) error {
	userId := c.Locals("user_id")
	if userId == nil {
		return http.NewErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized", nil)
	}

	// cast ke string dengan aman
	userIdStr, ok := userId.(string)
	if !ok {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid user_id type", nil)
	}

	// parse UUID
	userIdUUID, err := uuid.Parse(userIdStr)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	result, err := h.userService.GetUserById(c.Context(), userIdUUID)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return http.NewSuccessResponse(
		c,
		fiber.StatusOK,
		"Successfully get user by token",
		result,
	)
}

func (h *UserHandlers) GetAllUser(c *fiber.Ctx) error {

	// 1. Parse query parameters dengan default values
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}
	if limit > 50 {
		limit = 50
	}

	// 2. Parse query parameters page
	var offset int
	pageStr := c.Query("page", "1")
	pageInt := 1
	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			pageInt = 1
		}
		offset = (pageInt - 1) * limit
	}

	sortBy := c.Query("sort_by", "created_at")
	sortOrder := c.Query("sort_order", "DESC")

	// Validate sortOrder
	sortOrder = strings.ToUpper(sortOrder)
	if sortOrder != "ASC" && sortOrder != "DESC" {
		sortOrder = "DESC"
	}

	// 3. Optional: Get additional filters
	username := c.Query("username", "")
	email := c.Query("email", "")
	fullname := c.Query("fullname", "")

	userParams := dto.GetUserParams{
		Username:  username,
		Email:     email,
		FullName:  fullname,
		Limit:     limit,
		Offset:    offset,
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}

	users, totalItems, err := h.userService.GetAllUsers(c.Context(), &userParams)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusNotFound, err.Error(), nil)
	}

	pagination := http.Pagination{
		CurrentPage: pageInt,
		PageSize:    limit,
		TotalItems:  totalItems,
		TotalPages:  totalItems / limit,
		HasNext:     totalItems > limit*pageInt,
		HasPrevious: pageInt > 1,
	}

	return http.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Successfully retrieved users", users, pagination)
}

func (h *UserHandlers) CreateUser(c *fiber.Ctx) error {

	var req dto.CreateUserDto
	if err := c.BodyParser(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if err := h.validator.Struct(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, validation_utils.FormatValidationErrors(err), nil)
	}

	err := h.userService.CreateUser(c.Context(), &req)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusConflict, err.Error(), nil)
	}

	return http.NewSuccessResponse(c, fiber.StatusCreated, "Successfully created user", nil)
}

func (h *UserHandlers) UpdateUser(c *fiber.Ctx) error {

	id := c.Params("id", "")
	if id == "" {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "ID is required", nil)
	}

	var req dto.UpdateUserDto
	if err := c.BodyParser(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if err := h.validator.Struct(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, validation_utils.FormatValidationErrors(err), nil)
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	errUpdate := h.userService.UpdateUser(c.Context(), parsedId, &req)
	if errUpdate != nil {
		return http.NewErrorResponse(c, fiber.StatusConflict, errUpdate.Error(), nil)
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully updated user", nil)
}

func (h *UserHandlers) DeleteUser(c *fiber.Ctx) error {

	id := c.Params("id", "")
	if id == "" {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "ID is required", nil)
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	errDelete := h.userService.DeleteUser(c.Context(), parsedId)
	if errDelete != nil {
		return http.NewErrorResponse(c, fiber.StatusInternalServerError, errDelete.Error(), nil)
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully deleted user", nil)
}
