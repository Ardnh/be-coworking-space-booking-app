package dto

// ============== Response DTOs ==============
type BookingDto struct {
}

// ============== Request DTOs ==============
type CreateBookingDto struct {
	UserID      string `json:"user_id" validate:"required"`
	ResourceID  string `json:"resource_id" validate:"required"`
	BookingDate string `json:"booking_date" validate:"requiried"`
	TimeFrom    string `json:"time_from" validate:"required"`
	TimeTo      string `json:"time_to" validate:"required"`
	Capacity    int    `json:"capacity" validate:"required"`
}

type UpdateBookingDto struct {
}
