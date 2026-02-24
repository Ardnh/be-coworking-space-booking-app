package mapper

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
)

func ToResourceDTO(resource *entities.Resource) *dto.ResourceDto {
	return &dto.ResourceDto{
		ResourceID:         resource.ResourceID.String(),
		VendorID:           resource.VendorID.String(),
		ResourceTypeID:     resource.ResourceTypeID,
		ResourceName:       resource.ResourceName,
		Description:        resource.Description,
		Capacity:           resource.Capacity,
		OperationTimeStart: resource.OperationTimeFrom.String(),
		OperationTimeEnd:   resource.OperationTimeTo.String(),
		EndDate:            resource.EndDate.String(),
		PricePerUnit:       resource.PricePerUnit,
		Images:             resource.Images,
		Location:           resource.Location,
		Status:             resource.Status,
		CreatedAt:          resource.CreatedAt.Format("2006-01-02T15:04:05"),
		UpdatedAt:          resource.UpdatedAt.Format("2006-01-02T15:04:05"),
	}
}

func ToResourceTypeDTO(resourceType *entities.ResourceType) *dto.ResourceTypeDto {
	return &dto.ResourceTypeDto{
		ResourceTypeId:       resourceType.ResourceTypeID.String(),
		ResourceTypeName:     resourceType.ResourceTypeName,
		Description:          resourceType.Description,
		ParentResourceTypeId: resourceType.ParentResourceTypeID,
		Icon:                 resourceType.Icon,
		CreatedAt:            resourceType.CreatedAt.Format("2006-01-02T15:04:05"),
		UpdatedAt:            resourceType.UpdatedAt.Format("2006-01-02T15:04:05"),
	}
}
