package midddleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors 跨域中间件
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		// 允许任何域的访问
		ctx.Header("Access-Control-Allow-Origin", "*")
		// 设置缓存时间
		ctx.Header("Access-Control-Max-Age", "86400")
		// 设置允许的请求头字段
		ctx.Header("Access-Control-Allow-Headers",
			"Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		// 设置服务器允许接收的方法
		ctx.Header("Access-Control-Allow-Methods", "*")
		// 设置浏览器可以解析的请求头字段
		ctx.Header("Access-Control-Expose-Headers",
			"Content-Length, Access-Control-Allow-Origin, " +
			"Access-Control-Allow-Headers, Content-Type")
		// 设置浏览器可以传递cookie等校验字段
		ctx.Header("Access-Control-Allow-Credentials", "true")
		// 放行OPTIONS请求
		if method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
		}
		ctx.Next()
	}
}
