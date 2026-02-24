package services

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
)

type BlockedDateServiceImpl struct {
	repo repositories.BlockedDateRepository
}

func NewBlockedDateService(repo repositories.BlockedDateRepository) services.BlockedDateService {
	return &BlockedDateServiceImpl{
		repo: repo,
	}
}
