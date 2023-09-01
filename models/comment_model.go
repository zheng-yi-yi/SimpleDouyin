package models

import "time"

//  Comment. 评论表
type Comment struct {
	ID        uint      `gorm:"primarykey; comment:评论id"`
	UserId    uint      `gorm:"not null;   comment:发布评论的用户id;  type:INT"`
	VideoId   uint      `gorm:"not null;   comment:评论所属视频id;    type:INT"`
	Content   string    `gorm:"not null;   comment:评论内容;          type:VARCHAR(255)"`
	CreatedAt time.Time `gorm:"not null;   comment:评论发布日期;      type:DATETIME"`
	// 定义外键关系
	User  User  `gorm:"foreignKey:UserId; references:ID; comment:评论所属用户"`
	Video Video `gorm:"foreignKey:VideoId; references:ID; comment:评论所属视频"`
}
