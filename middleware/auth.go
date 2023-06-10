package middleware

import (
	"fmt"
	"net/http"
	"webapp1/util"

	"github.com/gin-gonic/gin"
)

func APIKey() gin.HandlerFunc {
	const API_KEY = "artinyaapabangmessi"
	const HEADER_KEY = "secret_key"

	return func(ctx *gin.Context) {
		// ambil value secret_key dari header
		// cek apakah cocok dengan API_KEY
		// jika cocok, next
		// jika tidak cocok, return 401 Unauthorized

		apiKeySecret := ctx.GetHeader(HEADER_KEY)

		if apiKeySecret != API_KEY {
			fmt.Println(apiKeySecret)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.BuildResponse("error", nil))
			return
		}

		ctx.Next()
	}
}