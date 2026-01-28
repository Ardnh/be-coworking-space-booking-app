package entities

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ReviewID   uuid.UUID    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	BookingID  uuid.UUID    `gorm:"type:uuid;not null"`
	UserID     uuid.UUID    `gorm:"type:uuid;not null;index"`
	ResourceID uuid.UUID    `gorm:"type:uuid;not null;index"`
	VendorID   uuid.UUID    `gorm:"type:uuid;not null;index"`
	Rating     int          `gorm:"type:int;not null;check:rating >= 1 AND rating <= 5"`
	Comment    *string      `gorm:"type:text"`  // Nullable
	Images     ReviewImages `gorm:"type:jsonb"` // Nullable JSONB

	// Relationships
	Booking  *Booking  `gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
	User     *Users    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Resource *Resource `gorm:"foreignKey:ResourceID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Vendor   *Vendor   `gorm:"foreignKey:VendorID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	// Timestamps
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	// NOTE: NO DeletedAt - table doesn't have soft delete
}

type ReviewImages []string

// TableName overrides the table name
func (Review) TableName() string {
	return "reviews"
}
