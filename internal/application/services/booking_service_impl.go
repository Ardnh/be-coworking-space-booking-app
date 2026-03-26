package services

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
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

func (s *BookingServiceImpl) CreateBooking(c context.Context, booking *dto.CreateBookingDto) (*dto.BookingDto, error) {

	return nil, nil
}

func (s *BookingServiceImpl) CancelBooking(c context.Context, bookingId string) error {

	return nil
}
