package routes

import (
	"github.com/gin-gonic/gin"
	"onlineVideo/handler"
	"onlineVideo/midddleware"
)

func SetupVideoRoutes(r *gin.RouterGroup) {
	video := r.Group("/video")
	video.Use(midddleware.AuthMiddleware())
	{
		// 根据视频id获取视频（视频播放）
		video.GET("/get", handler.GetVideoById)
		// 获取视频的交互状态和数据
		video.GET("/data", handler.GetVideoInteractive)
		// 根据用户id获取视频（用户主页视频列表）
		video.GET("/user/get", handler.GetUserVideo)
		// 获取视频列表（首页视频列表）
		video.GET("/list", handler.GetVideoList)
		// 获取推荐视频
		video.GET("/recommend", handler.GetRecommend)
		// 获取最新视频
		video.GET("/newest", handler.GetNewest)
		// 获取已上传的视频（视频管理）
		video.GET("/upload/list", handler.GetUploadVideo)
		// 上传视频信息
		video.POST("/upload", handler.UploadVideoInfo)
		// 获取收藏
		video.GET("/collect", handler.GetCollectVideo)
		// 修改视频信息
		video.PUT("/update", handler.UpdateVideoInfo)
		// 删除视频
		video.DELETE("/delete", handler.DeleteVideo)
		// 搜索视频
		video.GET("/search", handler.SearchVideo)
	}
}
