package dto

import (
	"time"

	"github.com/google/uuid"
)

// ============== Response DTOs ==============
type BlockedDate struct {
	BlockID    string    `json:"block_id"`
	ResourceID string    `json:"resource_id"`
	Date       string    `json:"date"`
	TimeFrom   string    `json:"time_from"`
	TimeTo     string    `json:"time_to"`
	Reason     string    `json:"reason"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ============== Request DTOs ==============
type CreateBlockedDateRequest struct {
	ResourceID uuid.UUID `json:"resource_id" validate:"required"`
	Date       string    `json:"start_date" validate:"required,datetime=2006-01-02"`
	TimeFrom   string    `json:"time_from" validate:"required,datetime=15:04"`
	TimeTo     string    `json:"time_to" validate:"required,datetime=15:04,gtfield=TimeFrom"`
	Reason     string    `json:"reason" validate:"required,min=1,max=500"`
	CreatedBy  string    `json:"created_by" validate:"required"`
}

type UpdateBlockedDateRequest struct {
	Date     *string `json:"start_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
	TimeFrom *string `json:"time_from,omitempty" validate:"omitempty,datetime=15:04"`
	TimeTo   *string `json:"time_to,omitempty" validate:"omitempty,datetime=15:04"`
	Reason   *string `json:"reason,omitempty" validate:"omitempty,min=1,max=500"`
}

// ============== Filter blocked date ==============
type BlockedDateFilter struct {
	ResourceID uuid.UUID `json:"resource_id,omitempty"`
	Date       string    `json:"date,omitempty" validate:"omitempty,datetime=2006-01-02"`
	TimeFrom   string    `json:"time_from,omitempty" validate:"omitempty,datetime=15:04"`
	TimeTo     string    `json:"time_to,omitempty" validate:"omitempty,datetime=15:04"`
}
