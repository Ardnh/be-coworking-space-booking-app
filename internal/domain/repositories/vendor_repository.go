package repositories

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type VendorRepository interface {
	GetAllVendors(ctx context.Context, vendorName string, city string, limit int, offset int, sortBy string, sortOrder string) ([]*entities.Vendor, int64, error)
	GetVendorByID(ctx context.Context, vendorID uuid.UUID) (*entities.Vendor, error)
	GetVendorsResourcesByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*entities.Resource, error)
	// GetVendorReviews(ctx context.Context, vendorID uuid.UUID) ([]*entities.Review, error)
	CreateVendor(ctx context.Context, vendor *entities.Vendor) (*entities.Vendor, error)
	UpdateVendor(ctx context.Context, vendor *entities.Vendor) (*entities.Vendor, error)
	DeleteVendor(ctx context.Context, vendorID uuid.UUID) error
}
