package services

import (
	"context"
	"mime/multipart"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
)

type ResourceService interface {
	CreateResource(ctx context.Context, req dto.CreateResourceRequestDto, images []*multipart.FileHeader) (*dto.ResourceResponseDto, error)
}
