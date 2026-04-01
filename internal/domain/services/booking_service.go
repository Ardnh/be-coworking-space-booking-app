package services

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
)

type BookingService interface {
	CreateBooking(ctx context.Context, booking *dto.CreateBookingRequestDto) (*dto.BookingDto, error)
	GetSlotAvailability(ctx context.Context, req dto.GetSlotAvailabilityRequestDto) ([]*dto.SlotAvailabilityDto, error)
	CancelBooking(ctx context.Context, bookingId string) error
}
