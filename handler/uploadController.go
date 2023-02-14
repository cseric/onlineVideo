package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"onlineVideo/service"
	"onlineVideo/utils"
	"path"
	"strconv"
	"time"
)

const (
	MaxAvatarSize = 1024 * 1024 * 5
	MaxCoverSize = 1024 * 1024 * 5
	MaxVideoSize = 1024 * 1024 * 500
)

// ImgAllowSuffix 允许的图片文件类型
var ImgAllowSuffix = [...]string{".jpg", ".jpeg", ".png"}

// UploadAvatar 上传头像
func UploadAvatar(c *gin.Context) {
	avatar, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}
	// 验证文件类型
	suffix := path.Ext(avatar.Filename)
	if !isImgAllowSuffix(suffix) {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.FileTypeNotAllow,
		})
		return
	}
	// 验证文件大小
	if avatar.Size >= MaxAvatarSize {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.FileSizeOverLimit,
		})
		return
	}
	// 生成文件名 avt+当前时间戳+类型后缀
	avatarName := "avt" + strconv.FormatInt(time.Now().UnixNano(), 10) + suffix
	// 获取保存路径
	dst := viper.GetString("file.avatar") + avatarName
	// 保存文件
	err = c.SaveUploadedFile(avatar, dst)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.UploadFailed,
		})
		return
	}
	uid := c.GetUint("user_id")
	res := service.UploadAvatar(uid, avatarName)
	c.JSON(http.StatusOK, res)
}

// UploadCover 上传封面
func UploadCover(c *gin.Context) {
	cover, err := c.FormFile("cover")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}

	// 验证文件类型
	suffix := path.Ext(cover.Filename)
	if !isImgAllowSuffix(suffix) {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.FileTypeNotAllow,
		})
		return
	}

	// 验证文件大小（不超过5M）
	if cover.Size >= MaxCoverSize {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.FileSizeOverLimit,
		})
		return
	}
	// 生成文件名 cov+当前时间戳+类型后缀
	coverName := "cov" + strconv.FormatInt(time.Now().UnixNano(), 10) + suffix
	// 获取保存路径
	dst := viper.GetString("file.cover") + coverName
	// 保存文件
	err = c.SaveUploadedFile(cover, dst)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.UploadFailed,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.UploadSuccess,
		Data: gin.H{
			"coverPath": coverName,
		},
	})
}

// UploadVideo 上传视频
func UploadVideo(c *gin.Context) {
	video, err := c.FormFile("video")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.RequestError,
		})
		return
	}

	// 验证文件类型
	suffix := path.Ext(video.Filename)
	if suffix != ".mp4" {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.FileTypeNotAllow,
		})
		return
	}

	// 验证文件大小
	if video.Size >= MaxVideoSize {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.FileSizeOverLimit,
		})
		return
	}
	// 生成文件名 vid+当前时间戳+类型后缀
	videoName := "vid" + strconv.FormatInt(time.Now().UnixNano(), 10) + suffix
	// 获取保存路径
	dst := viper.GetString("file.video") + videoName
	// 保存文件
	err = c.SaveUploadedFile(video, dst)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			Code: utils.CodeFail,
			Msg: utils.UploadFailed,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.UploadSuccess,
		Data: gin.H{
			"videoPath": videoName,
		},
	})
}

// 验证图片文件类型是否符合要求
func isImgAllowSuffix(suffix string) bool {
	for _, allowSuffix := range ImgAllowSuffix {
		if suffix == allowSuffix {
			return true
		}
	}
	return false
}
