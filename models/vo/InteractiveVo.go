package vo

// InteractiveInfoVo 视频交互信息
type InteractiveInfoVo struct {
	IsLike    bool  `json:"is_like"`
	IsCollect bool  `json:"is_collect"`
	IsFollow  bool  `json:"is_follow"`
	InteractiveDataVo
}

// InteractiveDataVo 视频点赞收藏数据
type InteractiveDataVo struct {
	Likes   int64 `json:"likes"`
	Collect int64 `json:"collect"`
}
