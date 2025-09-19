package middleware

import (
	"blog-main/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "请求头中缺少Auth信息",
			})
			ctx.Abort() // 阻止执行
			return
		}
		tokenParts := strings.SplitN(authHeader, " ", 2)
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "请求头中Auth格式有误",
			})
			ctx.Abort() // 阻止执行
			return
		}
		tokenString := tokenParts[1]
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "无效的Token",
			})
			ctx.Abort() // 阻止执行
			return
		}
		if claims.TokenType != "access" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "Token类型错误",
			})
			ctx.Abort()
			return
		}
		ctx.Set("userId", claims.UserId)
		ctx.Set("username", claims.Username)
		ctx.Next()
	}
}
