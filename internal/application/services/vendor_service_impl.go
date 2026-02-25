package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/mapper"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/entities"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/repositories"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	cldHelper "github.com/Ardnh/be-coworking-space-booking-app/internal/utils/cloudinary"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/google/uuid"
)

type VendorServiceImpl struct {
	repo repositories.VendorRepository
	cld  *cloudinary.Cloudinary
}

func NewVendorService(repo repositories.VendorRepository, cld *cloudinary.Cloudinary) services.VendorService {
	return &VendorServiceImpl{
		repo: repo,
		cld:  cld,
	}
}

func (s *VendorServiceImpl) GetAllVendors(ctx context.Context, params dto.VendorFilterRequest) ([]dto.VendorDto, int, error) {

	vendors, total, err := s.repo.GetAllVendors(ctx, params.Name, params.City, params.PageSize, params.Offset, params.SortBy, params.SortDirection)
	if err != nil {
		return nil, 0, err
	}

	result := mapper.ToVendorListDto(vendors)
	return result, total, nil
}

func (s *VendorServiceImpl) GetVendorByID(ctx context.Context, vendorID uuid.UUID) (*dto.VendorDto, error) {

	vendor, err := s.repo.GetVendorByID(ctx, vendorID)
	if err != nil {
		return nil, err
	}

	result := mapper.ToVendorDto(vendor)
	return &result, nil
}

func (s *VendorServiceImpl) GetVendorsResourcesByVendorID(ctx context.Context, vendorID uuid.UUID) ([]dto.ResourceDto, error) {

	result, err := s.repo.GetVendorsResourcesByVendorID(ctx, vendorID)
	if err != nil {
		return nil, err
	}

	resultDto := mapper.ToVendorResourceListDto(result)
	return resultDto, nil
}

func (s *VendorServiceImpl) CreateVendor(ctx context.Context, profileImage *multipart.FileHeader, vendor *dto.CreateVendorRequestDto) (*dto.VendorDto, error) {

	// === 1. Parse UUID dulu (murah, sebelum upload) ===
	ownerUserIdUuid, err := uuid.Parse(vendor.OwnerUserID)
	if err != nil {
		return nil, fmt.Errorf("invalid owner_user_id: %w", err)
	}

	// === 2. Upload image (opsional) ===
	var imageURL *string
	if profileImage != nil {
		url, err := s.uploadProfileImage(ctx, profileImage, vendor.OwnerUserID)
		if err != nil {
			return nil, err
		}
		imageURL = &url
	}

	// === 3. Simpan ke DB ===
	vendorEntity := entities.Vendor{
		OwnerUserID:  ownerUserIdUuid,
		VendorName:   vendor.VendorName,
		Address:      vendor.Address,
		City:         vendor.City,
		PhoneNumber:  vendor.PhoneNumber,
		Email:        vendor.Email,
		Description:  vendor.Description,
		ProfileImage: *imageURL,
	}

	result, err := s.repo.CreateVendor(ctx, &vendorEntity)
	if err != nil {
		// Rollback: hapus image dari Cloudinary jika DB gagal
		if imageURL != nil {
			_ = cldHelper.DeleteFromCloudinary(ctx, s.cld, *imageURL)
		}
		return nil, err
	}

	resultDto := mapper.ToVendorDto(result)
	return &resultDto, nil
}

// uploadProfileImage — validasi + upload, file ditutup otomatis
func (s *VendorServiceImpl) uploadProfileImage(ctx context.Context, fh *multipart.FileHeader, id string) (string, error) {
	allowedExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	const maxSize = 5 * 1024 * 1024

	ext := strings.ToLower(filepath.Ext(fh.Filename))
	if !allowedExt[ext] {
		return "", fmt.Errorf("file '%s' format tidak didukung", fh.Filename)
	}
	if fh.Size > maxSize {
		return "", fmt.Errorf("file '%s' melebihi batas ukuran 5MB", fh.Filename)
	}

	folder := fmt.Sprintf("/coworking-space-booking-app-assets/%s", id)
	url, err := cldHelper.UploadFile(ctx, s.cld, fh, folder)
	if err != nil {
		return "", fmt.Errorf("gagal upload '%s': %w", fh.Filename, err)
	}

	return url, nil
}

func (s *VendorServiceImpl) UpdateVendor(ctx context.Context, vendorID uuid.UUID, vendor *dto.UpdateVendorRequestDto) (*dto.VendorDto, error) {

	existingVendor, err := s.repo.GetVendorByID(ctx, vendorID)
	if err != nil {
		return nil, err
	}

	if vendor.VendorName != nil {
		existingVendor.VendorName = *vendor.VendorName
	}
	if vendor.Address != nil {
		existingVendor.Address = *vendor.Address
	}
	if vendor.City != nil {
		existingVendor.City = *vendor.City
	}
	if vendor.PhoneNumber != nil {
		existingVendor.PhoneNumber = *vendor.PhoneNumber
	}
	if vendor.Email != nil {
		existingVendor.Email = *vendor.Email
	}
	if vendor.Description != nil {
		existingVendor.Description = *vendor.Description
	}

	result, err := s.repo.UpdateVendor(ctx, existingVendor)
	if err != nil {
		return nil, err
	}

	resultDto := mapper.ToVendorDto(result)
	return &resultDto, nil
}

func (s *VendorServiceImpl) DeleteVendor(ctx context.Context, vendorID uuid.UUID) error {
	return s.repo.DeleteVendor(ctx, vendorID)
}
