package dto

import "github.com/google/uuid"

// ============== Response DTOs ==============
type ResourceDto struct {
	ResourceID         string    `json:"resource_id"`
	VendorID           string    `json:"vendor_id"`
	ResourceName       string    `json:"resource_name"`
	ResourceTypeID     string    `json:"resource_type_id"`
	Description        *string   `json:"description,omitempty"`
	Capacity           int       `json:"capacity"`
	OperationTimeStart string    `json:"operation_time_start"`
	OperationTimeEnd   string    `json:"operation_time_end"`
	EndDate            string    `json:"end_date"`
	PricePerUnit       float64   `json:"price_per_unit"`
	Images             []*string `json:"images,omitempty"`
	Location           *string   `json:"location,omitempty"`
	Status             string    `json:"status"`
	CreatedAt          string    `json:"created_at"`
	UpdatedAt          string    `json:"updated_at"`
}

type ResourceResponseDto struct {
	ResourceID         string          `json:"resource_id"`
	VendorID           string          `json:"vendor_id"`
	ResourceName       string          `json:"resource_name"`
	ResourceTypeID     string          `json:"resource_type_id"`
	ResourceType       ResourceTypeDto `json:"resource_type"`
	Description        *string         `json:"description,omitempty"`
	Capacity           int             `json:"capacity"`
	OperationTimeStart string          `json:"operation_time_start"`
	OperationTimeEnd   string          `json:"operation_time_end"`
	EndDate            string          `json:"end_date"`
	PricePerUnit       float64         `json:"price_per_unit"`
	Images             []string        `json:"images,omitempty"`
	Location           *string         `json:"location,omitempty"`
	Status             string          `json:"status"`
	CreatedAt          string          `json:"created_at"`
	UpdatedAt          string          `json:"updated_at"`
}

// ============== Request DTOs ==============
type CreateResourceRequestDto struct {
	VendorID          string                      `json:"vendor_id" validate:"required,uuid"`
	ResourceTypeID    string                      `json:"resource_type_id" validate:"required,uuid"`
	ResourceName      string                      `json:"resource_name" validate:"required,min=3"`
	Description       *string                     `json:"description,omitempty"`
	Capacity          int                         `json:"capacity" validate:"required,min=1"`
	OperationTimeFrom string                      `json:"operation_time_from" validate:"required"`
	OperationTimeTo   string                      `json:"operation_time_to" validate:"required"`
	EndDate           string                      `json:"end_date" validate:"required"`
	PricePerUnit      float64                     `json:"price_per_unit" validate:"required,gt=0"`
	Location          *string                     `json:"location,omitempty"`
	Status            string                      `json:"status,omitempty" validate:"omitempty,oneof=active inactive"`
	BlockedDates      []*CreateBlockedDateRequest `json:"blocked_date,omitempty"`
}

type UpdateResourceRequestDto struct {
	ResourceName      *string    `json:"resource_name,omitempty"`
	ResourceTypeID    *uuid.UUID `json:"resource_type_id,omitempty"`
	Description       *string    `json:"description,omitempty"`
	OperationTimeFrom *string    `json:"operation_time_from,omitempty"`
	OperationTimeTo   *string    `json:"operation_time_to,omitempty"`
	EndDate           *string    `json:"end_date,omitempty"`
	Capacity          *int       `json:"capacity,omitempty" validate:"omitempty,min=1"`
	PricePerUnit      *float64   `json:"price_per_unit,omitempty" validate:"omitempty,gt=0"`
	Images            []*string  `json:"images,omitempty"`
	Location          *string    `json:"location,omitempty"`
}

// ============== Filter ==============
type ResourceFilterDto struct {
	ResourceName       string `json:"resource_name"`
	ResourceTypeId     string `json:"resource_type_id"`
	OperationTimeStart string `json:"operation_time_start"`
	OperationTimeEnd   string `json:"operation_time_end"`
	SortBy             string `json:"sort_by"`
	SortOrder          string `json:"sort_order"`
	PageSize           int    `json:"page_size"`
	Page               int    `json:"page_number"`
	Limit              int    `json:"limit"`
}
