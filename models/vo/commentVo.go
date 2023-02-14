package vo

import "time"

// CommentInfoVo 评论信息
type CommentInfoVo struct {
	Id          uint      `json:"id"`
	Vid         uint      `json:"vid"`
	Uid         uint      `json:"uid"`
	Content     string    `json:"content"`
	CommentTime time.Time `json:"comment_time"`
	Title       string    `json:"title"`
	Username    string    `json:"username"`
}

// VideoCommentVo 视频评论
type VideoCommentVo struct {
	Id          uint      `json:"id"`
	Uid         uint      `json:"uid"`
	Content     string    `json:"content"`
	CommentTime time.Time `json:"comment_time"`
	Username    string    `json:"username"`
	Avatar      string    `json:"avatar"`
}
