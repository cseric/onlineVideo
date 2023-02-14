package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"onlineVideo/common"
	"onlineVideo/models"
	"onlineVideo/models/dto"
	"onlineVideo/models/vo"
	"onlineVideo/utils"
	"time"
)

// --------------------------------用户--------------------------------

// UploadVideoInfo 用户上传视频
func UploadVideoInfo(uid uint, uploadVideo *dto.UploadVideoDto) utils.Response {
	DB := common.GetDB()
	newVideo := models.Video{
		Title: uploadVideo.Title,
		Cover: uploadVideo.Cover,
		Path: uploadVideo.Path,
		Brief: uploadVideo.Brief,
		Uid: uid,
		UploadTime: time.Now(),
		Status: models.CodeNotAudit,
	}
	DB.Create(&newVideo)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.UploadSuccess,
	}
}

// GetCollectVideo 用户获取收藏
func GetCollectVideo(uid uint, page int, pageSize int) utils.Response {
	var collectVideo []vo.CollectVideoVo
	var total int64
	DB := common.GetDB()
	// 获取收藏数
	DB.Model(&models.Video{}).Joins("join interactive on interactive.vid = video.id").
		Where("status = ? and interactive.uid = ? and is_collect = true", models.CodeAuditPass, uid).Count(&total)
	// 分页获取收藏数据
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(&models.Video{}).Select("video.id, cover, title, interactive.collect_time").
		Joins("join interactive on interactive.vid = video.id").
		Where("status = ? and interactive.uid = ? and is_collect = true", models.CodeAuditPass, uid).
		Find(&collectVideo)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.GetDataSuccess,
		Data: gin.H{
			"total": total,
			"collects": collectVideo,
		},
	}
}

// UpdateVideoInfo 用户修改视频信息
func UpdateVideoInfo(uid uint, video *dto.UpdateVideoDto) utils.Response {
	DB := common.GetDB()
	// 验证是否是自己的视频
	if !isOwnVideo(DB, uid, video.Id) {
		return utils.Response{
			Code: utils.CodeFail,
			Msg: "修改失败",
		}
	}
	// 修改视频
	DB.Model(&models.Video{}).Where("id = ? and uid = ?", video.Id, uid).
		Updates(map[string]interface{}{
			"title": video.Title,
			"cover": video.Cover,
			"path": video.Path,
			"brief": video.Brief,
			"status": models.CodeNotAudit,
	})
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.UpdateSuccess,
	}
}

// DeleteVideo 用户删除视频
func DeleteVideo(vid uint, uid uint) utils.Response {
	DB := common.GetDB()
	// 验证是否是自己的视频
	if !isOwnVideo(DB, uid, vid) {
		return utils.Response{
			Code: utils.CodeFail,
			Msg: "删除失败",
		}
	}
	DB.Where("id = ?", vid).Delete(&models.Video{})	// 删除视频
	DB.Where("vid = ?", vid).Delete(&models.Comment{})	// 删除视频评论
	DB.Where("vid = ?", vid).Delete(&models.Interactive{})	// 删除视频交互数据
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.DeleteSuccess,
	}
}

// GetVideoList 获取首页视频列表
func GetVideoList(page int, pageSize int) utils.Response {
	var video []vo.ListVideoVo
	var total int64
	DB := common.GetDB()
	DB.Model(&models.Video{}).Joins("join user on user.id = video.uid").Where("video.status = ?", models.CodeAuditPass).Count(&total)
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(&models.Video{}).Select("video.id, video.cover, video.title, video.uid, user.username as author").
		Joins("join user on user.id = video.uid").Where("video.status = ?", models.CodeAuditPass).
		Order("video.upload_time DESC").Find(&video)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.GetDataSuccess,
		Data: gin.H{
			"total": total,
			"videos": video,
		},
	}
}

// GetRecommend 获取推荐视频
func GetRecommend(size int) utils.Response {
	DB := common.GetDB()
	var video []vo.ListVideoVo
	sqlStr := "select video.id, video.cover, video.title, video.uid, user.username as author from video " +
		"join user on user.id = video.uid where video.status = ? and video.id in " +
		"(select vid from interactive where is_like = true group by vid order by count(is_like) desc) " +
		"limit ?"
	DB.Raw(sqlStr, models.CodeAuditPass, size).Find(&video)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.GetDataSuccess,
		Data: gin.H{
			"recommends": video,
		},
	}
}

