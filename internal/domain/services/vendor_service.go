package services

import (
	"context"
	"mime/multipart"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/google/uuid"
)

type VendorService interface {
	GetAllVendors(ctx context.Context, params dto.VendorFilterRequest) ([]dto.VendorDto, int, error)
	GetVendorByID(ctx context.Context, vendorID uuid.UUID) (*dto.VendorDto, error)
	GetVendorsResourcesByVendorID(ctx context.Context, vendorID uuid.UUID) ([]dto.ResourceDto, error)
	// GetVendorReviews(ctx context.Context, vendorID uuid.UUID) ([]*entities.Review, error)
	CreateVendor(ctx context.Context, file *multipart.FileHeader, vendor *dto.CreateVendorRequestDto) (*dto.VendorDto, error)
	UpdateVendor(ctx context.Context, vendorID uuid.UUID, file *multipart.FileHeader, vendor *dto.UpdateVendorRequestDto) (*dto.VendorDto, error)
	DeleteVendor(ctx context.Context, vendorID uuid.UUID) error
}
