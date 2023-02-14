package routes

import (
	"github.com/gin-gonic/gin"
	"onlineVideo/handler"
	"onlineVideo/midddleware"
)

func SetupUserRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		// 用户登录
		user.POST("/login", handler.UserLogin)
		// 用户注册
		user.POST("/register",handler.UserRegister)
		// 验证用户是否存在
		user.PATCH("/check", handler.CheckUserIsExit)

		userAuth := user.Group("")
		userAuth.Use(midddleware.AuthMiddleware())
		{
			// 获取用户信息
			userAuth.GET("/info", handler.UserInfo)
			// 修改用户信息
			userAuth.PUT("/update/info", handler.UpdateUserInfo)
			// 修改用户密码
			userAuth.PUT("/update/pass", handler.UpdateUserPass)
			// 获取其他用户信息
			userAuth.GET("/other/info", handler.GetUserInfo)
			// 修改用户头像
			userAuth.PUT("/update/avatar", handler.UploadAvatar)
		}
	}
}
