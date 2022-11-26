package upldbiz

import (
	"context"
	"fmt"
	"image"
	"io"
	"log"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"bds/common"
	upldprv "bds/component/uploadprovider"
	upldmdl "bds/module/upload/model"
)

type CreateImageStorage interface {
	CreateImage(ctx context.Context, data *common.Image) error
}

type UploadBusiness struct {
	provider upldprv.UploadProvider
	imgStore CreateImageStorage
}

func NewUploadBusiness(provider upldprv.UploadProvider, imgStore CreateImageStorage) *UploadBusiness {
	return &UploadBusiness{provider: provider, imgStore: imgStore}
}

func (business UploadBusiness) Upload(ctx context.Context, file multipart.File, folder, fileName string) (*common.Image, error) {
	w, h, err := getImageDimension(file)
	if err != nil {
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := business.provider.UploadFile(ctx, file, fmt.Sprintf("%s/%s", folder, fileName))
	if err != nil {
		return nil, upldmdl.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err:", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
