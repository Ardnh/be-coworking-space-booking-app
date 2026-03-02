package entities

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	BookingID     uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID        uuid.UUID       `gorm:"type:uuid;not null;index:idx_bookings_user_id;index:idx_bookings_user_status,priority:1"`
	ResourceID    uuid.UUID       `gorm:"type:uuid;not null;index:idx_bookings_resource_id;index:idx_bookings_resource_date,priority:1;index:idx_bookings_resource_date_time,priority:1"`
	BookingDate   time.Time       `gorm:"type:date;not null;index:idx_bookings_booking_date;index:idx_bookings_resource_date,priority:2;index:idx_bookings_resource_date_time,priority:2"`
	TimeFrom      time.Time       `gorm:"type:time;not null;index:idx_bookings_resource_date_time,priority:3"`
	TimeTo        time.Time       `gorm:"type:time;not null;index:idx_bookings_resource_date_time,priority:4"`
	Duration      BookingDuration `gorm:"type:interval;not null"`
	TotalPrice    float64         `gorm:"type:decimal(10,2);not null;check:total_price > 0"`
	BookingStatus string          `gorm:"type:varchar(20);not null;default:'pending';index:idx_bookings_booking_status;index:idx_bookings_user_status,priority:2;check:booking_status IN ('pending','confirmed','cancelled','completed','rejected')"`
	PaymentStatus string          `gorm:"type:varchar(20);not null;default:'pending';index:idx_bookings_payment_status;check:payment_status IN ('pending','paid','failed','refunded','cancelled')"`
	BookingCode   string          `gorm:"type:varchar(20);not null;uniqueIndex:idx_bookings_booking_code;check:booking_code ~ '^[A-Z0-9]{3,20}$'"`
	Notes         *string         `gorm:"type:text"` // Nullable

	// Relationships
	User     *Users    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Resource *Resource `gorm:"foreignKey:ResourceID;constraint:OnDelete:RESTRICT"`

	// Reverse relationship
	Reviews []Review `gorm:"foreignKey:BookingID"`

	// Timestamps
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;index:idx_bookings_created_at"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}

type BookingDuration time.Duration

func (Booking) TableName() string {
	return "bookings"
}
