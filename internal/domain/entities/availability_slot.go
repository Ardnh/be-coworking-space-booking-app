package entities

import (
	"time"

	"github.com/google/uuid"
)

type AvailabilitySlot struct {
	SlotID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ResourceID     uuid.UUID `gorm:"type:uuid;not null;index:idx_availability_slots_resource_id;index:idx_availability_slots_resource_date,priority:1;uniqueIndex:uq_resource_slot,priority:1"`
	Date           time.Time `gorm:"type:date;not null;index:idx_availability_slots_date;uniqueIndex:uq_resource_slot,priority:2"`
	StartTime      time.Time `gorm:"type:timestamp;not null;uniqueIndex:uq_resource_slot,priority:3"`
	EndTime        time.Time `gorm:"type:timestamp;not null;uniqueIndex:uq_resource_slot,priority:4"`
	MaxCapacity    int       `gorm:"type:int;not null;check:max_capacity > 0"`
	BookedCapacity int       `gorm:"type:int;not null;default:0;check:booked_capacity >= 0 AND booked_capacity <= max_capacity"`
	IsAvailable    bool      `gorm:"type:boolean;not null;default:true;index:idx_availability_slots_is_available"`
	PriceOverride  float64   `gorm:"type:decimal(10,2);default:0;check:price_override >= 0"`

	// Relationships
	Resource *Resource `gorm:"foreignKey:ResourceID;constraint:OnDelete:CASCADE"`

	// Timestamps
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}

func (AvailabilitySlot) TableName() string {
	return "availability_slots"
}
