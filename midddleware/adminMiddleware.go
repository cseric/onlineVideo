package midddleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineVideo/common"
	"onlineVideo/models"
	"onlineVideo/utils"
	"strings"
)

// AdminMiddleware 管理员鉴权中间件
func AdminMiddleware(authority uint) func(c *gin.Context) {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")

		// 判断请求头是否有token
		if authToken == "" {
			c.JSON(http.StatusUnauthorized, utils.Response{
				Code: utils.CodeFail,
				Msg: utils.NoAuth,
			})
			c.Abort()
			return
		}

		// 验证token格式
		parts := strings.SplitN(authToken, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, utils.Response{
				Code: utils.CodeFail,
				Msg: utils.AuthFormatError,
			})
			c.Abort()
			return
		}

		// 解析和验证token是否有效
		token, claims, err := common.ParseToken(parts[1])
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, utils.Response{
				Code: utils.CodeFail,
				Msg: utils.ValidToken,
			})
			c.Abort()
			return
		}

		DB := common.GetDB()
		var admin models.Admin
		DB.First(&admin, claims.Id)
		if admin.Id == 0 {
			c.JSON(http.StatusUnauthorized, utils.Response{
				Code: utils.CodeFail,
				Msg: utils.UserNotExist,
			})
			c.Abort()
			return
		}
		if admin.Authority < authority {
			c.JSON(http.StatusUnauthorized, utils.Response{
				Code: utils.CodeNoEnoughAuth,
				Msg: utils.NoEnoughAuth,
			})
			c.Abort()
			return
		}

		// 将id存入上下文中
		c.Set("admin_id", claims.Id)
		c.Next()
	}
}
