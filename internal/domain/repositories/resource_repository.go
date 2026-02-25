package repositories

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
)

type ResourceRepository interface {
	CreateResource(ctx context.Context, resource *entities.Resource) (*entities.Resource, error)
}
