package repositories

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
)

type AuthRepository interface {
	Login(ctx context.Context, email string, password string) (*entities.Users, error)
}
