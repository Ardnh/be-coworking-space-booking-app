package dto

type BookingSlotDto struct {
	Date     string `json:"date" validate:"required"`
	SlotHour int    `json:"slot_hour" validate:"required"`
}
