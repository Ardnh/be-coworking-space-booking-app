package handlers

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	http "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type BlockedDateHandlers struct {
	blockedDateService services.BlockedDateService
	validator          *validator.Validate
	log                *logrus.Logger
}

func NewBlockedDateHandlers(blockedDateService services.BlockedDateService, validator *validator.Validate, log *logrus.Logger) *BlockedDateHandlers {
	return &BlockedDateHandlers{
		blockedDateService: blockedDateService,
		validator:          validator,
		log:                log,
	}
}

func (h *BlockedDateHandlers) GetBlockedDate(c *fiber.Ctx) error {

	// Get blocked date

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get blocked date", nil)
}

func (h *BlockedDateHandlers) GetBlockedDateById(c *fiber.Ctx) error {

	// Get by date and blocked date by id

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get blocked date by id", nil)
}

func (h *BlockedDateHandlers) CreateBlockedDate(c *fiber.Ctx) error {

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully create blocked date", nil)
}

func (h *BlockedDateHandlers) UpdateBlockedDate(c *fiber.Ctx) error {

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully update blocked date", nil)
}

func (h *BlockedDateHandlers) DeleteBlockedDate(c *fiber.Ctx) error {

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully delete blocked date", nil)
}
