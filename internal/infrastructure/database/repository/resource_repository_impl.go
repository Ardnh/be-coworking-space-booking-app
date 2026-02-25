package repository

import (
	"context"
	"fmt"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ResourceRespositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewResourceRespository(db *gorm.DB, redis *redis.Client) repositories.ResourceRepository {
	return &ResourceRespositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *ResourceRespositoryImpl) CreateResource(ctx context.Context, resource *entities.Resource) (*entities.Resource, error) {

	if err := r.db.WithContext(ctx).Model(&entities.Resource{}).Create(&resource).Error; err != nil {
		return nil, fmt.Errorf("gagal membuat resource: %w", err)
	}

	return resource, nil
}
