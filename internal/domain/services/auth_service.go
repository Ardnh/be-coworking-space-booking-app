package services

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (*string, *string, error)
	Register(ctx context.Context, user dto.CreateUserDto) error
}
