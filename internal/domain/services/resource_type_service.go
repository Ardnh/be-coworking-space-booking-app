package services

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/google/uuid"
)

type ResourceTypeService interface {
	GetAllResourceType(ctx context.Context, name string, limit int, offset int, sortBy string, sortOrder string) ([]*dto.ResourceTypeDto, int, error)
	GetResourceTypeById(ctx context.Context, resourceTypeID uuid.UUID) (*dto.ResourceTypeDto, error)
	CreateResourceType(ctx context.Context, resourceType *dto.CreateResourceTypeRequestDto) (*dto.ResourceTypeDto, error)
	UpdateResourceType(ctx context.Context, resourceTypeID uuid.UUID, resourceType *dto.UpdateResourceTypeRequestDto) (*dto.ResourceTypeDto, error)
	DeleteResourceType(ctx context.Context, resourceTypeID uuid.UUID) error
}
