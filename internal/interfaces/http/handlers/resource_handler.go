package handlers

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	http "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ResourceHandlers struct {
	vendorService services.ResourceService
	validator     *validator.Validate
	log           *logrus.Logger
}

func NewResourceHandler(vendorService services.ResourceService, validator *validator.Validate, log *logrus.Logger) *ResourceHandlers {
	return &ResourceHandlers{
		vendorService: vendorService,
		validator:     validator,
		log:           log,
	}
}

func (h *ResourceHandlers) GetReviewsByResourceId(c *fiber.Ctx) error {

	// id := c.Params("resourceId", "")
	// if id == "" {
	// 	return http.NewErrorResponse(c, fiber.StatusBadRequest, "ID is required", nil)
	// }

	// resourceIdUUID, err := uuid.Parse(id)
	// if err != nil {
	// 	return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	// }

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get vendor", nil)
}

func (h *ResourceHandlers) GetResourceById(c *fiber.Ctx) error {
	// id := c.Params("resourceId", "")
	// if id == "" {
	// 	return http.NewErrorResponse(c, fiber.StatusBadRequest, "ID is required", nil)
	// }

	// resourceIdUUID, err := uuid.Parse(id)
	// if err != nil {
	// 	return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	// }

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get vendor", nil)
}

func (h *ResourceHandlers) GetResource(c *fiber.Ctx) error {

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully get vendor", nil)
}

func (h *ResourceHandlers) CreateResource(c *fiber.Ctx) error {

	// // ── 1. Parse field teks ──────────────────────────────────────────────────
	// name := strings.TrimSpace(c.FormValue("resource_name"))
	// resourceTypeId := strings.TrimSpace(c.FormValue("resource_type_id"))
	// location := strings.TrimSpace(c.FormValue("location"))
	// description := strings.TrimSpace(c.FormValue("description"))

	// capacityStr := strings.TrimSpace(c.FormValue("capacity"))
	// priceStr := strings.TrimSpace(c.FormValue("price"))

	// // Validasi field wajib

	// // ── 2. Parse multiple file gambar ────────────────────────────────────────
	// form, err := c.MultipartForm()
	// if err != nil {
	// 	return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	// }

	// files := form.File["images"] // key "images" dari form-data
	// if len(files) == 0 {
	// 	return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	// }

	// // ── 3. Validasi & simpan setiap file ─────────────────────────────────────
	// allowedExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	// const maxSize = 5 * 1024 * 1024 // 5 MB per file

	// var savedImages []string

	// for _, file := range files {
	// 	// Cek ekstensi
	// 	ext := strings.ToLower(filepath.Ext(file.Filename))
	// 	if !allowedExt[ext] {
	// 		return http.NewErrorResponse(c, fiber.StatusBadRequest, fmt.Sprintf("Ekstensi file '%s' tidak diizinkan", ext), nil)
	// 	}

	// 	// Cek ukuran file
	// 	if file.Size > maxSize {
	// 		return http.NewErrorResponse(c, fiber.StatusBadRequest, fmt.Sprintf("File '%s' melebihi batas ukuran 5MB", file.Filename), nil)
	// 	}

	// 	// Generate nama file unik
	// 	uniqueName := fmt.Sprintf("%d_%s%s",
	// 		time.Now().UnixNano(),
	// 		strings.ReplaceAll(name, " ", "_"),
	// 		ext,
	// 	)
	// 	savePath := filepath.Join("uploads", uniqueName)

	// 	// Simpan file (upload to cloudinary)
	// 	if err := c.SaveFile(file, savePath); err != nil {
	// 		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	// 	}
	// }

	// ── 4. Buat objek produk & kembalikan respons ────────────────────────────
	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully create resource", nil)
}

func (h *ResourceHandlers) UpdateResource(c *fiber.Ctx) error {

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully update resource", nil)
}

func (h *ResourceHandlers) DeleteResource(c *fiber.Ctx) error {

	return http.NewSuccessResponse(c, fiber.StatusOK, "Successfully delete resource", nil)
}
