package mapper

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
)

func ToVendorListDto(vendor []*entities.Vendor) []dto.VendorDto {
	var vendorsListDto []dto.VendorDto
	for _, v := range vendor {
		vendorsListDto = append(vendorsListDto, ToVendorDto(v))
	}
	return vendorsListDto
}

func ToVendorDto(vendor *entities.Vendor) dto.VendorDto {
	return dto.VendorDto{
		VendorID:     vendor.VendorID.String(),
		OwnerUserID:  vendor.OwnerUserID.String(),
		VendorName:   vendor.VendorName,
		Address:      vendor.Address,
		City:         vendor.City,
		PhoneNumber:  vendor.PhoneNumber,
		Email:        vendor.Email,
		Description:  vendor.Description,
		Rating:       vendor.Rating,
		Status:       vendor.Status,
		ProfileImage: vendor.ProfileImage,
		CreatedAt:    vendor.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:    vendor.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func ToDetailVendorResponse(vendor *entities.Vendor) dto.VendorDto {
	return dto.VendorDto{
		VendorID:     vendor.VendorID.String(),
		OwnerUserID:  vendor.OwnerUserID.String(),
		VendorName:   vendor.VendorName,
		Address:      vendor.Address,
		City:         vendor.City,
		PhoneNumber:  vendor.PhoneNumber,
		Email:        vendor.Email,
		Description:  vendor.Description,
		Rating:       vendor.Rating,
		Status:       vendor.Status,
		ProfileImage: vendor.ProfileImage,
		CreatedAt:    vendor.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:    vendor.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func ToVendorResourceListDto(resources []*entities.Resource) []dto.ResourceDto {
	var resourcesListDto []dto.ResourceDto
	for _, r := range resources {
		resourcesListDto = append(resourcesListDto, ToVendorResourceDto(r))
	}

	return resourcesListDto
}

func ToVendorResourceDto(resource *entities.Resource) dto.ResourceDto {
	return dto.ResourceDto{
		ResourceID:   resource.ResourceID.String(),
		VendorID:     resource.VendorID.String(),
		ResourceName: resource.ResourceName,
		Description:  resource.Description,
		Capacity:     resource.Capacity,
		PricePerUnit: resource.PricePerUnit,
		Status:       resource.Status,
		CreatedAt:    resource.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:    resource.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}
