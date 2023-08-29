package models

import (
	"time"
)

// Video, 视频表
type Video struct {
	ID            uint      `gorm:"primarykey"`
	UserId        uint      `gorm:"not null; comment:作者ID"`
	PlayUrl       string    `gorm:"not null; comment:视频播放地址"`
	CoverUrl      string    `gorm:"not null; comment:视频封面地址"`
	FavoriteCount int64     `gorm:"not null; comment:点赞数量"`
	CommentCount  int64     `gorm:"not null; comment:视频的评论总数"`
	Description   string    `gorm:"not null; comment:视频描述"`
	CreatedAt     time.Time `gorm:"not null; comment:视频发布日期"`
	// 定义外键关系
	User User `gorm:"foreignKey:UserId; references:ID; comment:作者信息"`
}
