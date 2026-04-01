package services

import (
	"context"
	"fmt"
	"time"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/mapper"
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

func (s *BookingServiceImpl) CreateBooking(ctx context.Context, booking *dto.CreateBookingRequestDto) (*dto.BookingDto, error) {

	userIdUuid, err := uuid.Parse(booking.UserID)
	if err != nil {
		return nil, err
	}

	resourceIdUuid, err := uuid.Parse(booking.ResourceID)
	if err != nil {
		return nil, err
	}

	resource, err := s.resourceRepo.GetResourceById(ctx, resourceIdUuid)
	if err != nil {
		return nil, err
	}

	// Calculate total price based on resource price, booking time slots and capacity
	totalPrice := resource.PricePerUnit * float64(len(booking.BookingTime)*booking.Seats)
	bookingEntities := entities.Booking{
		UserID:        userIdUuid,
		ResourceID:    resourceIdUuid,
		TotalPrice:    totalPrice,
		Seats:         int(booking.Seats),
		BookingStatus: constants.BookingStatusConfirmed,
		PaymentStatus: constants.PaymentStatusPending,
		BookingCode:   helpers.GenerateBookingCode(),
	}

	bookingSlots := make([]*entities.BookingSlots, 0, len(booking.BookingTime))
	for _, slot := range booking.BookingTime {
		parsedDate, err := time.Parse("2006-01-02", slot.Date)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse date")
		}
		bookingSlots = append(bookingSlots, &entities.BookingSlots{
			SlotDate: parsedDate,
			SlotHour: slot.SlotHour,
		})
	}

	bookingResult, err := s.repo.CreateBooking(ctx, &bookingEntities, bookingSlots)
	if err != nil {
		return nil, err
	}

	bookingDto := mapper.ToBookingDto(bookingResult)
	return bookingDto, nil
}

func (s *BookingServiceImpl) GetSlotAvailability(ctx context.Context, req dto.GetSlotAvailabilityRequestDto) ([]dto.SlotAvailabilityDto, error) {

	return nil, nil
}

func (s *BookingServiceImpl) CancelBooking(ctx context.Context, bookingId string) error {

	return nil
}
