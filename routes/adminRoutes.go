package routes

import (
	"github.com/gin-gonic/gin"
	"onlineVideo/handler/control"
	"onlineVideo/midddleware"
	"onlineVideo/models"
)

// SetupAdminRoutes 配置管理员路由
func SetupAdminRoutes(r *gin.RouterGroup) {
	admin := r.Group("/admin")
	{
		// 管理员登录
		admin.POST("/login", control.AdminLogin)

		publicAuth := admin.Group("")
		publicAuth.Use(midddleware.AdminMiddleware(models.PublicAuth))
		{
			// 获取统计数据
			publicAuth.GET("/total_data", control.GetStatisticsData)
			// 修改管理员密码
			publicAuth.PUT("/update_pass", control.UpdateAdminPassword)
		}

		// 超级管理员
		superAdminAuth := admin.Group("")
		superAdminAuth.Use(midddleware.AdminMiddleware(models.SuperAdminAuth))
		{
			// 获取管理员列表
			superAdminAuth.GET("/admin_list",control.GetAdminList)
			// 增加管理员
			superAdminAuth.POST("/admin_add", control.AddAdmin)
			// 修改管理员
			superAdminAuth.PUT("/admin_update", control.UpdateAdmin)
			// 删除管理员
			superAdminAuth.DELETE("/admin_delete", control.DeleteAdmin)
		}

		// 管理员路由
		adminAuth := admin.Group("")
		adminAuth.Use(midddleware.AdminMiddleware(models.AdminAuth))
		{
			// 获取用户列表
			adminAuth.GET("/user_list", control.GetUserList)
			// 搜索用户
			adminAuth.GET("/user_search", control.SearchUser)
			// 删除用户
			adminAuth.DELETE("/user_delete", control.DeleteUser)

			// 获取视频列表
			adminAuth.GET("/video_list", control.GetVideoList)
			// 搜索视频
			adminAuth.GET("/video_search", control.SearchVideo)
			// 删除视频
			adminAuth.DELETE("/video_delete", control.DeleteVideo)

			// 获取评论列表
			adminAuth.GET("/comment_list", control.GetCommentList)
			// 搜索评论
			adminAuth.GET("/comment_search", control.SearchComment)
			// 删除评论
			adminAuth.DELETE("/comment_delete", control.DeleteComment)
		}

		// 审核员路由
		auditAuth := admin.Group("")
		auditAuth.Use(midddleware.AdminMiddleware(models.AuditAuth))
		{
			// 获取待审核视频列表
			auditAuth.GET("/audit_list", control.GetAuditList)
			// 审核通过
			auditAuth.PUT("/audit_pass", control.AuditVideoPass)
			// 审核不通过
			auditAuth.PUT("/audit_fail", control.AuditVideoFail)
		}
	}
}
