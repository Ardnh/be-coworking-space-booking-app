package repository

import (
	"context"
	"fmt"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ResourceRespositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewResourceRespository(db *gorm.DB, redis *redis.Client) repositories.ResourceRepository {
	return &ResourceRespositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *ResourceRespositoryImpl) GetAllResources(ctx context.Context, resourceName string, resourceTypeId string, operationTimeStart string, operationTimeEnd string, sortBy string, sortOrder string, pageSize int, page int) ([]*entities.Resource, int, error) {
	var resources []*entities.Resource

	baseQuery := r.db.WithContext(ctx).Model(&entities.Resource{})

	// Filter by resource name (partial match)
	if resourceName != "" {
		baseQuery = baseQuery.Where("resource_name ILIKE ?", "%"+resourceName+"%")
	}

	// Filter by resource type ID
	if resourceTypeId != "" {
		baseQuery = baseQuery.Where("resource_type_id = ?", resourceTypeId)
	}

	// Filter by operation time range
	if operationTimeStart != "" {
		baseQuery = baseQuery.Where("operation_time >= ?", operationTimeStart)
	}
	if operationTimeEnd != "" {
		baseQuery = baseQuery.Where("operation_time <= ?", operationTimeEnd)
	}

	// Sorting
	if sortBy != "" {
		if sortOrder == "" {
			sortOrder = "asc"
		}
		orderClause := sortBy + " " + sortOrder
		baseQuery = baseQuery.Order(orderClause)
	}

	// Pagination
	if pageSize > 0 {
		offset := (page - 1) * pageSize
		baseQuery = baseQuery.Limit(pageSize).Offset(offset)
	}

	// Execute query
	if err := baseQuery.Find(&resources).Error; err != nil {
		return nil, 0, err
	}

	return resources, 0, nil
}

func (r *ResourceRespositoryImpl) GetResourceByVendorId(ctx context.Context, vendorId uuid.UUID) ([]*entities.Resource, error) {

	var resources []*entities.Resource
	if err := r.db.WithContext(ctx).Model(&entities.Resource{}).Where("vendor_id = ?", vendorId).Find(&resources).Error; err != nil {
		return nil, err
	}

	return resources, nil
}

func (r *ResourceRespositoryImpl) GetResourceById(ctx context.Context, resourceId uuid.UUID) (*entities.Resource, error) {

	var resources *entities.Resource
	if err := r.db.WithContext(ctx).Model(&entities.Resource{}).Where("resource_id = ?", resourceId).First(&resources).Error; err != nil {
		return nil, err
	}

	return resources, nil
}

func (r *ResourceRespositoryImpl) CreateResource(ctx context.Context, resource *entities.Resource) (*entities.Resource, error) {

	if err := r.db.WithContext(ctx).Model(&entities.Resource{}).Create(&resource).Error; err != nil {
		return nil, fmt.Errorf("gagal membuat resource: %w", err)
	}

	return resource, nil
}

func (r *ResourceRespositoryImpl) UpdateResource(ctx context.Context, resource *entities.Resource) (*entities.Resource, error) {

	result := r.db.WithContext(ctx).Model(&entities.Resource{}).Where("resource_id = ?", resource.ResourceID).Updates(resource)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var updatedResource entities.Resource
	if err := r.db.WithContext(ctx).Model(&entities.Resource{}).Where("resource_id = ?", resource.ResourceID).First(&updatedResource).Error; err != nil {
		return nil, err
	}

	return &updatedResource, nil
}

func (r *ResourceRespositoryImpl) DeleteResource(ctx context.Context, resourceId uuid.UUID) error {

	result := r.db.WithContext(ctx).Where("resource_id = ?", resourceId).Delete(&entities.Resource{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
