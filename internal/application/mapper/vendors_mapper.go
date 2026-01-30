package mapper

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
)

func ToVendorListResponse(vendor []*entities.Vendor, page int, pageSize int, totalPages int, total int64) *dto.VendorListResponseDto {
	return &dto.VendorListResponseDto{
		Vendors:   ToVendorListDto(vendor),
		Total:     total,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPages,
	}
}

func ToVendorListDto(vendor []*entities.Vendor) []dto.VendorResponseDto {
	var vendorsListDto []dto.VendorResponseDto
	for _, v := range vendor {
		vendorsListDto = append(vendorsListDto, ToVendorDto(v))
	}
	return vendorsListDto
}

func ToVendorDto(vendor *entities.Vendor) dto.VendorResponseDto {
	return dto.VendorResponseDto{
		VendorID:    vendor.VendorID.String(),
		OwnerUserID: vendor.OwnerUserID.String(),
		VendorName:  vendor.VendorName,
		Address:     vendor.Address,
		City:        vendor.City,
		PhoneNumber: vendor.PhoneNumber,
		Email:       vendor.Email,
		Description: vendor.Description,
		Rating:      vendor.Rating,
		Status:      vendor.Status,
		CreatedAt:   vendor.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   vendor.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func ToDetailVendorResponse(vendor *entities.Vendor) dto.VendorResponseDto {
	return dto.VendorResponseDto{
		VendorID:    vendor.VendorID.String(),
		OwnerUserID: vendor.OwnerUserID.String(),
		VendorName:  vendor.VendorName,
		Address:     vendor.Address,
		City:        vendor.City,
		PhoneNumber: vendor.PhoneNumber,
		Email:       vendor.Email,
		Description: vendor.Description,
		Rating:      vendor.Rating,
		Status:      vendor.Status,
		CreatedAt:   vendor.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   vendor.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func ToVendorResourceListDto(resources []*entities.Resource) []dto.ResourceResponseDto {
	var resourcesListDto []dto.ResourceResponseDto
	for _, r := range resources {
		resourcesListDto = append(resourcesListDto, ToVendorResourceDto(r))
	}

	return resourcesListDto
}

func ToVendorResourceDto(resource *entities.Resource) dto.ResourceResponseDto {
	return dto.ResourceResponseDto{
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
