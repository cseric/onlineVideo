package vo

import "time"

// CollectVideoVo 收藏视频
type CollectVideoVo struct {
	Id          uint      `json:"id"`
	Cover       string    `json:"cover"`
	Title       string    `json:"title"`
	CollectTime time.Time `json:"collect_time"`
}

// ListVideoVo 首页视频列表
type ListVideoVo struct {
	Id     uint   `json:"id"`
	Cover  string `json:"cover"`
	Title  string `json:"title"`
	Uid    uint   `json:"uid"`
	Author string `json:"author"`
}

// UserVideoVo 用户视频
type UserVideoVo struct {
	Id         uint      `json:"id"`
	Cover      string    `json:"cover"`
	Title      string    `json:"title"`
	UploadTime time.Time `json:"upload_time"`
}

// PlayVideoVo 播放视频信息
type PlayVideoVo struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Cover      string    `json:"cover"`
	Path       string    `json:"path"`
	Brief      string    `json:"brief"`
	Uid        uint      `json:"uid"`
	UploadTime time.Time `json:"upload_time"`
	Author     string    `json:"author"`
	Avatar     string    `json:"avatar"`
	Sign       string    `json:"sign"`
}

// VideoVo 视频列表
type VideoVo struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Cover      string    `json:"cover"`
	Path       string    `json:"path"`
	Brief      string    `json:"brief"`
	Uid        uint      `json:"uid"`
	UploadTime time.Time `json:"upload_time"`
	Status     uint      `json:"status"`
	Remark     string    `json:"remark"`
}

// VideoListVo  视频列表
type VideoListVo struct {
	VideoVo
	InteractiveDataVo
}
