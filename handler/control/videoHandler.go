package control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineVideo/models"
	"onlineVideo/models/dto"
	"onlineVideo/service"
	"onlineVideo/utils"
	"strconv"
)

// GetVideoList 获取视频列表
func GetVideoList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "5"))
	if page <= 0 || pageSize <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.PageOrPageSizeError,
		})
		return
	}
	res := service.AdminGetVideoList(page, pageSize, models.CodeAuditPass)
	c.JSON(http.StatusOK, res)
}

// GetAuditList 获取审核视频列表
func GetAuditList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "5"))
	if page <= 0 || pageSize <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.PageOrPageSizeError,
		})
		return
	}
	res := service.AdminGetVideoList(page, pageSize, models.CodeNotAudit)
	c.JSON(http.StatusOK, res)
}

// SearchVideo 搜索视频
func SearchVideo(c *gin.Context) {
	keyword := c.Query("keyword")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "5"))
	if len(keyword) == 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.NoKeyword,
		})
		return
	}
	if page <= 0 || pageSize <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.PageOrPageSizeError,
		})
		return
	}
	res := service.AdminSearchVideo(keyword, page, pageSize)
	c.JSON(http.StatusOK, res)
}

// DeleteVideo 删除视频
func DeleteVideo(c *gin.Context) {
	var videoId dto.AdminIdDto
	err := c.ShouldBind(&videoId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	res := service.AdminDeleteVideo(videoId.Id)
	c.JSON(http.StatusOK, res)
}

// AuditVideoPass 视频审核通过
func AuditVideoPass(c *gin.Context) {
	var videoId dto.AdminIdDto
	err := c.ShouldBind(&videoId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	res := service.AdminAuditVideoPass(videoId.Id)
	c.JSON(http.StatusOK, res)
}

// AuditVideoFail 视频审核不通过
func AuditVideoFail(c *gin.Context) {
	var auditVideo dto.AuditVideoVo
	err := c.ShouldBind(&auditVideo)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	res := service.AdminAuditVideoFail(&auditVideo)
	c.JSON(http.StatusOK, res)
}
