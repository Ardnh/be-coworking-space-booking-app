package handlers

import (
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	http "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/responses"
)

type UserHandlers struct {
	userService services.UserService
	validator   *validator.Validate
}

func NewAuthHandlers(userService services.UserService, validator *validator.Validate) *UserHandlers {
	return &UserHandlers{
		userService: userService,
		validator:   validator,
	}
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
		HasNext:     true,
		HasPrevious: false,
	}

	return http.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Successfully retrieved users", users, pagination)
}

func (h *UserHandlers) CreateUser(c *fiber.Ctx) error {

}

func (h *UserHandlers) UpdateUser(c *fiber.Ctx) error {

}

func (h *UserHandlers) DeleteUser(c *fiber.Ctx) error {

}
