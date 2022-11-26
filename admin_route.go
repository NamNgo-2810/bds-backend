package main

import (
	appctx "bds/component/appctx"
	"bds/middleware"
	ginhm "bds/module/home/transport/ginhome"
	ginupld "bds/module/upload/transport/ginupload"

	"github.com/gin-gonic/gin"
)

func setupAdminRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group("/admin", middleware.RoleRequired(appContext, "admin"))

	{
		admin.POST("/home", ginhm.CreateHome(appContext))
		admin.PATCH("/home", ginhm.UpdateHome(appContext))
		admin.POST("/upload", ginupld.Upload(appContext))
	}
}
