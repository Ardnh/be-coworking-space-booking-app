package entities

import (
	"time"

	"github.com/google/uuid"
)

type BookingSlots struct {
	BookingSlotID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	BookingID     uuid.UUID `gorm:"type:uuid;not null"`
	SlotDate      time.Time `gorm:"type:date;not null"`
	SlotHour      int       `gorm:"not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`

	// Relasi ke Booking (opsional, untuk preload)
	Booking *Booking `gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
}

func (BookingSlots) TableName() string {
	return "booking_slots"
}
