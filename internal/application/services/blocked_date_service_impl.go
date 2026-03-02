package services

import (
	"context"
	"errors"
	"time"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/mapper"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type BlockedDateServiceImpl struct {
	repo repositories.BlockedDateRepository
	log  *logrus.Logger
}

func NewBlockedDateService(repo repositories.BlockedDateRepository, log *logrus.Logger) services.BlockedDateService {
	return &BlockedDateServiceImpl{
		repo: repo,
		log:  log,
	}
}

func (s *BlockedDateServiceImpl) GetBlockedDateByResourceId(ctx context.Context, resourceId uuid.UUID) ([]*dto.BlockedDate, error) {
	s.log.WithField("resource_id", resourceId).Info("fetching blocked dates by resource id")

	blockedDates, err := s.repo.GetBlockedDateByResourceId(ctx, resourceId)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"resource_id": resourceId,
			"error":       err,
		}).Error("failed to fetch blocked dates from repository")
		return nil, errors.New("terjadi kesalahan, coba lagi nanti")
	}

	blockedDateDTOs := mapper.ToBlockedDateListDTO(blockedDates)
	return blockedDateDTOs, nil
}

func (s *BlockedDateServiceImpl) CreateBlockedDate(ctx context.Context, blockedDateDto *dto.CreateBlockedDateRequest) (*dto.BlockedDate, error) {
	s.log.WithFields(logrus.Fields{
		"resource_id": blockedDateDto.ResourceID,
		"date":        blockedDateDto.Date,
	}).Info("creating blocked date")

	// Parse date dari request (bukan hardcoded)
	parsedDateTime, err := time.Parse(time.RFC3339, blockedDateDto.Date)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"date":  blockedDateDto.Date,
			"error": err,
		}).Warn("invalid date format")
		// User error — kembalikan pesan jelas
		return nil, errors.New("format tanggal tidak valid, gunakan format RFC3339")
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
		s.log.WithFields(logrus.Fields{
			"resource_id": blockedDateDto.ResourceID,
			"error":       err,
		}).Error("failed to create blocked date in repository")
		return nil, errors.New("terjadi kesalahan, coba lagi nanti")
	}

	s.log.WithFields(logrus.Fields{
		"block_id":    result.BlockID,
		"resource_id": blockedDateDto.ResourceID,
	}).Info("blocked date created successfully")

	return mapper.ToBlockedDateDTO(result), nil
}

func (s *BlockedDateServiceImpl) UpdateBlockedDate(ctx context.Context, blockId uuid.UUID, blockedDate *dto.UpdateBlockedDateRequest) (*dto.BlockedDate, error) {
	s.log.WithField("block_id", blockId).Info("updating blocked date")

	// ── 1. Ambil data blocked date yang sudah ada ──
	existingBlockedDate, err := s.repo.GetBlockedDateById(ctx, blockId)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"block_id": blockId,
			"error":    err,
		}).Error("failed to fetch blocked date from repository")
		return nil, errors.New("terjadi kesalahan, coba lagi nanti")
	}

	if existingBlockedDate == nil {
		s.log.WithField("block_id", blockId).Warn("blocked date not found")
		return nil, errors.New("blocked date tidak ditemukan")
	}

	// ── 2. TODO: Validasi konflik dengan booking yang sudah ada ──
	// 1. Ambil semua blocked date di tanggal yang sama dengan blockedDate.Date dan resourceId
	// 2. Ambil semua booking yang berpotensi bertabrakan dengan blockedDate di tanggal yang sama
	// 3. Jika ada booking yang bertabrakan, kembalikan error

	// ── 3. Terapkan partial update ──
	if blockedDate.Date != nil {
		parsedDateTime, err := time.Parse(time.RFC3339, *blockedDate.Date)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"date":  *blockedDate.Date,
				"error": err,
			}).Warn("invalid date format")
			// User error — kembalikan pesan jelas
			return nil, errors.New("format tanggal tidak valid, gunakan format RFC3339")
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

	// ── 4. Simpan ke DB ──
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
		s.log.WithFields(logrus.Fields{
			"block_id": blockId,
			"error":    err,
		}).Error("failed to update blocked date in repository")
		return nil, errors.New("terjadi kesalahan, coba lagi nanti")
	}

	s.log.WithField("block_id", blockId).Info("blocked date updated successfully")

	return mapper.ToBlockedDateDTO(result), nil
}

func (s *BlockedDateServiceImpl) DeleteBlockedDate(ctx context.Context, blockedDateId uuid.UUID) error {
	s.log.WithField("block_id", blockedDateId).Info("deleting blocked date")

	if err := s.repo.DeleteBlockedDate(ctx, blockedDateId); err != nil {
		s.log.WithFields(logrus.Fields{
			"block_id": blockedDateId,
			"error":    err,
		}).Error("failed to delete blocked date from repository")
		return errors.New("terjadi kesalahan, coba lagi nanti")
	}

	s.log.WithField("block_id", blockedDateId).Info("blocked date deleted successfully")

	return nil
}
