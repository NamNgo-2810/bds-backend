package ginhm

import (
	"net/http"
	"strconv"

	"bds/common"
	"bds/component/appctx"
	hmbiz "bds/module/home/business"
	hmmdl "bds/module/home/model"
	hmstrg "bds/module/home/storage"

	"github.com/gin-gonic/gin"
)

func UpdateHome(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data hmmdl.HomeUpdate
		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := hmstrg.NewSQLStore(db)
		biz := hmbiz.NewUpdateHomeBiz(store)

		if err := biz.UpdateHome(context.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
