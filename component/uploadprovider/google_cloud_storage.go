package upldprv

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"bds/common"

	"cloud.google.com/go/storage"
)

type CloudStorageProvider struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	domain     string
}

func NewCloudStorageProvider(cl *storage.Client, projectID, bucketName, domain string) *CloudStorageProvider {
	return &CloudStorageProvider{
		cl:         cl,
		projectID:  projectID,
		bucketName: bucketName,
		domain:     domain,
	}
}

func (provider *CloudStorageProvider) UploadFile(ctx context.Context, file multipart.File, dst string) (*common.Image, error) {
	wc := provider.cl.Bucket(provider.bucketName).Object(provider.domain + dst).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return nil, err
	}
	if err := wc.Close(); err != nil {
		return nil, err
	}

	img := &common.Image{
		Url:       fmt.Sprintf("%s/%s", provider.domain, dst),
		CloudName: "google cloud",
	}

	return img, nil
}
