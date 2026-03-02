package entities

import (
	"time"

	"github.com/google/uuid"
)

type BookingSlots struct {
	ID         uuid.UUID `gorm:"type:uuid;not null;primaryKey;autoIncrement:false"`
	BookingID  uuid.UUID `gorm:"type:uuid;not null;index:idx_booking_slots_booking_id"`
	ResourceID uuid.UUID `gorm:"type:uuid;not null;index:idx_booking_slots_resource_id"`
	SlotDate   time.Time `gorm:"type:date;not null;index:idx_booking_slots_slot_date"`
	SlotHour   int       `gorm:"type:int;not null;index:idx_booking_slots_slot_hour"`
	Seats      int       `gorm:"type:int;not null;check:seats > 0"`
}
