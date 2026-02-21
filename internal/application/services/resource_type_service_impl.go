package services

import (
	"context"
	"time"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/mapper"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	"github.com/google/uuid"
)

type ResourceTypeServiceImpl struct {
	repo repositories.ResourceTypeRepository
}

func NewResourceTypeService(repo repositories.ResourceTypeRepository) services.ResourceTypeService {
	return &ResourceTypeServiceImpl{
		repo: repo,
	}
}

func (s *ResourceTypeServiceImpl) GetResourceTypeById(ctx context.Context, resourceTypeID uuid.UUID) (*dto.ResourceTypeDto, error) {

	result, err := s.repo.GetResourceTypeById(ctx, resourceTypeID)
	if err != nil {
		return nil, err
	}

	resourceTypeDto := mapper.ToResourceTypeDto(result)
	return resourceTypeDto, nil
}

func (s *ResourceTypeServiceImpl) GetAllResourceType(ctx context.Context, name string, limit int, offset int, sortBy string, sortOrder string) ([]*dto.ResourceTypeDto, int, error) {

	result, total, err := s.repo.GetAllResourceType(ctx, name, limit, offset, sortBy, sortOrder)
	if err != nil {
		return nil, 0, err
	}

	resourceTypeListDto := mapper.ToResourceTypeListDto(result)
	return resourceTypeListDto, total, nil
}

func (s *ResourceTypeServiceImpl) CreateResourceType(ctx context.Context, resourceType *dto.CreateResourceTypeRequestDto) (*dto.ResourceTypeDto, error) {

	resourceTypeEntities := entities.ResourceType{
		ResourceTypeName: resourceType.ResourceTypeName,
		Description:      resourceType.Description,
		Icon:             resourceType.Icon,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	result, err := s.repo.CreateResourceType(ctx, &resourceTypeEntities)
	if err != nil {
		return nil, err
	}

	resourceTypeDto := mapper.ToResourceTypeDto(result)
	return resourceTypeDto, nil
}

func (s *ResourceTypeServiceImpl) UpdateResourceType(ctx context.Context, resourceTypeID uuid.UUID, resourceType *dto.UpdateResourceTypeRequestDto) (*dto.ResourceTypeDto, error) {

	existingResourceType, err := s.repo.GetResourceTypeById(ctx, resourceTypeID)
	if err != nil {
		return nil, err
	}

	if resourceType.ResourceTypeName != nil {
		existingResourceType.ResourceTypeName = *resourceType.ResourceTypeName
	}

	if resourceType.Description != nil {
		existingResourceType.Description = resourceType.Description
	}

	if resourceType.Icon != nil {
		existingResourceType.Icon = resourceType.Icon
	}

	if resourceType.ParentResourceTypeId != nil {
		existingResourceType.ParentResourceTypeID = resourceType.ParentResourceTypeId
	}

	result, err := s.repo.UpdateResourceType(ctx, existingResourceType)
	if err != nil {
		return nil, err
	}

	resourceTypeDto := mapper.ToResourceTypeDto(result)
	return resourceTypeDto, nil
}

func (s *ResourceTypeServiceImpl) DeleteResourceType(ctx context.Context, resourceTypeID uuid.UUID) error {

	err := s.repo.DeleteResourceType(ctx, resourceTypeID)
	if err != nil {
		return err
	}

	return nil
}
