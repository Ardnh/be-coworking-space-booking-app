package entities

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	PaymentID     uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	BookingID     uuid.UUID `gorm:"type:uuid;not null;index:idx_payments_booking_id"`
	UserID        int       `gorm:"type:integer;not null;index:idx_payments_user_id"` // INTEGER, not UUID!
	Amount        float64   `gorm:"type:decimal(10,2);not null;check:amount > 0"`
	PaymentMethod string    `gorm:"type:varchar(50);not null;check:payment_method IN ('credit_card','debit_card','bank_transfer','ewallet','cash')"`
	PaymentDate   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_payments_date,sort:desc"`
	PaymentStatus string    `gorm:"type:varchar(20);not null;index:idx_payments_status;check:payment_status IN ('pending','completed','failed','refunded','cancelled')"`
	Notes         *string   `gorm:"type:text"` // Nullable

	// Relationships
	Booking *Booking `gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
	// Note: No FK constraint for user_id in migration, but we can still have relationship
	// User    *Users   `gorm:"foreignKey:UserID"` // Uncomment if you want to add relationship

	// Timestamps
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;index:idx_payments_created_at,sort:desc"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	// NOTE: NO DeletedAt - no soft delete in this table
}

func (Payment) TableName() string {
	return "payments"
}
