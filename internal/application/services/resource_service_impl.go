package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/mapper"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	cldHelper "github.com/Ardnh/be-coworking-space-booking-app/internal/utils/cloudinary"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/google/uuid"
)

type ResourceServiceImpl struct {
	repo repositories.ResourceRepository
	cld  *cloudinary.Cloudinary
}

func NewResourceService(repo repositories.ResourceRepository, cld *cloudinary.Cloudinary) services.ResourceService {
	return &ResourceServiceImpl{
		repo: repo,
		cld:  cld,
	}
}

func (s *ResourceServiceImpl) CreateResource(ctx context.Context, req *dto.CreateResourceRequestDto, imageFiles []*multipart.FileHeader) (*dto.ResourceDto, error) {

	// === 1. Validasi semua file dulu sebelum upload ===
	allowedExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	const maxSize = 5 * 1024 * 1024

	for _, fh := range imageFiles {
		ext := strings.ToLower(filepath.Ext(fh.Filename))
		if !allowedExt[ext] {
			return nil, fmt.Errorf("file '%s' format tidak didukung", fh.Filename)
		}
		if fh.Size > maxSize {
			return nil, fmt.Errorf("file '%s' melebihi batas ukuran 5MB", fh.Filename)
		}
	}

	// === 2. Upload semua file ke Cloudinary ===
	var uploadedURLs []string
	folder := fmt.Sprintf("resources/%s", req.VendorID)

	for _, fh := range imageFiles {
		url, err := cldHelper.UploadFile(ctx, s.cld, fh, folder)
		if err != nil {
			// Rollback: hapus file yang sudah terupload
			cldHelper.RollbackUploads(ctx, s.cld, uploadedURLs)
			return nil, fmt.Errorf("gagal upload '%s': %w", fh.Filename, err)
		}
		uploadedURLs = append(uploadedURLs, url)
	}

	vendorIdUUID, err := uuid.Parse(req.VendorID)
	if err != nil {
		return nil, fmt.Errorf("gagal parse vendor ID '%s': %w", req.VendorID, err)
	}

	resourceTypeIdUUID, err := uuid.Parse(req.ResourceTypeID)
	if err != nil {
		return nil, fmt.Errorf("gagal parse resource type ID '%s': %w", req.ResourceTypeID, err)
	}

	// Format date only
	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("format end_date harus YYYY-MM-DD")
	}

	resourceEntities := entities.Resource{
		VendorID:          vendorIdUUID,
		ResourceName:      req.ResourceName,
		ResourceTypeID:    resourceTypeIdUUID,
		Description:       req.Description,
		Capacity:          req.Capacity,
		OperationTimeFrom: req.OperationTimeFrom,
		OperationTimeTo:   req.OperationTimeTo,
		EndDate:           endDate,
		PricePerUnit:      req.PricePerUnit,
		Images:            uploadedURLs,
		Location:          req.Location,
		Status:            req.Status,
	}

	result, err := s.repo.CreateResource(ctx, &resourceEntities)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat resource: %w", err)
	}

	resourceDto := mapper.ToResourceDTO(result)
	return resourceDto, nil
}
