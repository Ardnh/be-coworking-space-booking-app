package repositories

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type ResourceRepository interface {
	GetAllResources(ctx context.Context, resourceName string, resourceTypeId string, operationTimeStart string, operationTimeEnd string, sortBy string, sortOrder string, pageSize int, pageNumber int) ([]*entities.Resource, int, error)
	GetResourceByVendorId(ctx context.Context, vendorId uuid.UUID) ([]*entities.Resource, error)
	GetResourceById(ctx context.Context, resourceId uuid.UUID) (*entities.Resource, error)
	CreateResource(ctx context.Context, resource *entities.Resource) (*entities.Resource, error)
	UpdateResource(ctx context.Context, resource *entities.Resource) (*entities.Resource, error)
	DeleteResource(ctx context.Context, resourceId uuid.UUID) error
}
