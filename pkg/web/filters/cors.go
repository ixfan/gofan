package filters

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"net/http"
	"strings"
)

var AllowHeaders = "*"

func Cors(headers ...string) gin.HandlerFunc {
	var allowHeaders string
	if len(headers) <= 0 {
		allowHeaders = AllowHeaders
	} else {
		allowHeaders = strings.Join(headers, ",")
	}
	return func(context *gin.Context) {
		//过滤掉静态目录
		if strutil.Substr(context.Request.RequestURI, 0, 7) == "/static" {
			context.Next()
			return
		}
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", allowHeaders)
		context.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, OPTIONS, DELETE")
		context.Header("Access-Control-Allow-Credentials", "true")
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
