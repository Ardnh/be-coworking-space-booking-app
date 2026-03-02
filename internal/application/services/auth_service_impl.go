package services

import (
	"context"
	"errors"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/config"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	utils "github.com/Ardnh/be-coworking-space-booking-app/internal/utils/jwt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	repo repositories.AuthRepository
	log  *logrus.Logger
}

func NewAuthService(repo repositories.AuthRepository, log *logrus.Logger) services.AuthService {
	return &AuthServiceImpl{
		repo: repo,
		log:  log,
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, email string, password string) (*string, *string, error) {
	s.log.WithFields(logrus.Fields{
		"email": email,
	}).Info("attempting login")

	user, err := s.repo.Login(ctx, email, password)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"email": email,
			"error": err,
		}).Error("failed to fetch user from repository")
		// System error — jangan bocorkan detail DB ke user
		return nil, nil, errors.New("terjadi kesalahan, coba lagi nanti")
	}

	if user == nil {
		s.log.WithFields(logrus.Fields{
			"email": email,
		}).Warn("user not found")
		// Pesan samar — jangan bocorkan apakah email terdaftar atau tidak
		return nil, nil, errors.New("email atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"email": email,
		}).Warn("invalid credentials provided")
		// Pesan samar — sama dengan user not found agar tidak bisa di-enumerate
		return nil, nil, errors.New("email atau password salah")
	}

	config := config.LoadConfig()
	if config.App.JWTSecret == "" {
		s.log.Error("JWT secret is not configured")
		// System error — jangan bocorkan detail konfigurasi ke user
		return nil, nil, errors.New("terjadi kesalahan, coba lagi nanti")
	}

	secretKey := []byte(config.App.JWTSecret)
	token, expiredTimeISO, err := utils.GenerateToken(secretKey, user.UserID.String(), user.UserType)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"user_id":   user.UserID,
			"user_type": user.UserType,
			"error":     err,
		}).Error("failed to generate token")
		// System error — jangan bocorkan detail token generation ke user
		return nil, nil, errors.New("terjadi kesalahan, coba lagi nanti")
	}

	s.log.WithFields(logrus.Fields{
		"user_id":   user.UserID,
		"user_type": user.UserType,
	}).Info("login successful")

	return token, expiredTimeISO, nil
}