// GetNewest 获取最新视频
func GetNewest(size int) utils.Response {
	DB := common.GetDB()
	var video []vo.ListVideoVo
	DB.Model(&models.Video{}).Select("video.id, video.cover, video.title, video.uid, user.username as author").
		Joins("join user on user.id = video.uid").Where("video.status = ?", models.CodeAuditPass).
		Order("video.upload_time DESC").Limit(size).Find(&video)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.GetDataSuccess,
		Data: gin.H{
			"newest": video,
		},
	}
}

// GetUserVideo 用户获取个人主页视频
func GetUserVideo(uid uint, page int, pageSize int) utils.Response {
	DB := common.GetDB()

	// 验证用户是否存在
	var user models.User
	DB.Where("id = ?", uid).First(&user)
	if user.Id == 0 {
		return utils.Response{
			Code: utils.CodeFail,
			Msg: utils.UserNotExist,
		}
	}
	var total int64
	var userVideo []vo.UserVideoVo
	DB.Model(&models.Video{}).Where("status = ? and uid = ?", models.CodeAuditPass, uid).Count(&total)
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(&models.Video{}).Select("id, cover, title, upload_time").
		Where("status = ? and uid = ?", models.CodeAuditPass, uid).
		Order("upload_time DESC").Find(&userVideo)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.GetDataSuccess,
		Data: gin.H{
			"total": total,
			"userVideo": userVideo,
		},
	}
}

// GetUploadVideo 用户获取上传的视频
func GetUploadVideo(uid uint, page int, pageSize int) utils.Response {
	DB := common.GetDB()
	var total int64
	var videos []vo.VideoVo
	DB.Model(&models.Video{}).Where("uid = ?", uid).Count(&total)
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(&models.Video{}).Where("uid = ?", uid).Find(&videos)

	videosData := GetVideosData(videos)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.GetDataSuccess,
		Data: gin.H{
			"total": total,
			"videos": videosData,
		},
	}
}

// GetVideoById 用户根据视频id获取视频
func GetVideoById(vid uint, uid uint) utils.Response {
	DB := common.GetDB()
	var playVideo vo.PlayVideoVo
	// 获取视频信息
	DB.Model(&models.Video{}).
		Select("video.id, video.title, video.cover, video.path, video.brief, video.uid, video.upload_time, user.username as author, user.avatar, user.sign").
		Joins("join user on user.id = video.uid").Where("video.id = ?", vid).First(&playVideo)
	isOwn := isOwnVideo(DB, uid, vid)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.GetDataSuccess,
		Data: gin.H{
			"video": playVideo,
			"isOwn": isOwn,
		},
	}
}

// GetVideoInteractiveInfo 获取用户和视频的交互状态
func GetVideoInteractiveInfo(vid uint, uid uint) utils.Response {
	DB := common.GetDB()
	InteractiveData := GetInteractiveInfo(DB, uid, vid)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.GetDataSuccess,
		Data: gin.H {
			"interactive": InteractiveData,
		},
	}
}

// SearchVideo 用户搜索视频
func SearchVideo(keyword string, page int, pageSize int) utils.Response {
	var video []vo.ListVideoVo
	var total int64
	DB := common.GetDB()
	keyword = "%" + keyword + "%"
	DB.Model(&models.Video{}).Joins("join user on user.id = video.uid").Where("video.status = ? and title like ?", models.CodeAuditPass, keyword).Count(&total)
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(&models.Video{}).Select("video.id, video.cover, video.title, video.uid, user.username as author").
		Joins("join user on user.id = video.uid").Where("video.status = ? and title like ?", models.CodeAuditPass, keyword).
		Order("video.upload_time DESC").Find(&video)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.GetDataSuccess,
		Data: gin.H{
			"total": total,
			"videos": video,
		},
	}
}

// --------------------------------管理员--------------------------------

// AdminGetVideoList 管理员获取视频列表
func AdminGetVideoList(page int, pageSize int, status int) utils.Response {
	var total int64
	var videos []vo.VideoVo

	DB := common.GetDB()
	DB.Model(&models.Video{}).Where("status = ?", status).Count(&total)

	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(&models.Video{}).Where("status = ?", status).Find(&videos)

	videosData := GetVideosData(videos)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg:  utils.GetDataSuccess,
		Data: gin.H{
			"total":  total,
			"videos": videosData,
		},
	}
}

