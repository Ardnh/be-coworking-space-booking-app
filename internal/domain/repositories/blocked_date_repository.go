package repositories

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type BlockedDateRepository interface {
	GetBlockedDateByResourceId(ctx context.Context, resourceId uuid.UUID) ([]*entities.BlockedDate, error)
	CreateBlockedDate(ctx context.Context, blockedDate *entities.BlockedDate) (*entities.BlockedDate, error)
	UpdateBlockedDate(ctx context.Context, blockedDate *entities.BlockedDate) (*entities.BlockedDate, error)
	DeleteBlockedDate(ctx context.Context, blockedDateId uuid.UUID) error
}
