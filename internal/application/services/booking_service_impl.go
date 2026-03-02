package services

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	"github.com/sirupsen/logrus"
)

type BookingServiceImpl struct {
	repo repositories.BookingRespository
	log  *logrus.Logger
}

func NewBookingRepository(repo repositories.BookingRespository, log *logrus.Logger) services.BookingService {
	return &BookingServiceImpl{
		repo: repo,
		log:  log,
	}
}
