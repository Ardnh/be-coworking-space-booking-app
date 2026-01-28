package services

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/google/uuid"
)

type VendorService interface {
	GetAllVendors(ctx context.Context, params dto.VendorFilterRequest) ([]*dto.VendorListResponse, error)
	GetVendorByID(ctx context.Context, vendorID uuid.UUID) (*dto.VendorDetailResponse, error)
	GetVendorsResourcesByVendorID(ctx context.Context, vendorID uuid.UUID) (*dto.VendorPublicProfileResponse, error)
	// GetVendorReviews(ctx context.Context, vendorID uuid.UUID) ([]*entities.Review, error)
	CreateVendor(ctx context.Context, vendor *dto.CreateVendorRequest) error
	UpdateVendor(ctx context.Context, vendor *dto.UpdateVendorRequest) error
	DeleteVendor(ctx context.Context, vendorID uuid.UUID) error
}
