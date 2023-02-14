package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineVideo/models/dto"
	"onlineVideo/service"
	"onlineVideo/utils"
	"strconv"
)

// UploadVideoInfo 上传视频信息
func UploadVideoInfo(c *gin.Context) {
	var uploadVideo dto.UploadVideoDto
	err := c.ShouldBind(&uploadVideo)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	uid := c.GetUint("user_id")
	res := service.UploadVideoInfo(uid, &uploadVideo)
	c.JSON(http.StatusOK, res)
}

// GetCollectVideo 获取收藏视频
func GetCollectVideo(c *gin.Context) {
	uid := c.GetUint("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "5"))
	if page <= 0 || pageSize <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.PageOrPageSizeError,
		})
		return
	}
	res := service.GetCollectVideo(uid, page, pageSize)
	c.JSON(http.StatusOK, res)
}

// UpdateVideoInfo 更新视频信息
func UpdateVideoInfo(c *gin.Context) {
	var video dto.UpdateVideoDto
	err := c.ShouldBind(&video)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	uid := c.GetUint("user_id")
	res := service.UpdateVideoInfo(uid, &video)
	c.JSON(http.StatusOK, res)
}

// DeleteVideo 删除视频
func DeleteVideo(c *gin.Context) {
	var videoId dto.VideoIdDto
	err := c.ShouldBind(&videoId)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	if videoId.Id == 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.VideoNotExist,
		})
		return
	}
	uid := c.GetUint("user_id")
	res := service.DeleteVideo(videoId.Id, uid)
	c.JSON(http.StatusOK, res)
}

// GetUserVideo 获取用户视频
func GetUserVideo(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Query("uid"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "5"))
	if page <= 0 || pageSize <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.PageOrPageSizeError,
		})
		return
	}
	res := service.GetUserVideo(uint(uid), page, pageSize)
	c.JSON(http.StatusOK, res)
}

// GetVideoList 获取首页视频
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
	res := service.GetVideoList(page, pageSize)
	c.JSON(http.StatusOK, res)
}

// GetRecommend 获取推荐视频
func GetRecommend(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", "5"))
	if size <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	res := service.GetRecommend(size)
	c.JSON(http.StatusOK, res)
}

// GetNewest 获取最新视频
func GetNewest(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", "5"))
	if size <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	res := service.GetNewest(size)
	c.JSON(http.StatusOK, res)
}

// GetUploadVideo 获取上传的视频
func GetUploadVideo(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "5"))
	if page <= 0 || pageSize <= 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.PageOrPageSizeError,
		})
		return
	}
	uid := c.GetUint("user_id")
	res := service.GetUploadVideo(uid, page, pageSize)
	c.JSON(http.StatusOK, res)
}

// GetVideoById 根据id获取视频
func GetVideoById(c *gin.Context) {
	vid, _ := strconv.Atoi(c.Query("vid"))
	if vid == 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	uid := c.GetUint("user_id")
	res := service.GetVideoById(uint(vid), uid)
	c.JSON(http.StatusOK, res)
}

// GetVideoInteractive 获取视频交互信息
func GetVideoInteractive(c *gin.Context) {
	vid, _ := strconv.Atoi(c.Query("vid"))
	if vid == 0 {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	uid := c.GetUint("user_id")
	res := service.GetVideoInteractiveInfo(uint(vid), uid)
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
	res := service.SearchVideo(keyword, page, pageSize)
	c.JSON(http.StatusOK, res)
}