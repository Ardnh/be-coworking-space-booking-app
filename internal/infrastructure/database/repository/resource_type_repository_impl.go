package repository

import (
	"context"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ResourceTypeRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewResourceType(db *gorm.DB, redis *redis.Client) repositories.ResourceTypeRepository {
	return &ResourceRespositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *ResourceRespositoryImpl) GetResourceTypeById(ctx context.Context, resourceTypeId uuid.UUID) (*entities.ResourceType, error) {

	var resourceType entities.ResourceType
	err := r.db.WithContext(ctx).
		Where("resource_type_id = ?", resourceTypeId).
		First(&resourceType).Error

	if err != nil {
		return nil, err
	}

	return &resourceType, err
}

func (r *ResourceRespositoryImpl) GetAllResourceType(ctx context.Context, name string, limit int, offset int, sortBy string, sortOrder string) ([]*entities.ResourceType, int, error) {

	var resourceType []*entities.ResourceType
	var total int64

	baseQuery := r.db.WithContext(ctx).Model(&entities.ResourceType{})

	if name != "" {
		baseQuery = baseQuery.Where("resource_type_name = ?", name)
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

	if err := baseQuery.Find(&resourceType).Error; err != nil {
		return nil, 0, err
	}

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return resourceType, int(total), nil
}

func (r *ResourceRespositoryImpl) CreateResourceType(ctx context.Context, resourceType *entities.ResourceType) (*entities.ResourceType, error) {

	if err := r.db.WithContext(ctx).Model(&entities.ResourceType{}).Create(&resourceType).Error; err != nil {
		return nil, err
	}

	return resourceType, nil
}

func (r *ResourceRespositoryImpl) UpdateResourceType(ctx context.Context, resourceType *entities.ResourceType) (*entities.ResourceType, error) {

	result := r.db.WithContext(ctx).
		Model(&entities.ResourceType{}).
		Where("resource_type_id = ? ", resourceType.ResourceTypeID).
		Updates(resourceType)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var updatedResourceType entities.ResourceType
	err := r.db.WithContext(ctx).
		Model(&entities.ResourceType{}).
		First(&updatedResourceType).Error

	if err != nil {
		return nil, err
	}

	return &updatedResourceType, nil
}

func (r *ResourceRespositoryImpl) DeleteResourceType(ctx context.Context, resourceTypeId uuid.UUID) error {

	var resourceType entities.ResourceType
	err := r.db.WithContext(ctx).
		Where("resource_type_id = ?", resourceTypeId).
		First(&resourceType).Error

	if err != nil {
		return err
	}

	result := r.db.WithContext(ctx).
		Where("resource_type_id = ?", resourceType.ResourceTypeID).
		Delete(&entities.ResourceType{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
