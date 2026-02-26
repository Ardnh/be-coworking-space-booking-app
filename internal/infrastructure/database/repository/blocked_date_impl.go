package repository

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BlockedDateRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewBlockedDateRepository(db *gorm.DB, redis *redis.Client) repositories.BlockedDateRepository {
	return &BlockedDateRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *BlockedDateRepositoryImpl) GetBlockedDateById(ctx context.Context, blockedDateId uuid.UUID) (*entities.BlockedDate, error) {

	var blockedDate *entities.BlockedDate
	if err := r.db.WithContext(ctx).Model(&entities.BlockedDate{}).Where("block_id = ?", blockedDateId).First(&blockedDate).Error; err != nil {
		return nil, err
	}

	return blockedDate, nil
}

func (r *BlockedDateRepositoryImpl) GetBlockedDateByResourceId(ctx context.Context, resourceId uuid.UUID) ([]*entities.BlockedDate, error) {

	var blockedDateList []*entities.BlockedDate
	if err := r.db.WithContext(ctx).Model(&entities.BlockedDate{}).Where("resource_id = ?", resourceId).Find(&blockedDateList).Error; err != nil {
		return nil, err
	}

	return blockedDateList, nil
}

func (r *BlockedDateRepositoryImpl) CreateBlockedDate(ctx context.Context, blockedDate *entities.BlockedDate) (*entities.BlockedDate, error) {

	if err := r.db.WithContext(ctx).Model(&entities.BlockedDate{}).Create(&blockedDate).Error; err != nil {
		return nil, err
	}

	return blockedDate, nil
}

func (r *BlockedDateRepositoryImpl) UpdateBlockedDate(ctx context.Context, blockedDate *entities.BlockedDate) (*entities.BlockedDate, error) {

	errUpdate := r.db.WithContext(ctx).Model(&entities.BlockedDate{}).Where("block_id = ?", blockedDate.BlockID).Updates(blockedDate).Error
	if errUpdate != nil {
		return nil, errUpdate
	}

	var updatedBlockedDate entities.BlockedDate
	if err := r.db.WithContext(ctx).Model(&entities.BlockedDate{}).Where("block_id = ?", blockedDate.BlockID).First(&updatedBlockedDate).Error; err != nil {
		return nil, err
	}

	return &updatedBlockedDate, nil
}

func (r *BlockedDateRepositoryImpl) DeleteBlockedDate(ctx context.Context, blockedDateId uuid.UUID) error {

	err := r.db.WithContext(ctx).Where("block_id = ?", blockedDateId).Delete(&entities.BlockedDate{}).Error
	if err != nil {
		return err
	}

	return nil
}
