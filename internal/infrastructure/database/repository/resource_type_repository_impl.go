package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	errConst "github.com/Ardnh/be-coworking-space-booking-app/pkg/errors"
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errConst.ErrNotFound
		}
		return nil, err
	}

	return &resourceType, err
}

type cachedResourceTypeResult struct {
	Data  []*entities.ResourceType `json:"data"`
	Total int                      `json:"total"`
}

func (r *ResourceRespositoryImpl) GetAllResourceType(ctx context.Context, name string, limit int, offset int, sortBy string, sortOrder string) ([]*entities.ResourceType, int, error) {

	// Build cache key based on query params
	cacheKey := fmt.Sprintf("resource_type:all:%s:%d:%d:%s:%s", name, limit, offset, sortBy, sortOrder)

	// Try get from Redis
	cached, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var result cachedResourceTypeResult
		if jsonErr := json.Unmarshal([]byte(cached), &result); jsonErr == nil {
			return result.Data, result.Total, nil
		}
	}

	var resourceTypes []*entities.ResourceType
	var total int64

	baseQuery := r.db.WithContext(ctx).Model(&entities.ResourceType{})
	countQuery := r.db.WithContext(ctx).Model(&entities.ResourceType{})

	if name != "" {
		baseQuery = baseQuery.Where("resource_type_name ILIKE ?", "%"+name+"%")
		countQuery = countQuery.Where("resource_type_name ILIKE ?", "%"+name+"%")
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

	if err := baseQuery.Find(&resourceTypes).Error; err != nil {
		return nil, 0, err
	}

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Store to Redis
	result := cachedResourceTypeResult{Data: resourceTypes, Total: int(total)}
	if jsonData, jsonErr := json.Marshal(result); jsonErr == nil {
		r.redis.Set(ctx, cacheKey, jsonData, 5*time.Minute)
	}

	return resourceTypes, int(total), nil
}

func (r *ResourceRespositoryImpl) CreateResourceType(ctx context.Context, resourceType *entities.ResourceType) (*entities.ResourceType, error) {

	if err := r.db.WithContext(ctx).Model(&entities.ResourceType{}).Create(&resourceType).Error; err != nil {
		return nil, err
	}

	pattern := "resource_type:all:*"
	var cursor uint64
	totalDeleted := 0
	for {
		keys, nextCursor, err := r.redis.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			break
		}
		if len(keys) > 0 {
			r.redis.Del(ctx, keys...)
			totalDeleted += len(keys)
		}
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}
	return resourceType, nil
}

func (r *ResourceRespositoryImpl) UpdateResourceType(ctx context.Context, resourceType *entities.ResourceType) (*entities.ResourceType, error) {

	result := r.db.WithContext(ctx).
		Model(&entities.ResourceType{}).
		Where("resource_type_id = ?", resourceType.ResourceTypeID).
		Updates(resourceType)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errConst.ErrNotFound
		}
		return nil, result.Error
	}

	var updatedResourceType entities.ResourceType
	err := r.db.WithContext(ctx).
		Model(&entities.ResourceType{}).
		Where("resource_type_id = ? ", resourceType.ResourceTypeID).
		First(&updatedResourceType).Error

	if err != nil {
		return nil, err
	}

	pattern := "resource_type:all:*"
	var cursor uint64
	for {
		keys, nextCursor, _ := r.redis.Scan(ctx, cursor, pattern, 100).Result()
		if len(keys) > 0 {
			r.redis.Del(ctx, keys...)
		}
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}

	return &updatedResourceType, nil
}

func (r *ResourceRespositoryImpl) DeleteResourceType(ctx context.Context, resourceTypeId uuid.UUID) error {

	var resourceType entities.ResourceType
	err := r.db.WithContext(ctx).
		Where("resource_type_id = ?", resourceTypeId).
		First(&resourceType).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errConst.ErrNotFound
		}
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

	pattern := "resource_type:all:*"
	var cursor uint64
	for {
		keys, nextCursor, _ := r.redis.Scan(ctx, cursor, pattern, 100).Result()
		if len(keys) > 0 {
			r.redis.Del(ctx, keys...)
		}
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}

	return nil
}
