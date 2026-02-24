package repository

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
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
