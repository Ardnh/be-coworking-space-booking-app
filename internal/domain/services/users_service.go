package services

import (
	"context"

	"github.com/google/uuid"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
)

type UserService interface {
	GetAllUsers(ctx context.Context, params *dto.GetUserParams) ([]*dto.UserDto, int, error)
	GetUserById(ctx context.Context, userId uuid.UUID) (*dto.UserDto, error)
	CreateUser(ctx context.Context, user *dto.CreateUserDto) error
	UpdateUser(ctx context.Context, userID uuid.UUID, user *dto.UpdateUserDto) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}
