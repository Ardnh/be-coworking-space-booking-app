package services

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
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

func (s *VendorServiceImpl) GetAllVendors(ctx context.Context, params dto.VendorFilterRequest) ([]*dto.VendorListResponse, error) {
	return nil, nil
}

func (s *VendorServiceImpl) GetVendorByID(ctx context.Context, vendorID uuid.UUID) (*dto.VendorDetailResponse, error) {
	return nil, nil
}

func (s *VendorServiceImpl) GetVendorsResourcesByVendorID(ctx context.Context, vendorID uuid.UUID) (*dto.VendorPublicProfileResponse, error) {
	return nil, nil
}

func (s *VendorServiceImpl) CreateVendor(ctx context.Context, vendor *dto.CreateVendorRequest) error {
	return nil
}

func (s *VendorServiceImpl) UpdateVendor(ctx context.Context, vendor *dto.UpdateVendorRequest) error {
	return nil
}

func (s *VendorServiceImpl) DeleteVendor(ctx context.Context, vendorID uuid.UUID) error {
	return nil
}
