package services

import (
	"context"
	"time"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/mapper"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	"github.com/google/uuid"
)

type BlockedDateServiceImpl struct {
	repo repositories.BlockedDateRepository
}

func NewBlockedDateService(repo repositories.BlockedDateRepository) services.BlockedDateService {
	return &BlockedDateServiceImpl{
		repo: repo,
	}
}

func (s *BlockedDateServiceImpl) GetBlockedDateByResourceId(ctx context.Context, resourceId uuid.UUID) ([]*dto.BlockedDate, error) {

	blockedDates, err := s.repo.GetBlockedDateByResourceId(ctx, resourceId)
	if err != nil {
		return nil, err
	}

	blockedDateDTOs := mapper.ToBlockedDateListDTO(blockedDates)
	return blockedDateDTOs, nil
}

func (s *BlockedDateServiceImpl) CreateBlockedDate(ctx context.Context, blockedDateDto *dto.CreateBlockedDateRequest) (*dto.BlockedDate, error) {

	parsedDateTime, err := time.Parse(time.RFC3339, "2025-01-15T08:30:00Z")
	if err != nil {
		return nil, err
	}

	blockedDateEntities := entities.BlockedDate{
		ResourceID: blockedDateDto.ResourceID,
		Date:       parsedDateTime,
		TimeFrom:   blockedDateDto.TimeFrom,
		TimeTo:     blockedDateDto.TimeTo,
		Reason:     blockedDateDto.Reason,
		CreatedBy:  blockedDateDto.CreatedBy,
		CreatedAt:  time.Now(),
	}

	result, err := s.repo.CreateBlockedDate(ctx, &blockedDateEntities)
	if err != nil {
		return nil, err
	}

	return mapper.ToBlockedDateDTO(result), nil
}

func (s *BlockedDateServiceImpl) UpdateBlockedDate(ctx context.Context, blockId uuid.UUID, blockedDate *dto.UpdateBlockedDateRequest) (*dto.BlockedDate, error) {

	var existingBlockedDate *entities.BlockedDate
	existingBlockedDate, err := s.repo.GetBlockedDateById(ctx, blockId)
	if err != nil {
		return nil, err
	}

	// Tambah pengecekan jika di tanggal dan waktu from to nya
	// sudah ada booking maka update Blocked time ditolak
	// 1. Ambil semua blocked date di tanggal yang sama dengan blockedDate.Date dan resourceId
	// 2. Ambil semua booking yang berpotensi bertabrakan dengan blockedDate di tanggal yang sama dengan blockedDate.Date dan resourceId
	// 3. Jika ada booking yang bertabrakan, kembalikan error

	if blockedDate.Date != nil {
		parsedDateTime, err := time.Parse(time.RFC3339, *blockedDate.Date)
		if err != nil {
			return nil, err
		}

		existingBlockedDate.Date = parsedDateTime
	}

	if blockedDate.TimeFrom != nil {
		existingBlockedDate.TimeFrom = *blockedDate.TimeFrom
	}

	if blockedDate.TimeTo != nil {
		existingBlockedDate.TimeTo = *blockedDate.TimeTo
	}

	if blockedDate.Reason != nil {
		existingBlockedDate.Reason = *blockedDate.Reason
	}

	blockedDateEntities := entities.BlockedDate{
		BlockID:   blockId,
		Date:      existingBlockedDate.Date,
		TimeFrom:  existingBlockedDate.TimeFrom,
		TimeTo:    existingBlockedDate.TimeTo,
		Reason:    existingBlockedDate.Reason,
		UpdatedAt: time.Now(),
	}

	result, err := s.repo.UpdateBlockedDate(ctx, &blockedDateEntities)
	if err != nil {
		return nil, err
	}

	return mapper.ToBlockedDateDTO(result), nil
}

func (s *BlockedDateServiceImpl) DeleteBlockedDate(ctx context.Context, blockedDateId uuid.UUID) error {

	if err := s.repo.DeleteBlockedDate(ctx, blockedDateId); err != nil {
		return err
	}

	return nil
}
