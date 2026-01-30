package services

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/google/uuid"
)

type VendorService interface {
	GetAllVendors(ctx context.Context, params dto.VendorFilterRequest) (*dto.VendorListResponseDto, error)
	GetVendorByID(ctx context.Context, vendorID uuid.UUID) (*dto.VendorResponseDto, error)
	GetVendorsResourcesByVendorID(ctx context.Context, vendorID uuid.UUID) ([]dto.ResourceResponseDto, error)
	// GetVendorReviews(ctx context.Context, vendorID uuid.UUID) ([]*entities.Review, error)
	CreateVendor(ctx context.Context, vendor *dto.CreateVendorRequestDto) (*dto.VendorResponseDto, error)
	UpdateVendor(ctx context.Context, vendorID uuid.UUID, vendor *dto.UpdateVendorRequestDto) (*dto.VendorResponseDto, error)
	DeleteVendor(ctx context.Context, vendorID uuid.UUID) error
}
