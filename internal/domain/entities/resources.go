package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Resource struct {
	ResourceID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	VendorID           uuid.UUID `gorm:"type:uuid;not null;index"`
	ResourceTypeID     uuid.UUID `gorm:"type:uuid;not null"`
	ResourceName       string    `gorm:"type:varchar(255);not null;index:idx_resources_resource_name"`
	Description        *string   `gorm:"type:text"`
	MaxSeatsPerSession int       `gorm:"type:int;not null;check:max_seats_per_session > 0"` // fix: hapus comment "per jam"
	SessionDuration    int       `gorm:"type:smallint;not null;check:session_duration > 0"` // fix: smallint, tambah check
	OperationTimeFrom  int       `gorm:"type:smallint;not null;check:operation_time_from >= 0 AND operation_time_from <= 23"`
	OperationTimeTo    int       `gorm:"type:smallint;not null;check:operation_time_to >= 1 AND operation_time_to <= 24"` // 24 = midnight
	StartDate          time.Time `gorm:"type:date;not null"`                                                              // fix: tambah StartDate
	EndDate            time.Time `gorm:"type:date;not null"`
	PricePerSession    float64   `gorm:"type:decimal(10,2);not null;check:price_per_session > 0"` // fix: PricePerUnit → PricePerSession
	Images             []*string `gorm:"type:jsonb;serializer:json"`
	Location           *string   `gorm:"type:varchar(255)"`
	Status             string    `gorm:"type:varchar(20);not null;default:'active';check:status IN ('active','inactive')"`
	// Relationships
	Vendor       *Vendor       `gorm:"foreignKey:VendorID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	ResourceType *ResourceType `gorm:"foreignKey:ResourceTypeID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	// Timestamps
	CreatedAt time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"index"`
}

// TableName overrides the table name
func (Resource) TableName() string {
	return "resources"
}

// Constants for resource status
const (
	ResourceStatusActive   = "active"
	ResourceStatusInactive = "inactive"
)

// ResourceImages is a custom type for JSONB images field
type ResourceImages []string

// Scan implements sql.Scanner interface for reading from database
func (ri *ResourceImages) Scan(value interface{}) error {
	if value == nil {
		*ri = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal JSONB value")
	}

	return json.Unmarshal(bytes, ri)
}

// Value implements driver.Valuer interface for writing to database
func (ri ResourceImages) Value() (driver.Value, error) {
	if ri == nil {
		return nil, nil
	}
	return json.Marshal(ri)
}

// BeforeSave hook for validation
func (r *Resource) BeforeSave(tx *gorm.DB) error {
	// Validate status
	if r.Status != ResourceStatusActive && r.Status != ResourceStatusInactive {
		return errors.New("invalid resource status: must be 'active' or 'inactive'")
	}

	// Validate capacity
	if r.Capacity < 0 {
		return errors.New("capacity cannot be negative")
	}

	// Validate price
	if r.PricePerUnit < 0 {
		return errors.New("price per unit cannot be negative")
	}

	return nil
}

// BeforeCreate hook
func (r *Resource) BeforeCreate(tx *gorm.DB) error {
	// Set default status if not provided
	if r.Status == "" {
		r.Status = ResourceStatusActive
	}
	return nil
}
