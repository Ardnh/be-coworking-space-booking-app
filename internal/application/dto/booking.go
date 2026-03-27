package dto

// ============== Response DTOs ==============
type BookingDto struct {
}

// ============== Request DTOs ==============
type CreateBookingDto struct {
	UserID      string           `json:"user_id" validate:"required"`
	ResourceID  string           `json:"resource_id" validate:"required"`
	TotalPrice  float64          `json:"total_price" validate:"required"`
	BookingCode string           `json:"booking_code,omitempty"`
	Seats       int              `json:"seats" validate:"required,gte=1"`
	BookingTime []BookingSlotDto `json:"booking_time" validate:"required"`
}
