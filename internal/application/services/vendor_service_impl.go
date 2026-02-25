package services

import (
	"context"
	"fmt"
	"log"
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

// UpdateVendor memperbarui data vendor berdasarkan field yang dikirim (partial update).
// Jika profileImage dikirim, gambar baru diupload ke Cloudinary terlebih dahulu.
// Urutan proses: upload baru → update DB → hapus lama.
// Jika update DB gagal, gambar baru di-rollback (dihapus dari Cloudinary).
func (s *VendorServiceImpl) UpdateVendor(
	ctx context.Context,
	vendorID uuid.UUID,
	profileImage *multipart.FileHeader,
	vendor *dto.UpdateVendorRequestDto,
) (*dto.VendorDto, error) {

	// ── 1. Ambil data vendor yang sudah ada ──
	existingVendor, err := s.repo.GetVendorByID(ctx, vendorID)
	if err != nil {
		return nil, err
	}

	// ── 2. Upload image baru jika ada ──
	// Simpan URL lama untuk dihapus setelah DB sukses.
	// Upload dulu baru hapus, agar image lama tetap aman jika upload gagal.
	var newImageURL *string
	var oldImageURL *string

	if profileImage != nil {
		url, err := s.uploadProfileImage(ctx, profileImage, existingVendor.OwnerUserID.String())
		if err != nil {
			return nil, err
		}
		newImageURL = &url

		// ⚠️ FIX: simpan URL lama, bukan URL baru
		if existingVendor.ProfileImage != "" {
			oldImageURL = &existingVendor.ProfileImage
		}
	}

	// ── 3. Terapkan perubahan (partial update) ──
	// Hanya field yang dikirim (tidak nil) yang diperbarui.
	if newImageURL != nil {
		existingVendor.ProfileImage = *newImageURL
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

	// ── 4. Simpan ke DB ──
	// Jika gagal, rollback image baru yang sudah terupload di Cloudinary.
	result, err := s.repo.UpdateVendor(ctx, existingVendor)
	if err != nil {
		if newImageURL != nil {
			_ = cldHelper.DeleteFromCloudinary(ctx, s.cld, *newImageURL)
		}
		return nil, err
	}

	// ── 5. DB sukses → hapus image lama dari Cloudinary ──
	// Dilakukan terakhir karena jika gagal hapus, hanya jadi orphan file
	// (tidak merusak data). Bisa dibersihkan berkala.
	if oldImageURL != nil {
		_ = cldHelper.DeleteFromCloudinary(ctx, s.cld, *oldImageURL)
	}

	resultDto := mapper.ToVendorDto(result)
	return &resultDto, nil
}

// DeleteVendor menghapus vendor dari database dan menghapus profile image dari Cloudinary.
// Urutan: hapus DB dulu → hapus Cloudinary.
// Jika hapus Cloudinary gagal, hanya di-log (tidak merusak data, hanya orphan file).
func (s *VendorServiceImpl) DeleteVendor(ctx context.Context, vendorID uuid.UUID) error {

	// ── 1. Ambil data vendor ──
	existingVendor, err := s.repo.GetVendorByID(ctx, vendorID)
	if err != nil {
		return err
	}

	// ── 2. Hapus dari DB ──
	err = s.repo.DeleteVendor(ctx, existingVendor.VendorID)
	if err != nil {
		return err
	}

	// ── 3. Hapus image dari Cloudinary ──
	// DB sudah terhapus, jadi jangan return error.
	// Jika gagal, hanya jadi orphan file di Cloudinary.
	if existingVendor.ProfileImage != "" {
		err = cldHelper.DeleteFromCloudinary(ctx, s.cld, existingVendor.ProfileImage)
		if err != nil {
			log.Printf("WARNING: gagal hapus image Cloudinary untuk vendor %s: %v", vendorID, err)
		}
	}

	return nil
}
