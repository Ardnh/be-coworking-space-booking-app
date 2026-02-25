package cloudinary

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// uploadFile — buka, upload, dan TUTUP file per iterasi
func UploadFile(ctx context.Context, cld *cloudinary.Cloudinary, fh *multipart.FileHeader, folder string) (string, error) {
	file, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	url, err := UploadToCloudinary(ctx, cld, file, folder)
	if err != nil {
		return "", err
	}
	return url, nil
}

func UploadToCloudinary(ctx context.Context, cld *cloudinary.Cloudinary, file multipart.File, folder string) (string, error) {

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: folder,
	})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}

func DeleteFromCloudinary(ctx context.Context, cld *cloudinary.Cloudinary, publicID string) error {
	_, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	return err
}

// rollbackUploads — hapus file dari Cloudinary jika ada kegagalan
func RollbackUploads(ctx context.Context, cld *cloudinary.Cloudinary, urls []string) {
	for _, url := range urls {
		_ = DeleteFromCloudinary(ctx, cld, url)
	}
}
