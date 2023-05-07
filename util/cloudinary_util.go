package util

import (
	"context"
	"mime/multipart"
	"stage01-project-backend/config"
	"stage01-project-backend/httperror"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var c = config.CloudConfig

func InitiateCloudinary() (*cloudinary.Cloudinary, error) {

	cld, err := cloudinary.NewFromParams(c.CloudName, c.ApiKey, c.ApiSecret)
	if err != nil {
		return nil, httperror.ErrFailedInitiateCloudinary
	}

	cld.Config.URL.Secure = true
	return cld, nil
}

func UploadImage(cld *cloudinary.Cloudinary, input *multipart.FileHeader) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	file, err := input.Open()
	if err != nil {
		return "", httperror.ErrFailedOpenFile
	}
	defer file.Close()

	//upload file
	uploadParam := uploader.UploadParams{
		Folder: c.FolderName,
	}

	uploadResult, err := cld.Upload.Upload(ctx, file, uploadParam)
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}

func GetPublicIdFromUrl(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1][:strings.Index(parts[len(parts)-1], ".")]
}

func DeleteImage(cld *cloudinary.Cloudinary, publicId string) error {

	delParams := uploader.DestroyParams{
		PublicID: publicId,
	}

	_, err := cld.Upload.Destroy(context.Background(), delParams)
	if err != nil {
		return httperror.ErrFailedToDeleteImage
	}

	return nil
}
