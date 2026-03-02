package repository

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type BookingRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewBookingRepostory(db *gorm.DB, redis *redis.Client) repositories.BookingRespository {
	return &BookingRepositoryImpl{
		db:    db,
		redis: redis,
	}
}
