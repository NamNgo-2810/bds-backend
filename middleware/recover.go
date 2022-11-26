package middleware

import (
	"bds/common"
	"bds/component/appctx"

	"github.com/gin-gonic/gin"
)

func Recover(ac appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				context.Header("Content-Type", "application/json")
				if appErr, ok := err.(*common.AppError); ok {
					context.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
				}

				appErr := common.ErrInternal(err.(error))
				context.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
			}
		}()

		context.Next()
	}
}
