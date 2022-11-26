package upldprv

import (
	"context"
	"mime/multipart"

	"bds/common"
)

type UploadProvider interface {
	UploadFile(ctx context.Context, file multipart.File, fileName string) (*common.Image, error)
}
