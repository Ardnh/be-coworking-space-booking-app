package services

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/google/uuid"
)

type BlockedDateService interface {
	GetBlockedDateByResourceId(ctx context.Context, resourceId uuid.UUID) ([]*dto.BlockedDate, error)
	CreateBlockedDate(ctx context.Context, blockedDateDto *dto.CreateBlockedDateRequest) (*dto.BlockedDate, error)
	UpdateBlockedDate(ctx context.Context, blockId uuid.UUID, blockedDate *dto.UpdateBlockedDateRequest) (*dto.BlockedDate, error)
	DeleteBlockedDate(ctx context.Context, blockedDateId uuid.UUID) error
}
