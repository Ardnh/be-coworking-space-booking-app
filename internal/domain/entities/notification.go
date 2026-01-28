package entities

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	NotificationID   uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	BookingID        uuid.UUID  `gorm:"type:uuid;not null;index:idx_notifications_booking_id"`
	UserID           uuid.UUID  `gorm:"type:uuid;not null;index:idx_notifications_user_id;index:idx_notifications_user_unread,priority:1,where:is_read = false"`
	NotificationType string     `gorm:"type:varchar(255);not null;index:idx_notifications_type"`
	Title            string     `gorm:"type:varchar(255);not null"`
	Message          string     `gorm:"type:text;not null"`
	IsRead           bool       `gorm:"type:boolean;not null;default:false;index:idx_notifications_is_read,where:is_read = false;index:idx_notifications_user_unread,priority:2,where:is_read = false"`
	IsSent           bool       `gorm:"type:boolean;not null;default:false;index:idx_notifications_is_sent,where:is_sent = false;index:idx_notifications_unsent,priority:1,where:is_sent = false"`
	SentAt           *time.Time `gorm:"type:timestamp"` // Nullable

	// Relationships
	Booking *Booking `gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE"`
	User    *Users   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`

	// Timestamps
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;index:idx_notifications_created_at,sort:desc;index:idx_notifications_user_unread,priority:3,sort:desc,where:is_read = false;index:idx_notifications_unsent,priority:2,where:is_sent = false"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	// NOTE: NO DeletedAt - no soft delete
}

func (Notification) TableName() string {
	return "notifications"
}
