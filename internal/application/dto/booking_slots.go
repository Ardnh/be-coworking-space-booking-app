package dto

type BookingSlotDto struct {
	Date     string `json:"date" validate:"required"`
	TimeFrom string `json:"time_from" validate:"required"`
	TimeTo   string `json:"time_to" validate:"required"`
}
