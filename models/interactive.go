package models

import "time"

type Interactive struct {
	Id          uint      `gorm:"primarykey"`
	Uid         uint      `gorm:"not null;"`                        // 用户id
	Vid         uint      `gorm:"not null"`                         // 视频id
	IsCollect   bool      `gorm:"is_collect;default:false;"`        // 是否收藏
	IsLike      bool      `gorm:"is_like;default:false;"`           // 是否点赞
	CollectTime time.Time `gorm:"upload_time;default:'1970-01-01'"` // 收藏时间
}
