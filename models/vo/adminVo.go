package vo

// AdminInfoVo 管理员信息
type AdminInfoVo struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// AdminListVo 管理员列表
type AdminListVo struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Authority uint   `json:"authority"`
}
