package repositories

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type ResourceTypeRepository interface {
	GetAllResourceType(ctx context.Context, name string, limit int, offset int, sortBy string, sortOrder string) ([]*entities.ResourceType, int, error)
	GetResourceTypeById(ctx context.Context, resourceTypeID uuid.UUID) (*entities.ResourceType, error)
	CreateResourceType(ctx context.Context, resourceType *entities.ResourceType) (*entities.ResourceType, error)
	UpdateResourceType(ctx context.Context, resourceType *entities.ResourceType) (*entities.ResourceType, error)
	DeleteResourceType(ctx context.Context, resourceTypeID uuid.UUID) error
}
