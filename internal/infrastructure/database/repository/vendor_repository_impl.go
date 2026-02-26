package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

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

type cachedVendorResult struct {
	Data  []*entities.Vendor `json:"data"`
	Total int                `json:"total"`
}

func NewVendorRepository(db *gorm.DB, redis *redis.Client) repositories.VendorRepository {
	return &VendorRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *VendorRepositoryImpl) GetAllVendors(ctx context.Context, vendorName string, city string, limit int, offset int, sortBy string, sortOrder string) ([]*entities.Vendor, int, error) {

	cacheKey := fmt.Sprintf("resource_type:all:%s:%d:%d:%s:%s", vendorName, limit, offset, sortBy, sortOrder)

	// Try get from Redis
	cached, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var result cachedVendorResult
		if jsonErr := json.Unmarshal([]byte(cached), &result); jsonErr == nil {
			return result.Data, result.Total, nil
		}
	}

	var vendor []*entities.Vendor
	var total int64 = 0

	baseQuery := r.db.WithContext(ctx).Model(&entities.Vendor{})

	if vendorName != "" {
		baseQuery = baseQuery.Where("name LIKE ?", "%"+vendorName+"%")
	}

	if city != "" {
		baseQuery = baseQuery.Where("city = ?", city)
	}

	if limit > 0 {
		baseQuery = baseQuery.Limit(limit)
	}

	if offset > 0 {
		baseQuery = baseQuery.Offset(offset)
	}

	if sortBy != "" {
		baseQuery = baseQuery.Order(sortBy + " " + sortOrder)
	}

	if err := baseQuery.Find(&vendor).Error; err != nil {
		return nil, 0, err
	}

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return vendor, int(total), nil
}

func (r *VendorRepositoryImpl) GetVendorByID(ctx context.Context, vendorID uuid.UUID) (*entities.Vendor, error) {

	var vendor entities.Vendor
	if err := r.db.WithContext(ctx).Model(&entities.Vendor{}).Where("vendor_id = ?", vendorID).First(&vendor).Error; err != nil {
		return nil, err
	}

	return &vendor, nil
}

func (r *VendorRepositoryImpl) GetVendorsResourcesByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*entities.Resource, error) {

	var isVendorExists bool
	if err := r.db.WithContext(ctx).Model(&entities.Vendor{}).Where("id = ?", vendorID).First(&isVendorExists).Error; err != nil {
		return nil, err
	}

	if !isVendorExists {
		return nil, errors.New("vendor not found")
	}

	var vendorResources []*entities.Resource
	if err := r.db.WithContext(ctx).Model(&entities.Resource{}).Where("vendor_id = ?", vendorID).Find(&vendorResources).Error; err != nil {
		return nil, err
	}

	return vendorResources, nil
}

func (r *VendorRepositoryImpl) CreateVendor(ctx context.Context, vendor *entities.Vendor) (*entities.Vendor, error) {

	if err := r.db.WithContext(ctx).Model(&entities.Vendor{}).Create(&vendor).Error; err != nil {
		return nil, err
	}

	return vendor, nil
}

func (r *VendorRepositoryImpl) UpdateVendor(ctx context.Context, vendor *entities.Vendor) (*entities.Vendor, error) {

	result := r.db.WithContext(ctx).
		Model(&entities.Vendor{}).
		Where("vendor_id = ?", vendor.VendorID).
		Updates(vendor)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var updatedVendor entities.Vendor
	err := r.db.WithContext(ctx).
		Where("vendor_id = ?", vendor.VendorID).
		First(&updatedVendor).Error

	if err != nil {
		return nil, err
	}

	return &updatedVendor, nil
}

func (r *VendorRepositoryImpl) DeleteVendor(ctx context.Context, vendorID uuid.UUID) error {

	result := r.db.WithContext(ctx).
		Where("vendor_id = ?", vendorID).
		Delete(&entities.Vendor{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
