package models

import "time"

type Comment struct {
	Id          uint      `gorm:"primarykey"`
	Vid         uint      `gorm:"not null;"`     // 视频id
	Uid         uint      `gorm:"not null;"`     // 用户id
	Content     string    `gorm:"not null;"`     // 评论内容
	CommentTime time.Time `gorm:"comment_time;"` // 评论时间
}
