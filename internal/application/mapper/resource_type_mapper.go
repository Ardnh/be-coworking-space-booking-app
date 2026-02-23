package mapper

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
)

func ToResourceTypeListDto(resourceType []*entities.ResourceType) []*dto.ResourceTypeDto {

	var resourceTypeListDto []*dto.ResourceTypeDto
	for _, v := range resourceType {
		resourceTypeListDto = append(resourceTypeListDto, ToResourceTypeDto(v))
	}

	return resourceTypeListDto
}

func ToResourceTypeDto(resourceType *entities.ResourceType) *dto.ResourceTypeDto {

	return &dto.ResourceTypeDto{
		ResourceTypeId:       resourceType.ResourceTypeID.String(),
		ResourceTypeName:     resourceType.ResourceTypeName,
		ParentResourceTypeId: resourceType.ParentResourceTypeID,
		Description:          resourceType.Description,
		Icon:                 resourceType.Icon,
		CreatedAt:            resourceType.CreatedAt.Format("2006-01-02T15:04:05"),
		UpdatedAt:            resourceType.UpdatedAt.Format("2006-01-02T15:04:05"),
	}
}
