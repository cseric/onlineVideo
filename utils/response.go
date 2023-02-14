package utils

import "github.com/gin-gonic/gin"

// 结果码
const (
	CodeSuccess      = 2000 // 执行成功
	CodeFail         = 4000 // 执行失败
	CodeNoEnoughAuth = 4001 // 权限不足
	CodeServerError  = 5000 // 服务异常
)

type Response struct {
	Code int    `json:"code"` // 结果码
	Msg  string `json:"msg"`  // 消息
	Data gin.H  `json:"data"` // 数据
}

// 消息
const (
	UserNotExist        = "用户不存在"
	UserIsExist         = "用户已存在"
	VideoNotExist       = "视频不存在"
	RequestError        = "请求错误"
	PasswordError       = "密码错误"
	LoginSuccess        = "登录成功"
	RegisterSuccess     = "注册成功"
	ServerError         = "服务异常"
	NoAuth              = "请求头中auth为空"
	AuthFormatError     = "请求头中auth格式有误"
	ValidToken          = "无效的token"
	UpdateSuccess       = "修改成功"
	GetDataSuccess      = "获取数据成功"
	PageOrPageSizeError = "请求页码或请求页数有误"
	NoKeyword           = "搜索内容不能为空"
	SearchSuccess       = "搜索成功"
	DeleteSuccess       = "删除成功"
	CommentSuccess      = "评论成功"
	FollowSuccess       = "关注成功"
	UnFollowSuccess     = "取关成功"
	UploadSuccess       = "上传成功"
	UploadFailed        = "上传失败"
	FileTypeNotAllow    = "文件类型不符"
	FileSizeOverLimit   = "文件过大"
	AddSuccess          = "添加成功"
	NoEnoughAuth        = "权限不足"
)
