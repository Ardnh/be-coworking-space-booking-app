package handlers

import (
	"strings"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/application/dto"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/domain/services"
	http "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/responses"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ResourceHandlers struct {
	service   services.ResourceService
	validator *validator.Validate
	log       *logrus.Logger
}

func NewResourceHandler(vendorService services.ResourceService, validator *validator.Validate, log *logrus.Logger) *ResourceHandlers {
	return &ResourceHandlers{
		service:   vendorService,
		validator: validator,
		log:       log,
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

	// ── 1. Parse field teks ──────────────────────────────────────────────────
	// String
	vendorId := strings.TrimSpace(c.FormValue("vandor_id"))
	resourceTypeId := strings.TrimSpace(c.FormValue("resource_type_id"))
	resourceName := strings.TrimSpace(c.FormValue("resource_name"))
	location := strings.TrimSpace(c.FormValue("location"))
	description := strings.TrimSpace(c.FormValue("description"))
	operationTimeFrom := strings.TrimSpace(c.FormValue("operation_time_from"))
	operationTimeTo := strings.TrimSpace(c.FormValue("operation_time_to"))
	endDate := strings.TrimSpace(c.FormValue("end_date"))

	// Number
	capacityStr := strings.TrimSpace(c.FormValue("capacity"))
	pricePerUnitStr := strings.TrimSpace(c.FormValue("price_per_unit"))

	// Array
	blockedDateStr := strings.TrimSpace(c.FormValue("blocked_date"))

	// Validasi field wajib
	req := dto.CreateResourceRequestDto{}

	// ── 2. Parse multiple file gambar ────────────────────────────────────────
	form, err := c.MultipartForm()
	if err != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	files := form.File["images"]
	if len(files) == 0 {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, "At least one image is required", nil)
	}

	result, errCreate := h.service.CreateResource(c.Context(), req, files)
	if errCreate != nil {
		return http.NewErrorResponse(c, fiber.StatusBadRequest, errCreate.Error(), nil)
	}

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
	// 	uniqueName := fmt.Sprintf("%d_%s_%s%s",
	// 		time.Now().UnixNano(),
	// 		uuid.New().String()[:8], // 8 char random
	// 		strings.ReplaceAll(resourceName, " ", "_"),
	// 		ext,
	// 	)
	// 	savePath := filepath.Join("uploads", uniqueName)

	// 	// 	// Simpan file (upload to cloudinary)
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
