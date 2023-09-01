package models

import (
	"time"
)

// Video, 视频表
type Video struct {
	ID            uint      `gorm:"primarykey"`
	UserId        uint      `gorm:"not null; comment:作者ID           type:INT"`
	PlayUrl       string    `gorm:"not null; comment:视频播放地址;    type:VARCHAR(255)"`
	CoverUrl      string    `gorm:"not null; comment:视频封面地址;    type:VARCHAR(255)"`
	FavoriteCount int64     `gorm:"not null; comment:点赞数量;        type:BIGINT"`
	CommentCount  int64     `gorm:"not null; comment:视频的评论总数;  type:BIGINT"`
	Description   string    `gorm:"not null; comment:视频描述;        type:TEXT"`
	CreatedAt     time.Time `gorm:"not null; comment:视频发布日期;    type:DATETIME"`
	// 定义外键关系
	User User `gorm:"foreignKey:UserId; references:ID; comment:作者信息"`
}
