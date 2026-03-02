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
	"github.com/Ardnh/be-coworking-space-booking-app/internal/utils/helpers"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ResourceServiceImpl struct {
	repo repositories.ResourceRepository
	cld  *cloudinary.Cloudinary
	log  *logrus.Logger
}

func NewResourceService(repo repositories.ResourceRepository, cld *cloudinary.Cloudinary, log *logrus.Logger) services.ResourceService {
	return &ResourceServiceImpl{
		repo: repo,
		cld:  cld,
		log:  log,
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
	var uploadedURLs []*string
	folder := fmt.Sprintf("resources/%s", req.VendorID)

	for _, fh := range imageFiles {
		url, err := cldHelper.UploadFile(ctx, s.cld, fh, folder)
		if err != nil {
			// Rollback: hapus file yang sudah terupload
			cldHelper.RollbackUploads(ctx, s.cld, uploadedURLs)
			return nil, fmt.Errorf("gagal upload '%s': %w", fh.Filename, err)
		}
		uploadedURLs = append(uploadedURLs, &url)
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

func (s *ResourceServiceImpl) GetAllResources(ctx context.Context, filter *dto.ResourceFilterDto) ([]*dto.ResourceDto, int, error) {

	result, totalItems, err := s.repo.GetAllResources(ctx, filter.ResourceName, filter.ResourceTypeId, filter.OperationTimeStart, filter.OperationTimeEnd, filter.SortBy, filter.SortOrder, filter.PageSize, filter.Page)
	if err != nil {
		return nil, 0, fmt.Errorf("gagal mendapatkan resource: %w", err)
	}

	resourceDtos := mapper.ToResourceListDTO(result)
	return resourceDtos, totalItems, nil
}

func (s *ResourceServiceImpl) GetResourceByVendorId(ctx context.Context, vendorId uuid.UUID) ([]*dto.ResourceDto, error) {

	result, err := s.repo.GetResourceByVendorId(ctx, vendorId)
	if err != nil {
		return nil, err
	}

	resourceDtos := mapper.ToResourceListDTO(result)
	return resourceDtos, nil
}

func (s *ResourceServiceImpl) GetResourceById(ctx context.Context, resourceId uuid.UUID) (*dto.ResourceDto, error) {

	result, err := s.repo.GetResourceById(ctx, resourceId)
	if err != nil {
		return nil, err
	}

	resourceDto := mapper.ToResourceDTO(result)
	return resourceDto, nil
}

func (s *ResourceServiceImpl) UpdateResource(ctx context.Context, resourceId uuid.UUID, newResourceImage []*multipart.FileHeader, req *dto.UpdateResourceRequestDto) (*dto.ResourceDto, error) {

	existingResource, err := s.repo.GetResourceById(ctx, resourceId)
	if err != nil {
		return nil, err
	}

	if req.ResourceName != nil {
		existingResource.ResourceName = *req.ResourceName
	}

	if req.Description != nil {
		existingResource.Description = req.Description
	}

	if req.ResourceTypeID != nil {
		existingResource.ResourceTypeID = *req.ResourceTypeID
	}

	if req.Capacity != nil {
		existingResource.Capacity = *req.Capacity
	}

	if req.PricePerUnit != nil {
		existingResource.PricePerUnit = *req.PricePerUnit
	}

	if req.Location != nil {
		existingResource.Location = req.Location
	}

	if req.OperationTimeFrom != nil {
		existingResource.OperationTimeFrom = *req.OperationTimeFrom
	}

	if req.OperationTimeTo != nil {
		existingResource.OperationTimeTo = *req.OperationTimeTo
	}

	if req.EndDate != nil {
		parsedTime, err := time.Parse(time.RFC3339, *req.EndDate)
		if err != nil {
			return nil, err
		}

		existingResource.EndDate = parsedTime
	}

	retainedImages := req.Images         // Urls from frontend
	oldImages := existingResource.Images // Urls from database

	if oldImages != nil {
		oldImagesArr := oldImages
		for _, img := range oldImagesArr {
			if !helpers.Contains(retainedImages, *img) {
				err := cldHelper.DeleteFromCloudinary(ctx, s.cld, *img)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	// Upload newResourceImage ke Cloudinary
	var uploadedNewImageUrls []*string
	if len(newResourceImage) > 0 {
		folder := fmt.Sprintf("resources/%s", existingResource.VendorID)
		for _, img := range newResourceImage {
			url, err := cldHelper.UploadFile(ctx, s.cld, img, folder)
			if err != nil {
				// Rollback: hapus file yang sudah terupload
				cldHelper.RollbackUploads(ctx, s.cld, uploadedNewImageUrls)
				return nil, fmt.Errorf("gagal upload '%s': %w", img.Filename, err)
			}
			uploadedNewImageUrls = append(uploadedNewImageUrls, &url)
		}
	}

	// Ambil URL hasil upload, append ke retainedImages
	retainedImages = append(retainedImages, uploadedNewImageUrls...)

	// Update existingResource.Images dengan retainedImages terbaru
	existingResource.Images = retainedImages

	result, err := s.repo.UpdateResource(ctx, existingResource)
	if err != nil {
		return nil, err
	}

	resourceDto := mapper.ToResourceDTO(result)
	return resourceDto, nil
}

func (s *ResourceServiceImpl) DeleteResource(ctx context.Context, resourceId uuid.UUID) error {

	err := s.repo.DeleteResource(ctx, resourceId)
	if err != nil {
		return err
	}

	return nil
}