// AdminSearchVideo 管理员搜索视频
func AdminSearchVideo(keyword string, page int, pageSize int) utils.Response {
	var total int64
	var videos []vo.VideoVo

	DB := common.GetDB()
	keyword = "%" + keyword + "%"
	DB.Model(&models.Video{}).
		Where("status = ? and ( id like ? or title like ? or uid like ? )", models.CodeAuditPass, keyword, keyword, keyword).Count(&total)
	DB = DB.Limit(pageSize).Offset((page - 1) * pageSize)
	DB.Model(&models.Video{}).
		Where("status = ? and ( id like ? or title like ? or uid like ? )", models.CodeAuditPass, keyword, keyword, keyword).Scan(&videos)
	videosData := GetVideosData(videos)
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.SearchSuccess,
		Data: gin.H{
			"total": total,
			"videos": videosData,
		},
	}
}

// AdminDeleteVideo 管理员删除视频
func AdminDeleteVideo(videoId uint) utils.Response {
	DB := common.GetDB()
	DB.Delete(&models.Video{}, videoId)
	DB.Where("vid = ?", videoId).Delete(&models.Comment{})
	DB.Where("vid = ?", videoId).Delete(&models.Interactive{})
	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.DeleteSuccess,
	}
}

// AdminAuditVideoPass 管理员审核视频通过
func AdminAuditVideoPass(videoId uint) utils.Response {
	DB := common.GetDB()
	DB.Model(&models.Video{}).Where("id = ?", videoId).Updates(map[string]interface{}{
		"status": models.CodeAuditPass,
		"remark": "无",
	})

	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.UpdateSuccess,
	}
}

// AdminAuditVideoFail 管理员审核视频不通过
func AdminAuditVideoFail(auditVideo *dto.AuditVideoVo) utils.Response {
	DB := common.GetDB()
	DB.Model(&models.Video{}).Select("status", "remark").Where("id = ?", auditVideo.Id).Updates(map[string]interface{}{"status": models.CodeAuditFail, "remark": auditVideo.Remark})

	return utils.Response{
		Code: utils.CodeSuccess,
		Msg: utils.UpdateSuccess,
	}
}

// --------------------------------功能性函数--------------------------------

// GetInteractiveInfo 获取交互数据
func GetInteractiveInfo(DB *gorm.DB, uid uint, vid uint) vo.InteractiveInfoVo {
	var video models.Video
	DB.First(&video, vid)

	like, collect := isLikeAndCollect(DB, uid, vid)
	follow := isFollow(DB, uid, video.Uid)
	data := GetInteractiveData(DB, vid)

	return vo.InteractiveInfoVo{
		IsLike: like,
		IsCollect: collect,
		IsFollow: follow,
		InteractiveDataVo: data,
	}
}

// GetVideosData 获取视频数据
func GetVideosData(videos []vo.VideoVo) []vo.VideoListVo {
	var videoList []vo.VideoListVo
	DB := common.GetDB()
	var data vo.InteractiveDataVo
	for i := 0; i < len(videos); i++ {
		data = GetInteractiveData(DB, videos[i].Id)
		videoList = append(videoList, vo.VideoListVo{
			VideoVo:           videos[i],
			InteractiveDataVo: data,
		})
	}
	return videoList
}

// GetInteractiveData 获取视频收藏和点赞数据
func GetInteractiveData(DB *gorm.DB, vid uint) vo.InteractiveDataVo {
	var likes int64
	var collect int64
	DB.Model(&models.Interactive{}).Where("vid = ? and is_like = true", vid).Count(&likes)
	DB.Model(&models.Interactive{}).Where("vid = ? and is_collect = true", vid).Count(&collect)
	return vo.InteractiveDataVo{
		Likes: likes,
		Collect: collect,
	}
}

// isLikeAndCollect 获取点赞和收藏状态
func isLikeAndCollect(DB *gorm.DB, uid uint, vid uint) (bool, bool) {
	var interactive models.Interactive
	DB.Model(&models.Interactive{}).Where("uid = ? and vid = ?", uid, vid).First(&interactive)
	if interactive.Id != 0 {
		return interactive.IsLike, interactive.IsCollect
	}
	return false, false
}

// isOwnVideo 查询视频是否为当前用户的
func isOwnVideo(DB *gorm.DB, uid uint, vid uint) bool {
	var video models.Video
	DB.Where("uid = ? and id = ?", uid, vid).First(&video)
	if video.Id != 0 {
		return true
	}
	return false
}
