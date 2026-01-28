package entities

import (
	"time"

	"github.com/google/uuid"
)

type PricingRule struct {
	RuleID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ResourceID      uuid.UUID  `gorm:"type:uuid;not null;index:idx_pricing_rules_resource_id"`
	DayOfWeek       int16      `gorm:"type:smallint;not null;index:idx_pricing_rules_day_of_week;check:day_of_week BETWEEN 0 AND 6"`
	TimeStart       time.Time  `gorm:"type:time;not null"`
	TimeEnd         time.Time  `gorm:"type:time;not null"`
	PriceMultiplier *float64   `gorm:"type:decimal(5,2)"`  // Nullable - untuk percentage markup/discount
	FixedPrice      *float64   `gorm:"type:decimal(10,2)"` // Nullable - untuk fixed price override
	IsActive        bool       `gorm:"type:boolean;not null;default:true"`
	StartDate       *time.Time `gorm:"type:date"` // Nullable - untuk seasonal rules
	EndDate         *time.Time `gorm:"type:date"` // Nullable - untuk seasonal rules

	// Relationships
	Resource *Resource `gorm:"foreignKey:ResourceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	// Timestamps
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	// NOTE: NO UpdatedAt and DeletedAt in this table
}

func (PricingRule) TableName() string {
	return "pricing_rules"
}
