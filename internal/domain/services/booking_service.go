package services

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
)

type BookingService interface {
	CreateBooking(c context.Context, booking *dto.CreateBookingDto) (*dto.BookingDto, error)
	CancelBooking(c context.Context, bookingId string) error
}
