package models

import "time"

//  Comment. 评论表
type Comment struct {
	ID        uint      `gorm:"primarykey; comment:评论id"`
	UserId    uint      `gorm:"not null;   comment:发布评论的用户id"`
	VideoId   uint      `gorm:"not null;   comment:评论所属视频id"`
	Content   string    `gorm:"not null;   comment:评论内容"`
	CreatedAt time.Time `gorm:"not null;   comment:评论发布日期"`
	// Cancel    uint      `gorm:"not null;   comment:默认评论发布为0，取消后为1"`
	// 定义外键关系
	User  User  `gorm:"foreignKey:UserId; references:ID; comment:评论所属用户"`
	Video Video `gorm:"foreignKey:VideoId; references:ID; comment:评论所属视频"`
}
