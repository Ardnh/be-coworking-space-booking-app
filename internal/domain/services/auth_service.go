package services

import (
	"context"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (*string, *string, error)
}
