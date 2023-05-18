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
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		if origin != "" {
			context.Header("Access-Control-Allow-Origin", origin) // 可将将 * 替换为指定的域名
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			context.Header("Access-Control-Allow-Headers", allowHeaders)
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			context.Header("Access-Control-Allow-Credentials", "true")
		} else {
			context.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			context.Header("Access-Control-Allow-Headers", allowHeaders)
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			context.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
