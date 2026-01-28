package entities

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	CategoryID       uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CategoryName     string     `gorm:"type:varchar(255);not null;index:idx_categories_category_name"`
	Description      *string    `gorm:"type:text"`                                         // Nullable
	ParentCategoryID *uuid.UUID `gorm:"type:uuid;index:idx_categories_parent_category_id"` // Nullable, self-referencing
	Icon             *string    `gorm:"type:varchar(255)"`                                 // Nullable

	// Self-referencing relationships
	ParentCategory *Category  `gorm:"foreignKey:ParentCategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	SubCategories  []Category `gorm:"foreignKey:ParentCategoryID"`

	// Reverse relationship with Resources
	Resources []Resource `gorm:"foreignKey:CategoryID"`

	// Timestamps
	CreatedAt time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"index"`
}

// TableName overrides the table name
func (Category) TableName() string {
	return "categories"
}
