package services

import (
	"context"
	"errors"
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
	"github.com/sirupsen/logrus"
)

type VendorServiceImpl struct {
	repo repositories.VendorRepository
	cld  *cloudinary.Cloudinary
	log  *logrus.Logger
}

func NewVendorService(repo repositories.VendorRepository, cld *cloudinary.Cloudinary, log *logrus.Logger) services.VendorService {
	return &VendorServiceImpl{
		repo: repo,
		cld:  cld,
		log:  log,
	}
}

func (s *VendorServiceImpl) GetAllVendors(ctx context.Context, params dto.VendorFilterRequest) ([]dto.VendorDto, int, error) {
	s.log.WithFields(logrus.Fields{
		"name":           params.Name,
		"city":           params.City,
		"page_size":      params.PageSize,
		"offset":         params.Offset,
		"sort_by":        params.SortBy,
		"sort_direction": params.SortDirection,
	}).Info("fetching all vendors")

	vendors, total, err := s.repo.GetAllVendors(ctx, params.Name, params.City, params.PageSize, params.Offset, params.SortBy, params.SortDirection)
	if err != nil {
		s.log.WithField("error", err).Error("failed to fetch vendors from repository")
		return nil, 0, errors.New("terjadi kesalahan, coba lagi nanti")
	}

	result := mapper.ToVendorListDto(vendors)
	return result, total, nil
}

func (s *VendorServiceImpl) GetVendorByID(ctx context.Context, vendorID uuid.UUID) (*dto.VendorDto, error) {
	s.log.WithField("vendor_id", vendorID).Info("fetching vendor by id")

	vendor, err := s.repo.GetVendorByID(ctx, vendorID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"vendor_id": vendorID,
			"error":     err,
		}).Error("failed to fetch vendor from repository")
		return nil, errors.New("terjadi kesalahan, coba lagi nanti")
	}

	if vendor == nil {
		s.log.WithField("vendor_id", vendorID).Warn("vendor not found")
		return nil, errors.New("vendor tidak ditemukan")
	}

	result := mapper.ToVendorDto(vendor)
	return &result, nil
}

func (s *VendorServiceImpl) GetVendorsResourcesByVendorID(ctx context.Context, vendorID uuid.UUID) ([]dto.ResourceDto, error) {
	s.log.WithField("vendor_id", vendorID).Info("fetching vendor resources")

	result, err := s.repo.GetVendorsResourcesByVendorID(ctx, vendorID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"vendor_id": vendorID,
			"error":     err,
		}).Error("failed to fetch vendor resources from repository")
		return nil, errors.New("terjadi kesalahan, coba lagi nanti")
	}

	resultDto := mapper.ToVendorResourceListDto(result)
	return resultDto, nil
}

func (s *VendorServiceImpl) CreateVendor(ctx context.Context, profileImage *multipart.FileHeader, vendor *dto.CreateVendorRequestDto) (*dto.VendorDto, error) {
	s.log.WithField("owner_user_id", vendor.OwnerUserID).Info("creating vendor")

	// === 1. Parse UUID dulu (murah, sebelum upload) ===
	ownerUserIdUuid, err := uuid.Parse(vendor.OwnerUserID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"owner_user_id": vendor.OwnerUserID,
			"error":         err,
		}).Warn("invalid owner_user_id format")
		// User error — kembalikan pesan jelas
		return nil, errors.New("format owner_user_id tidak valid")
	}

	// === 2. Upload image (opsional) ===
	var imageURL *string
	if profileImage != nil {
		url, err := s.uploadProfileImage(ctx, profileImage, vendor.OwnerUserID)
		if err != nil {
			// User error dari validasi (format/size) — kembalikan pesan aslinya
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
		s.log.WithFields(logrus.Fields{
			"owner_user_id": vendor.OwnerUserID,
			"error":         err,
		}).Error("failed to create vendor in repository")

		// Rollback: hapus image dari Cloudinary jika DB gagal
		if imageURL != nil {
			if rollbackErr := cldHelper.DeleteFromCloudinary(ctx, s.cld, *imageURL); rollbackErr != nil {
				s.log.WithFields(logrus.Fields{
					"image_url": *imageURL,
					"error":     rollbackErr,
				}).Error("failed to rollback image upload after db error")
			}
		}
		return nil, errors.New("terjadi kesalahan, coba lagi nanti")
	}

	s.log.WithFields(logrus.Fields{
		"vendor_id":     result.VendorID,
		"owner_user_id": vendor.OwnerUserID,
	}).Info("vendor created successfully")

	resultDto := mapper.ToVendorDto(result)
	return &resultDto, nil
}

// uploadProfileImage — validasi + upload, file ditutup otomatis
func (s *VendorServiceImpl) uploadProfileImage(ctx context.Context, fh *multipart.FileHeader, id string) (string, error) {
	allowedExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	const maxSize = 5 * 1024 * 1024

	ext := strings.ToLower(filepath.Ext(fh.Filename))
	if !allowedExt[ext] {
		s.log.WithFields(logrus.Fields{
			"filename":  fh.Filename,
			"extension": ext,
		}).Warn("unsupported file format")
		// User error — kembalikan pesan jelas
		return "", fmt.Errorf("file '%s' format tidak didukung", fh.Filename)
	}

	if fh.Size > maxSize {
		s.log.WithFields(logrus.Fields{
			"filename": fh.Filename,
			"size":     fh.Size,
		}).Warn("file size exceeds limit")
		// User error — kembalikan pesan jelas
		return "", fmt.Errorf("file '%s' melebihi batas ukuran 5MB", fh.Filename)
	}

	folder := fmt.Sprintf("/coworking-space-booking-app-assets/%s", id)
	url, err := cldHelper.UploadFile(ctx, s.cld, fh, folder)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"filename": fh.Filename,
			"folder":   folder,
			"error":    err,
		}).Error("failed to upload image to cloudinary")
		// System error — pesan generik
		return "", errors.New("terjadi kesalahan saat upload gambar, coba lagi nanti")
	}

	return url, nil
}

