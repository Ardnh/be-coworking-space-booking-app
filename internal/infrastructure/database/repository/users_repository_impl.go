package repository

import (
	"context"
	"errors"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUsersRepository(db *gorm.DB, redis *redis.Client) repositories.UsersRepository {
	return &userRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *userRepositoryImpl) GetAllUsers(ctx context.Context, usernameQuery string, emailQuery string, fullNameQuery string, limit int, offset int, sortBy string, sortOrder string) ([]*entities.Users, int, error) {

	var users []*entities.Users

	query := r.db.WithContext(ctx).Table("users u")

	// Search filter
	if usernameQuery != "" {
		searchPattern := "%" + usernameQuery + "%"
		query = query.Where("u.username LIKE ?", searchPattern)
	}
	if emailQuery != "" {
		searchPattern := "%" + emailQuery + "%"
		query = query.Where("u.email LIKE ?", searchPattern)
	}
	if fullNameQuery != "" {
		searchPattern := "%" + fullNameQuery + "%"
		query = query.Where("u.full_name LIKE ?", searchPattern)
	}

	// Pagination
	query = query.Limit(limit).Offset(offset)

	// Sorting
	if sortBy != "" {
		query = query.Order(sortBy + " " + sortOrder)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepositoryImpl) CreateUser(ctx context.Context, user *entities.Users) error {

	// Check if email already exists
	var existingUser entities.Users
	err := r.db.WithContext(ctx).
		Where("email = ?", user.Email).
		First(&existingUser).Error

	if err == nil {
		// Email found, return error
		return errors.New("email already registered")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Database error (bukan "not found")
		return err
	}

	// Email tidak ditemukan, proceed create
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryImpl) UpdateUser(ctx context.Context, user *entities.Users) error {

	// Check if user already exists
	var existingUser entities.Users
	err := r.db.WithContext(ctx).
		Where("user_id = ?", user.UserID).
		First(&existingUser).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err := r.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryImpl) DeleteUser(ctx context.Context, userID string) error {

	var existingUser entities.Users
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		First(&existingUser).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err := r.db.Delete(&entities.Users{}, userID).Error; err != nil {
		return err
	}

	return nil
}
