package repository

import (
	"context"
	"errors"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type authRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewAuthRepository(db *gorm.DB, redis *redis.Client) repositories.AuthRepository {
	return &authRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *authRepositoryImpl) Login(ctx context.Context, email string, password string) (*entities.Users, error) {

	var user entities.Users
	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Database error (bukan "not found")
		return nil, err
	}

	return &user, nil
}
