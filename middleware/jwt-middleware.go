package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"net/http"
	"golangapi/util"
)

func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		accessToken := strings.Replace(token, "Bearer ", "", 1)	
		if ok := util.Parse(accessToken); !ok {
			ctx.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}