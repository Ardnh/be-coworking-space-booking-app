package mapper

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
)

func ToBlockedDateListDTO(blockedDates []*entities.BlockedDate) []*dto.BlockedDate {
	var result []*dto.BlockedDate
	for _, blockedDate := range blockedDates {
		result = append(result, ToBlockedDateDTO(blockedDate))
	}
	return result
}

func ToBlockedDateDTO(blockedDate *entities.BlockedDate) *dto.BlockedDate {
	return &dto.BlockedDate{
		BlockID:    blockedDate.BlockID.String(),
		ResourceID: blockedDate.ResourceID.String(),
		Date:       blockedDate.Date.Format("2006-01-02T15:04:05"),
		TimeFrom:   blockedDate.TimeFrom,
		TimeTo:     blockedDate.TimeTo,
		Reason:     blockedDate.Reason,
		CreatedBy:  blockedDate.CreatedBy,
		CreatedAt:  blockedDate.CreatedAt,
		UpdatedAt:  blockedDate.UpdatedAt,
	}
}
