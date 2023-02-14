package models

import "time"

type Video struct {
	Id         uint      `gorm:"primarykey"`                // 视频id
	Title      string    `gorm:"type:varchar(50);not null"` // 视频标题
	Cover      string    `gorm:"not null"`                  // 视频封面
	Path       string    `gorm:"not null"`                  // 视频路径
	Brief      string    `gorm:"type:varchar(100);"`        // 视频介绍
	Uid        uint      `gorm:"not null"`                  // 用户（作者）id
	UploadTime time.Time `gorm:"upload_time;not null"`      // 上传时间
	Status     uint      `gorm:"default:0"`                 // 视频状态（0-待审核 1-审核未通过 2-审核通过）
	Remark     string    `gorm:"default:'无'"`               // 备注信息，用来填写审核信息
}

const (
	CodeNotAudit  = 0
	CodeAuditFail = 1
	CodeAuditPass = 2
)
