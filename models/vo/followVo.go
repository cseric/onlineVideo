package vo

// FollowVo 关注和粉丝信息
type FollowVo struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Sign     string `json:"sign"`
}
