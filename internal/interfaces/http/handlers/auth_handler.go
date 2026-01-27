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

type AuthHandlers struct {
	service     services.AuthService
	userService services.UserService
	validator   *validator.Validate
	log         *logrus.Logger
}

func NewAuthHandlers(authService services.AuthService, userService services.UserService, validator *validator.Validate, log *logrus.Logger) *AuthHandlers {
	return &AuthHandlers{
		service:     authService,
		userService: userService,
		validator:   validator,
		log:         log,
	}
}

func (h *AuthHandlers) Login(c *fiber.Ctx) error {

	var req dto.LoginRequestDto
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validator.Struct(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, validation_utils.FormatValidationErrors(err), nil)
	}

	token, expiresAt, err := h.service.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	response := map[string]interface{}{
		"token":      token,
		"expires_at": expiresAt,
	}

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully login user", response)
}

func (h *AuthHandlers) Register(c *fiber.Ctx) error {

	var req dto.CreateUserDto
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validator.Struct(&req); err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, validation_utils.FormatValidationErrors(err), nil)
	}

	err := h.userService.CreateUser(c.Context(), &req)
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return http.NewSuccessResponse(c, fiber.StatusCreated, "Successfully registered user", nil)
}
