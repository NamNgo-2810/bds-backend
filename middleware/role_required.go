package middleware

import (
	"errors"

	"bds/common"
	"bds/component/appctx"

	"github.com/gin-gonic/gin"
)

func RoleRequired(appCtx appctx.AppContext, allowRoles ...string) func(*gin.Context) {
	return func(context *gin.Context) {
		u := context.MustGet(common.CurrentUser).(common.Requester)

		hasFound := false

		for _, item := range allowRoles {
			if u.GetRole() == item {
				hasFound = true
				break
			}
		}

		if !hasFound {
			panic(common.ErrNoPermission(errors.New("invalid role")))
		}

		context.Next()
	}
}
