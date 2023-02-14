package control

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineVideo/models/dto"
	"onlineVideo/service"
	"onlineVideo/utils"
	"strconv"
)

// AdminLogin 管理员登录
func AdminLogin(c *gin.Context) {
	var loginAdmin dto.AdminLoginDto
	err := c.ShouldBind(&loginAdmin)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	res := service.AdminLogin(&loginAdmin)
	c.JSON(http.StatusOK, res)
}

// UpdateAdminPassword 修改管理员密码
func UpdateAdminPassword(c *gin.Context) {
	var adminPass dto.AdminPasswordDto
	id := c.GetUint("admin_id")
	err := c.ShouldBind(&adminPass)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	res := service.UpdateAdminPassword(id, &adminPass)
	c.JSON(http.StatusOK, res)
}

// GetStatisticsData 获取统计数据
func GetStatisticsData(c *gin.Context) {
	res := service.GetStatisticsData()
	c.JSON(http.StatusOK, res)
}

// GetAdminList 获取管理员列表
func GetAdminList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "5"))
	if page <= 0 || pageSize <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.PageOrPageSizeError,
		})
		return
	}
	res := service.GetAdminList(page, pageSize)
	c.JSON(http.StatusOK, res)
}

// AddAdmin 添加管理员
func AddAdmin(c *gin.Context) {
	var addAdmin dto.AddAdminDto
	err := c.ShouldBind(&addAdmin)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	res := service.AddAdmin(&addAdmin)
	c.JSON(http.StatusOK, res)
}

// UpdateAdmin 修改管理员
func UpdateAdmin(c *gin.Context) {
	var updateAdmin dto.UpdateAdminDto
	err := c.ShouldBind(&updateAdmin)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	res := service.UpdateAdmin(&updateAdmin)
	c.JSON(http.StatusOK, res)
}

// DeleteAdmin 删除管理员
func DeleteAdmin(c *gin.Context) {
	var adminId dto.AdminIdDto
	err := c.ShouldBind(&adminId)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	res := service.DeleteAdmin(adminId.Id)
	c.JSON(http.StatusOK, res)
}