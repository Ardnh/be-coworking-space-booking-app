package entities

import (
	"time"

	"github.com/google/uuid"
)

type BookingHistory struct {
	HistoryID    uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	BookingID    uuid.UUID       `gorm:"type:uuid;not null;index:idx_booking_history_booking_id,priority:1;index:idx_booking_history_audit,priority:1"`
	UserID       uuid.UUID       `gorm:"type:uuid;not null;index:idx_booking_history_user_id,priority:1"`
	OldStatus    string          `gorm:"type:varchar(255);not null;index:idx_booking_history_status_transition,priority:1"`
	NewStatus    string          `gorm:"type:varchar(255);not null;index:idx_booking_history_status_transition,priority:2;index:idx_booking_history_audit,priority:3"`
	ChangedBy    uuid.UUID       `gorm:"type:uuid;not null;index:idx_booking_history_changed_by,priority:1"`
	ChangeReason string          `gorm:"type:text;not null"`
	ChangedAt    time.Time       `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;index:idx_booking_history_booking_id,priority:2,sort:desc;index:idx_booking_history_user_id,priority:2,sort:desc;index:idx_booking_history_changed_by,priority:2,sort:desc;index:idx_booking_history_changed_at,sort:desc;index:idx_booking_history_audit,priority:2,sort:desc"`
	Metadata     HistoryMetadata `gorm:"type:jsonb;index:idx_booking_history_metadata,type:gin"` // GIN index for JSONB

	// Relationships
	Booking       *Booking `gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
	User          *Users   `gorm:"foreignKey:UserID;constraint:OnDelete:RESTRICT"`
	ChangedByUser *Users   `gorm:"foreignKey:ChangedBy;constraint:OnDelete:RESTRICT"`

	// NOTE: NO CreatedAt, UpdatedAt, DeletedAt - this is immutable audit log
}

type HistoryMetadata map[string]interface{}

func (BookingHistory) TableName() string {
	return "booking_history"
}