// UpdateVendor memperbarui data vendor berdasarkan field yang dikirim (partial update).
// Jika profileImage dikirim, gambar baru diupload ke Cloudinary terlebih dahulu.
// Urutan proses: upload baru → update DB → hapus lama.
// Jika update DB gagal, gambar baru di-rollback (dihapus dari Cloudinary).
func (s *VendorServiceImpl) UpdateVendor(ctx context.Context, vendorID uuid.UUID, profileImage *multipart.FileHeader, vendor *dto.UpdateVendorRequestDto) (*dto.VendorDto, error) {
	s.log.WithField("vendor_id", vendorID).Info("updating vendor")

	// ── 1. Ambil data vendor yang sudah ada ──
	existingVendor, err := s.repo.GetVendorByID(ctx, vendorID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"vendor_id": vendorID,
			"error":     err,
		}).Error("failed to fetch vendor from repository")
		return nil, errors.New("terjadi kesalahan, coba lagi nanti")
	}

	if existingVendor == nil {
		s.log.WithField("vendor_id", vendorID).Warn("vendor not found")
		return nil, errors.New("vendor tidak ditemukan")
	}

	// ── 2. Upload image baru jika ada ──
	var newImageURL *string
	var oldImageURL *string

	if profileImage != nil {
		url, err := s.uploadProfileImage(ctx, profileImage, existingVendor.OwnerUserID.String())
		if err != nil {
			// Error sudah di-log di uploadProfileImage
			return nil, err
		}
		newImageURL = &url

		if existingVendor.ProfileImage != "" {
			oldImageURL = &existingVendor.ProfileImage
		}
	}

	// ── 3. Terapkan perubahan (partial update) ──
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
	result, err := s.repo.UpdateVendor(ctx, existingVendor)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"vendor_id": vendorID,
			"error":     err,
		}).Error("failed to update vendor in repository")

		// Rollback image baru yang sudah terupload
		if newImageURL != nil {
			if rollbackErr := cldHelper.DeleteFromCloudinary(ctx, s.cld, *newImageURL); rollbackErr != nil {
				s.log.WithFields(logrus.Fields{
					"image_url": *newImageURL,
					"error":     rollbackErr,
				}).Error("failed to rollback image upload after db error")
			}
		}
		return nil, errors.New("terjadi kesalahan, coba lagi nanti")
	}

	// ── 5. DB sukses → hapus image lama dari Cloudinary ──
	if oldImageURL != nil {
		if err := cldHelper.DeleteFromCloudinary(ctx, s.cld, *oldImageURL); err != nil {
			// Hanya log, tidak return error — hanya orphan file
			s.log.WithFields(logrus.Fields{
				"image_url": *oldImageURL,
				"error":     err,
			}).Warn("failed to delete old image from cloudinary, may cause orphan file")
		}
	}

	s.log.WithField("vendor_id", vendorID).Info("vendor updated successfully")

	resultDto := mapper.ToVendorDto(result)
	return &resultDto, nil
}

// DeleteVendor menghapus vendor dari database dan menghapus profile image dari Cloudinary.
// Urutan: hapus DB dulu → hapus Cloudinary.
// Jika hapus Cloudinary gagal, hanya di-log (tidak merusak data, hanya orphan file).
func (s *VendorServiceImpl) DeleteVendor(ctx context.Context, vendorID uuid.UUID) error {
	s.log.WithField("vendor_id", vendorID).Info("deleting vendor")

	// ── 1. Ambil data vendor ──
	existingVendor, err := s.repo.GetVendorByID(ctx, vendorID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"vendor_id": vendorID,
			"error":     err,
		}).Error("failed to fetch vendor from repository")
		return errors.New("terjadi kesalahan, coba lagi nanti")
	}

	if existingVendor == nil {
		s.log.WithField("vendor_id", vendorID).Warn("vendor not found")
		return errors.New("vendor tidak ditemukan")
	}

	// ── 2. Hapus dari DB ──
	err = s.repo.DeleteVendor(ctx, existingVendor.VendorID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"vendor_id": vendorID,
			"error":     err,
		}).Error("failed to delete vendor from repository")
		return errors.New("terjadi kesalahan, coba lagi nanti")
	}

	// ── 3. Hapus image dari Cloudinary ──
	// DB sudah terhapus, jangan return error — hanya orphan file jika gagal
	if existingVendor.ProfileImage != "" {
		if err := cldHelper.DeleteFromCloudinary(ctx, s.cld, existingVendor.ProfileImage); err != nil {
			s.log.WithFields(logrus.Fields{
				"vendor_id": vendorID,
				"image_url": existingVendor.ProfileImage,
				"error":     err,
			}).Warn("failed to delete image from cloudinary, may cause orphan file")
		}
	}

	s.log.WithField("vendor_id", vendorID).Info("vendor deleted successfully")

	return nil
}
