package services

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
)

type ResourceServiceImpl struct {
	repo repositories.ResourceRepository
}

func NewResourceService(repo repositories.ResourceRepository) services.ResourceService {
	return &ResourceServiceImpl{
		repo: repo,
	}
}
