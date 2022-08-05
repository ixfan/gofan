package filters

import (
	"github.com/gin-gonic/gin"
	"github.com/ixfan/gofan/pkg/web/auth"
	"net/http"
)

func AccessToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		user, err := auth.ParseToken(token)
		if err != nil || user.UserId <= 0 {
			context.JSON(http.StatusOK, map[string]interface{}{
				"code":    403,
				"message": "Token过期或者无效，请重新登录",
			})
			context.Abort()
		} else {
			context.Set("AuthUser", user)
			context.Next()
		}
	}
}
