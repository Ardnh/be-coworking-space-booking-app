package repository

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type VendorRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewVendorRepository(db *gorm.DB, redis *redis.Client) repositories.VendorRepository {
	return &VendorRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *VendorRepositoryImpl) GetAllVendors(ctx context.Context, vendorName string, city string, limit int, offset int, sortBy string, sortOrder string) ([]*entities.Vendor, error) {

	return nil, nil
}

func (r *VendorRepositoryImpl) GetVendorByID(ctx context.Context, vendorID uuid.UUID) (*entities.Vendor, error) {

	return nil, nil
}

func (r *VendorRepositoryImpl) GetVendorsResourcesByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*entities.Resource, error) {

	return nil, nil
}

func (r *VendorRepositoryImpl) GetVendorReviews(ctx context.Context, vendorID uuid.UUID) ([]*entities.Review, error) {

	return nil, nil
}

func (r *VendorRepositoryImpl) CreateVendor(ctx context.Context, vendor *entities.Vendor) error {

	return nil
}

func (r *VendorRepositoryImpl) UpdateVendor(ctx context.Context, vendor *entities.Vendor) error {

	return nil
}

func (r *VendorRepositoryImpl) DeleteVendor(ctx context.Context, vendorID uuid.UUID) error {

	return nil
}
