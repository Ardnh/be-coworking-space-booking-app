package dto

type ResourceResponseDto struct {
	ResourceID   string   `json:"resource_id"`
	VendorID     string   `json:"vendor_id"`
	ResourceName string   `json:"resource_name"`
	ResourceType string   `json:"resource_type"`
	CategoryID   string   `json:"category_id"`
	Description  *string  `json:"description,omitempty"`
	Capacity     int      `json:"capacity"`
	PricePerUnit float64  `json:"price_per_unit"`
	Images       []string `json:"images,omitempty"`
	Location     *string  `json:"location,omitempty"`
	Status       string   `json:"status"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
}

type CreateResourceRequestDto struct {
	VendorID     string    `json:"vendor_id" validate:"required"`
	ResourceName string    `json:"resource_name" validate:"required,min=3"`
	ResourceType string    `json:"resource_type" validate:"required"`
	CategoryID   string    `json:"category_id" validate:"required"`
	Description  *string   `json:"description,omitempty"`
	Capacity     int       `json:"capacity" validate:"required,min=1"`
	PricePerUnit float64   `json:"price_per_unit" validate:"required,gt=0"`
	Images       []*string `json:"images,omitempty"`
	Location     *string   `json:"location,omitempty"`
	Status       string    `json:"status,omitempty" validate:"omitempty,oneof=active inactive"`
}

type UpdateResourceRequestDto struct {
	ResourceName *string   `json:"resource_name,omitempty"`
	ResourceType *string   `json:"resource_type,omitempty"`
	CategoryID   *string   `json:"category_id,omitempty"`
	Description  *string   `json:"description,omitempty"`
	Capacity     *int      `json:"capacity,omitempty" validate:"omitempty,min=1"`
	PricePerUnit *float64  `json:"price_per_unit,omitempty" validate:"omitempty,gt=0"`
	Images       []*string `json:"images,omitempty"`
	Location     *string   `json:"location,omitempty"`
	Status       *string   `json:"status,omitempty" validate:"omitempty,oneof=active inactive"`
}
