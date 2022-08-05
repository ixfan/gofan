package filters

import (
	"github.com/gin-gonic/gin"
)

func LimitAccess(headers ...string) gin.HandlerFunc {
	return func(context *gin.Context) {
		
		context.Next()
	}
}
