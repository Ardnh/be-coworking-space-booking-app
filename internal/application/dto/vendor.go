package dto

type VendorResponseDto struct {
	VendorID    string  `json:"vendor_id"`
	OwnerUserID string  `json:"owner_user_id"`
	VendorName  string  `json:"vendor_name"`
	Address     string  `json:"address"`
	City        string  `json:"city"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	RatingStars string  `json:"rating_stars"`
	Status      string  `json:"status"`
	StatusBadge string  `json:"status_badge"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type VendorListResponseDto struct {
	Vendors   []VendorResponseDto `json:"vendors"`
	Total     int                 `json:"total"`
	Page      int                 `json:"page"`
	PageSize  int                 `json:"page_size"`
	TotalPage int                 `json:"total_page"`
}

// ==================== REQUEST DTOs ====================
// CreateVendorRequest - DTO untuk membuat vendor baru
type CreateVendorRequestDto struct {
	OwnerUserID string `json:"owner_user_id" validate:"required,uuid"`
	VendorName  string `json:"vendor_name" validate:"required,min=3,max=255"`
	Address     string `json:"address" validate:"required,min=5,max=255"`
	City        string `json:"city" validate:"required,min=2,max=255"`
	PhoneNumber string `json:"phone_number" validate:"required,min=10,max=20"`
	Email       string `json:"email" validate:"required,email,max=255"`
	Description string `json:"description" validate:"required,min=10,max=2000"`
}

// UpdateVendorRequest - DTO untuk update vendor (semua field optional)
type UpdateVendorRequestDto struct {
	VendorName  *string `json:"vendor_name,omitempty" validate:"omitempty,min=3,max=255"`
	Address     *string `json:"address,omitempty" validate:"omitempty,min=5,max=255"`
	City        *string `json:"city,omitempty" validate:"omitempty,min=2,max=255"`
	PhoneNumber *string `json:"phone_number,omitempty" validate:"omitempty,min=10,max=20"`
	Email       *string `json:"email,omitempty" validate:"omitempty,email,max=255"`
	Description *string `json:"description,omitempty" validate:"omitempty,min=10,max=2000"`
}

// UpdateVendorStatusRequest - DTO untuk update status vendor
type UpdateVendorStatusRequestDto struct {
	Status string `json:"status" validate:"required,oneof=active inactive"`
}

// VendorFilterRequest - DTO untuk filter vendors
type VendorFilterRequest struct {
	Name          string   `json:"name,omitempty"`
	City          string   `json:"city,omitempty"`
	MinRating     *float64 `json:"min_rating,omitempty" validate:"omitempty,gte=0,lte=5"`
	SearchQuery   string   `json:"search_query,omitempty"`
	Page          int      `json:"page,omitempty" validate:"omitempty,min=1"`
	PageSize      int      `json:"page_size,omitempty" validate:"omitempty,min=1,max=100"`
	Offset        int      `json:"offset,omitempty" validate:"omitempty,min=0"`
	SortBy        string   `json:"sort_by,omitempty" validate:"omitempty,oneof=name rating created_at"`
	SortDirection string   `json:"sort_direction,omitempty" validate:"omitempty,oneof=asc desc"`
}
