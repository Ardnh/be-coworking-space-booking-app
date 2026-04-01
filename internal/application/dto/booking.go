package dto

// ============== Response DTOs ==============
type BookingDto struct {
	BookingID     string  `json:"booking_id"`
	UserID        string  `json:"user_id"`
	ResourceID    string  `json:"resource_id"`
	Seats         int     `json:"seats"`
	TotalPrice    float64 `json:"total_price"`
	BookingStatus string  `json:"booking_status"`
	PaymentStatus string  `json:"payment_status"`
	BookingCode   string  `json:"booking_code"`
	CreatedAt     string  `json:"created_at"`
}

type SlotAvailabilityDto struct {
	AvailableTime  int  `json:"available_time"`
	AvailableSeats int  `json:"available_seats"`
	IsAvailable    bool `json:"is_available"`
}

// ============== Request DTOs ==============
type CreateBookingRequestDto struct {
	UserID      string           `json:"user_id" validate:"required"`
	ResourceID  string           `json:"resource_id" validate:"required"`
	TotalPrice  float64          `json:"total_price" validate:"required"`
	BookingCode string           `json:"booking_code,omitempty"`
	Seats       int              `json:"seats" validate:"required,gte=1"`
	BookingTime []BookingSlotDto `json:"booking_time" validate:"required"`
}

type GetSlotAvailabilityRequestDto struct {
	ResourceID   string `params:"resourceId" validate:"required"`
	SelectedDate string `query:"date"         validate:"required,datetime=2006-01-02"`
	Seats        int    `query:"seats"        validate:"required,min=1"`
}
