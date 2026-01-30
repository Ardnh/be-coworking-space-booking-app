package repository

import (
	"context"
	"errors"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/google/uuid"

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

func (r *userRepositoryImpl) GetAllUsers(
	ctx context.Context,
	usernameQuery string,
	emailQuery string,
	fullNameQuery string,
	limit int,
	offset int,
	sortBy string,
	sortOrder string,
) ([]*entities.Users, int, error) {

	var users []*entities.Users
	var total int64

	baseQuery := r.db.WithContext(ctx).Table("users u")

	// Search filter
	if usernameQuery != "" {
		baseQuery = baseQuery.Where("u.username LIKE ?", "%"+usernameQuery+"%")
	}
	if emailQuery != "" {
		baseQuery = baseQuery.Where("u.email LIKE ?", "%"+emailQuery+"%")
	}
	if fullNameQuery != "" {
		baseQuery = baseQuery.Where("u.full_name LIKE ?", "%"+fullNameQuery+"%")
	}

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Sorting (whitelist)
	allowedSort := map[string]bool{
		"username":   true,
		"email":      true,
		"full_name":  true,
		"created_at": true,
	}

	if allowedSort[sortBy] {
		order := "ASC"
		if sortOrder == "desc" {
			order = "DESC"
		}
		baseQuery = baseQuery.Order(sortBy + " " + order)
	}

	// Pagination
	if limit > 0 {
		baseQuery = baseQuery.Limit(limit).Offset(offset)
	}

	// Fetch data
	if err := baseQuery.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
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

func (r *userRepositoryImpl) DeleteUser(ctx context.Context, userID uuid.UUID) error {

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
