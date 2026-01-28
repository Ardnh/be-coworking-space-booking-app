package entities

import (
	"time"

	"github.com/google/uuid"
)

type BlockedDate struct {
	BlockID    uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ResourceID uuid.UUID `gorm:"type:uuid;not null;index:idx_blocked_dates_resource_id;index:idx_blocked_dates_date_range,priority:1"`
	StartDate  time.Time `gorm:"type:date;not null;index:idx_blocked_dates_start_date;index:idx_blocked_dates_date_range,priority:2"`
	EndDate    time.Time `gorm:"type:date;not null;index:idx_blocked_dates_end_date;index:idx_blocked_dates_date_range,priority:3"`
	Reason     string    `gorm:"type:text;not null"`
	CreatedBy  uuid.UUID `gorm:"type:uuid;not null"`

	// Relationships
	Resource *Resource `gorm:"foreignKey:ResourceID;constraint:OnDelete:CASCADE"`
	Creator  *Users    `gorm:"foreignKey:CreatedBy;constraint:OnDelete:CASCADE"`

	// Timestamps
	CreatedAt time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"index:idx_blocked_dates_deleted_at"` // Soft delete
}

func (BlockedDate) TableName() string {
	return "blocked_dates"
}
