package entities

import (
	"time"

	"github.com/google/uuid"
)

type Vendor struct {
	VendorID    uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	OwnerUserID uuid.UUID `gorm:"type:uuid;not null"`
	VendorName  string    `gorm:"type:varchar(255);not null;index:idx_vendors_name"`
	Address     string    `gorm:"type:varchar(255);not null"`
	City        string    `gorm:"type:varchar(255);not null"`
	PhoneNumber string    `gorm:"column:phone_number;type:varchar(20);not null"`
	Email       string    `gorm:"type:varchar(255);uniqueIndex:idx_vendors_email;not null"`
	Description string    `gorm:"type:text;not null"`
	Rating      float64   `gorm:"type:decimal(10,2);not null;default:0"`
	Status      string    `gorm:"type:varchar(20);not null;default:'active';check:status IN ('active', 'inactive')"`

	User *Users `gorm:"foreignKey:OwnerUserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"index:idx_vendors_deleted_at"`
}

func (Vendor) TableName() string {
	return "vendors"
}

// Constants sesuai migration
const (
	VendorStatusActive   = "active"
	VendorStatusInactive = "inactive"
)
