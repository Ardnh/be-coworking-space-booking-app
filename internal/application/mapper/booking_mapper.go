package mapper

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
)

func ToBookingDto(booking *entities.Booking) *dto.BookingDto {
	return &dto.BookingDto{
		BookingID:     booking.BookingID.String(),
		UserID:        booking.UserID.String(),
		ResourceID:    booking.ResourceID.String(),
		Seats:         booking.Seats,
		TotalPrice:    booking.TotalPrice,
		BookingStatus: booking.BookingStatus,
		PaymentStatus: booking.PaymentStatus,
		BookingCode:   booking.BookingCode,
		CreatedAt:     booking.CreatedAt.Format("2006-01-02T15:04:05"),
	}
}

func ToSlotAvailabilityDto(resourceOpenHour int, resourceCloseHour int, resourceMaxSlotPerHour int, slot []*entities.AvailabilitySlot) []dto.SlotAvailabilityDto {

	//

}
