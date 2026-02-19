package repository

import (
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
