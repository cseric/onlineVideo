package midddleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineVideo/common"
	"onlineVideo/models"
	"onlineVideo/utils"
	"strings"
)

// AuthMiddleware 用户鉴权中间件
func AuthMiddleware() func(c *gin.Context) {
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

		// 验证用户是否存在
		DB := common.GetDB()
		var user models.User
		DB.First(&user, claims.Id)
		if user.Id == 0 {
			c.JSON(http.StatusUnauthorized, utils.Response{
				Code: utils.CodeFail,
				Msg: utils.UserNotExist,
			})
			c.Abort()
			return
		}

		// 将id存入上下文中
		c.Set("user_id", claims.Id)
		c.Next()
	}
}
