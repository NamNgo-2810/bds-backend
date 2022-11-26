package ginhm

import (
	"net/http"

	"bds/common"
	"bds/component/appctx"
	hmbiz "bds/module/home/business"
	hmmdl "bds/module/home/model"
	hmstrg "bds/module/home/storage"

	"github.com/gin-gonic/gin"
)

func CreateHome(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		// requester := context.MustGet(common.CurrentUser).(common.Requester)
		var data hmmdl.HomeCreate
		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := hmstrg.NewSQLStore(db)
		biz := hmbiz.NewCreateHomeBiz(store)

		if err := biz.CreateHome(context.Request.Context(), &data); err != nil {
			panic(err)
		}

		// data.Mask(false)

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
