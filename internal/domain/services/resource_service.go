package services

import (
	"context"
	"mime/multipart"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/google/uuid"
)

type ResourceService interface {
	GetAllResources(ctx context.Context, filter *dto.ResourceFilterDto) ([]*dto.ResourceDto, int, error)
	GetResourceByVendorId(ctx context.Context, vendorId uuid.UUID) ([]*dto.ResourceDto, error)
	GetResourceById(ctx context.Context, resourceId uuid.UUID) (*dto.ResourceDto, error)
	CreateResource(ctx context.Context, req *dto.CreateResourceRequestDto, images []*multipart.FileHeader) (*dto.ResourceDto, error)
	UpdateResource(ctx context.Context, resourceId uuid.UUID, newResourceImage []*multipart.FileHeader, req *dto.UpdateResourceRequestDto) (*dto.ResourceDto, error)
	DeleteResource(ctx context.Context, resourceId uuid.UUID) error
}
