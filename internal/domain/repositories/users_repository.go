package repositories

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type UsersRepository interface {
	GetAllUsers(ctx context.Context, usernameQuery string, emailQuery string, fullNameQuery string, limit int, offset int, sortBy string, sortOrder string) ([]*entities.Users, int, error)
	CreateUser(ctx context.Context, user *entities.Users) error
	UpdateUser(ctx context.Context, user *entities.Users) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}
