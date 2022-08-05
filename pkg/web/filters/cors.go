package filters

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Cors(headers ...string) gin.HandlerFunc {
	var allowHeaders string
	if len(headers) <= 0 {
		allowHeaders = "*"
	} else {
		allowHeaders = strings.Join(headers, ",")
	}
	return func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", allowHeaders)
		context.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, OPTIONS, DELETE")
		context.Header("Access-Control-Allow-Credentials", "false")
		context.Header("Cache-Control", "no-cache, private")
		context.Header("Content-Type", "application/json")
		if strings.ToUpper(context.Request.Method) == "OPTIONS" {
			context.JSON(http.StatusOK, nil)
			context.Abort()
			return
		}
		context.Next()
	}
}
