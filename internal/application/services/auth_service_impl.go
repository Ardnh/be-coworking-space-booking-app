package services

import (
	"context"
	"errors"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/config"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	utils "github.com/Ardnh/be-coworking-space-booking-app/internal/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) services.AuthService {
	return &AuthServiceImpl{
		repo: repo,
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, email string, password string) (*string, *string, error) {

	user, err := s.repo.Login(ctx, email, password)
	if err != nil {
		return nil, nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, nil, errors.New("Invalid credentials")
	}

	config := config.LoadConfig()
	if config.App.JWTSecret == "" {
		return nil, nil, errors.New("Failed to load jwt secret")
	}

	// Secret key untuk signing
	secretKey := []byte(config.App.JWTSecret)

	token, expiredTimeISO, err := utils.GenerateToken(secretKey, user.UserID.String(), user.UserType)
	if err != nil {
		return nil, nil, err
	}

	return token, expiredTimeISO, nil
}

func (s *AuthServiceImpl) Register(ctx context.Context, user dto.CreateUserDto) error {

	return nil
}
