package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"onlineVideo/handler"
	"onlineVideo/midddleware"
)

// SetupRoutes 初始化路由
func SetupRoutes() *gin.Engine {
	// 设置gin框架运行模式为debug
	gin.SetMode(viper.GetString("app.server_mode"))
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), midddleware.Cors())

	v1 := router.Group("/api/v1")
	{
		// 管理员路由
		SetupAdminRoutes(v1)
		// 用户路由
		SetupUserRoutes(v1)
		// 视频路由
		SetupVideoRoutes(v1)

		// 点赞/收藏
		interactive := v1.Group("/interactive")
		interactive.Use(midddleware.AuthMiddleware())
		{
			// 点赞
			interactive.POST("/like/add", handler.AddLike)
			// 取消点赞
			interactive.PUT("/like/cancel", handler.CancelLike)
			// 收藏
			interactive.POST("/collect/add", handler.AddCollect)
			// 取消收藏
			interactive.PUT("/collect/cancel",handler.CancelCollect)
		}

		// 关注/粉丝
		follow := v1.Group("/follow")
		follow.Use(midddleware.AuthMiddleware())
		{
			// 关注
			follow.POST("/add", handler.AddFollow)
			// 取消关注
			follow.DELETE("/cancel", handler.CancelFollow)
			// 关注列表
			follow.GET("/list", handler.GetFollowList)
			// 粉丝列表
			follow.GET("/fans", handler.GetFansList)
			// 关注状态获取
			follow.GET("/status", handler.GetFollowStatus)
			// 统计数据
			follow.GET("/data", handler.GetFollowData)
		}

		// 评论
		comment := v1.Group("/comment")
		comment.Use(midddleware.AuthMiddleware())
		{
			// 获取评论（根据视频id获取）
			comment.GET("/get", handler.GetVideoComment)
			// 删除评论
			comment.DELETE("/delete", handler.DeleteVideoComment)
			// 发送评论
			comment.POST("/send", handler.AddComment)
		}

		// 上传视频/封面
		upload := v1.Group("/upload")
		upload.Use(midddleware.AuthMiddleware())
		{
			// 上传封面
			upload.POST("/cover", handler.UploadCover)
			// 上传视频
			upload.POST("/video", handler.UploadVideo)
		}
	}

	// 放行静态资源
	router.Static("/res/avatar", "./upload/avatar")
	router.Static("/res/cover", "./upload/cover")
	router.Static("/res/video", "./upload/video")

	router.Static("/user", "./web/user")
	router.Static("/admin", "./web/admin")

	return router
}
