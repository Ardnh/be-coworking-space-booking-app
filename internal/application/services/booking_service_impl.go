package services

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/utils/helpers"
	"github.com/Ardnh/be-coworking-space-booking-app/pkg/constants"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type BookingServiceImpl struct {
	repo         repositories.BookingRespository
	resourceRepo repositories.ResourceRepository
	log          *logrus.Logger
}

func NewBookingRepository(repo repositories.BookingRespository, resourceRepo repositories.ResourceRepository, log *logrus.Logger) services.BookingService {
	return &BookingServiceImpl{
		repo:         repo,
		resourceRepo: resourceRepo,
		log:          log,
	}
}

func (s *BookingServiceImpl) CreateBooking(ctx context.Context, booking *dto.CreateBookingDto) (*dto.BookingDto, error) {

	userIdUuid, err := uuid.Parse(booking.UserID)
	if err != nil {
		return nil, err
	}

	resourceIdUuid, err := uuid.Parse(booking.ResourceID)
	if err != nil {
		return nil, err
	}

	resource, err := s.resourceRepo.GetResourceById(c, booking.ResourceID)
	if err != nil {
		return nil, err
	}

	// Calculate total price based on resource price, booking time slots and capacity
	totalPrice := resource.PricePerUnit * float64(len(booking.BookingTime)*booking.Capacity)
	bookingEntities := entities.Booking{
		UserID:        userIdUuid,
		ResourceID:    resourceIdUuid,
		TotalPrice:    totalPrice,
		Seats:         int(booking.Seats),
		BookingStatus: constants.BookingStatusConfirmed,
		PaymentStatus: constants.PaymentStatusPending,
		BookingCode:   helpers.GenerateBookingCode(),
	}

	bookingSlots := make([]entities.BookingSlots, 0, len(booking.BookingTime))
	for _, slot := range booking.BookingTime {
		bookingSlots = append(bookingSlots, entities.BookingSlots{
			SlotDate: slot.Date,
			SlotHour: slot.SlotHour,
		})
	}

	return nil, nil
}

func (s *BookingServiceImpl) CancelBooking(ctx context.Context, bookingId string) error {

	return nil
}
