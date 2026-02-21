package dto

import "github.com/google/uuid"

type ResourceTypeDto struct {
	ResourceTypeId       string     `json:"resource_type_id" validate:"required,uuid"`
	ResourceTypeName     string     `json:"resource_type_name" validate:"required,min=3,max=255"`
	Description          *string    `json:"description" validate:"omitempty"`
	ParentResourceTypeId *uuid.UUID `json:"parent_resource_type_id" validate:"omitempty"`
	Icon                 *string    `json:"icon" validate:"omitempty"`
}

// ==================== REQUEST DTOs ====================
type CreateResourceTypeRequestDto struct {
	ResourceTypeName     string     `json:"resource_type_name" validate:"required,min=3,max=255"`
	Description          *string    `json:"description" validate:"omitempty"`
	ParentResourceTypeId *uuid.UUID `json:"parent_resource_type_id" validate:"omitempty"`
	Icon                 *string    `json:"icon" validate:"omitempty"`
}

type UpdateResourceTypeRequestDto struct {
	ResourceTypeId       string     `json:"resource_type_id" validate:"required,uuid"`
	ResourceTypeName     *string    `json:"resource_type_name" validate:"required,min=3,max=255"`
	Description          *string    `json:"description" validate:"omitempty"`
	ParentResourceTypeId *uuid.UUID `json:"parent_resource_type_id" validate:"omitempty"`
	Icon                 *string    `json:"icon" validate:"omitempty"`
}
