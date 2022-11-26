package ginupld

import (
	"mime/multipart"
	"net/http"

	"bds/common"
	"bds/component/appctx"
	upldbiz "bds/module/upload/business"

	"github.com/gin-gonic/gin"
)

func Upload(appCtx appctx.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := ctx.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer func(file multipart.File) {
			file.Close()
		}(file)

		dataBytes := make([]byte, fileHeader.Size)
		_, err = file.Read(dataBytes)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		biz := upldbiz.NewUploadBusiness(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(ctx.Request.Context(), file, folder, fileHeader.Filename)
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
