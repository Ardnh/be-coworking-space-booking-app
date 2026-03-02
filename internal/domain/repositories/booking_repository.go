package repositories

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
)

type BookingRespository interface {
	CreateBooking(ctx context.Context, booking *entities.Booking, bookignSlots []*entities.BookingSlots) (*entities.Booking, error)
}
