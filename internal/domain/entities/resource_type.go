package entities

import (
	"time"

	"github.com/google/uuid"
)

type ResourceType struct {
	ResourceTypeID       uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ResourceTypeName     string     `gorm:"type:varchar(255);not null;index:idx_resources_resource_name"`
	Description          *string    `gorm:"type:text"`
	ParentResourceTypeID *uuid.UUID `gorm:"type:uuid;index:idx_resources_parent_resource_id"`
	Icon                 *string    `gorm:"type:varchar(255)"`

	// Self-referencing relationships
	ParentCategory *ResourceType  `gorm:"foreignKey:ParentResourceTypeID;references:ResourceTypeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	SubCategories  []ResourceType `gorm:"foreignKey:ParentResourceTypeID;references:ResourceTypeID"`

	// Reverse relationship with Resources
	Resources []Resource `gorm:"foreignKey:ResourceTypeID"`

	// Timestamps
	CreatedAt time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"index"`
}

// TableName overrides the table name
func (ResourceType) TableName() string {
	return "resource_type"
}
