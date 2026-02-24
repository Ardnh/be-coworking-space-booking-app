package services

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/mapper"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	"github.com/google/uuid"
)

type VendorServiceImpl struct {
	repo repositories.VendorRepository
}

func NewVendorService(repo repositories.VendorRepository) services.VendorService {
	return &VendorServiceImpl{
		repo: repo,
	}
}

func (s *VendorServiceImpl) GetAllVendors(ctx context.Context, params dto.VendorFilterRequest) ([]dto.VendorDto, int, error) {

	vendors, total, err := s.repo.GetAllVendors(ctx, params.Name, params.City, params.PageSize, params.Offset, params.SortBy, params.SortDirection)
	if err != nil {
		return nil, 0, err
	}

	result := mapper.ToVendorListDto(vendors)
	return result, total, nil
}

func (s *VendorServiceImpl) GetVendorByID(ctx context.Context, vendorID uuid.UUID) (*dto.VendorDto, error) {

	vendor, err := s.repo.GetVendorByID(ctx, vendorID)
	if err != nil {
		return nil, err
	}

	result := mapper.ToVendorDto(vendor)
	return &result, nil
}

func (s *VendorServiceImpl) GetVendorsResourcesByVendorID(ctx context.Context, vendorID uuid.UUID) ([]dto.ResourceDto, error) {

	result, err := s.repo.GetVendorsResourcesByVendorID(ctx, vendorID)
	if err != nil {
		return nil, err
	}

	resultDto := mapper.ToVendorResourceListDto(result)
	return resultDto, nil
}

func (s *VendorServiceImpl) CreateVendor(ctx context.Context, vendor *dto.CreateVendorRequestDto) (*dto.VendorDto, error) {

	ownerUserIdUuid, err := uuid.Parse(vendor.OwnerUserID)
	if err != nil {
		return nil, err
	}

	vendorEntities := entities.Vendor{
		OwnerUserID: ownerUserIdUuid,
		VendorName:  vendor.VendorName,
		Address:     vendor.Address,
		City:        vendor.City,
		PhoneNumber: vendor.PhoneNumber,
		Email:       vendor.Email,
		Description: vendor.Description,
	}

	result, err := s.repo.CreateVendor(ctx, &vendorEntities)
	if err != nil {
		return nil, err
	}

	resultDto := mapper.ToVendorDto(result)
	return &resultDto, nil
}

func (s *VendorServiceImpl) UpdateVendor(ctx context.Context, vendorID uuid.UUID, vendor *dto.UpdateVendorRequestDto) (*dto.VendorDto, error) {

	existingVendor, err := s.repo.GetVendorByID(ctx, vendorID)
	if err != nil {
		return nil, err
	}

	if vendor.VendorName != nil {
		existingVendor.VendorName = *vendor.VendorName
	}
	if vendor.Address != nil {
		existingVendor.Address = *vendor.Address
	}
	if vendor.City != nil {
		existingVendor.City = *vendor.City
	}
	if vendor.PhoneNumber != nil {
		existingVendor.PhoneNumber = *vendor.PhoneNumber
	}
	if vendor.Email != nil {
		existingVendor.Email = *vendor.Email
	}
	if vendor.Description != nil {
		existingVendor.Description = *vendor.Description
	}

	result, err := s.repo.UpdateVendor(ctx, existingVendor)
	if err != nil {
		return nil, err
	}

	resultDto := mapper.ToVendorDto(result)
	return &resultDto, nil
}

func (s *VendorServiceImpl) DeleteVendor(ctx context.Context, vendorID uuid.UUID) error {
	return s.repo.DeleteVendor(ctx, vendorID)
}
