package repository

import (
	"context"
	"fmt"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookingRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewBookingRepostory(db *gorm.DB, redis *redis.Client) repositories.BookingRespository {
	return &BookingRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *BookingRepositoryImpl) CreateBooking(ctx context.Context, booking *entities.Booking, bookingSlots []*entities.BookingSlots) (*entities.Booking, error) {
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		// 1. Ambil resource untuk dapat kapasitas
		var resource entities.Resource
		if err := tx.First(&resource, booking.ResourceID).Error; err != nil {
			return fmt.Errorf("resource not found: %w", err)
		}

		// 2. Kumpulkan slot hours yang diminta
		slotHours := make([]int, len(bookingSlots))
		for i, slot := range bookingSlots {
			slotHours[i] = slot.SlotHour
		}

		// 3. Lock & hitung seat terpakai per slot
		type SlotUsage struct {
			SlotHour    int
			BookedSeats int
		}

		var usages []SlotUsage
		err := tx.Model(&entities.BookingSlots{}).
			Select("slot_hour, COALESCE(SUM(bookings.seats), 0) as booked_seats").
			Joins("JOIN bookings ON bookings.booking_id = booking_slots.booking_id AND bookings.status = 'confirmed'").
			Where("booking_slots.resource_id = ? AND booking_slots.slot_date = ? AND booking_slots.slot_hour IN ?", booking.ResourceID, bookingSlots[0].SlotDate, slotHours).
			Group("slot_hour").
			Clauses(clause.Locking{Strength: "UPDATE", Options: "NOWAIT"}).
			Find(&usages).Error

		if err != nil {
			return fmt.Errorf("failed to check availability: %w", err)
		}

		// 4. Buat map untuk lookup cepat
		usageMap := make(map[int]int)
		for _, u := range usages {
			usageMap[u.SlotHour] = u.BookedSeats
		}

		// 5. Validasi semua slot harus tersedia
		for _, hour := range slotHours {
			booked := usageMap[hour]
			available := resource.Capacity - booked
			if available < booking.Seats {
				return fmt.Errorf("slot %d:00-%d:00 hanya tersisa %d seat, butuh %d", hour, hour+1, available, booking.Seats)
			}
		}

		// 6. Insert booking + slots (cascade)
		booking.BookingStatus = "confirmed"
		if err := tx.Create(booking).Error; err != nil {
			return fmt.Errorf("failed to create booking: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return booking, nil
}
