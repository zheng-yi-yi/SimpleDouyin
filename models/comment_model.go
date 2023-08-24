package models

import "time"

//  @Description: 评论表
type Comment struct {
	ID        uint      `json:"id"                    gorm:"primarykey;autoIncrement"`
	UserId    uint      `json:"user_id,omitempty"     gorm:"not null; comment:发布评论的用户ID"`
	VideoId   uint      `json:"video_id,omitempty"    gorm:"not null; comment:评论所属视频ID"`
	Content   string    `json:"content,omitempty"     gorm:"not null; comment:评论内容"`
	CreatedAt time.Time `json:"create_date,omitempty" gorm:"not null; comment:评论发布日期"`
	Cancel    uint      `json:"cancel,omitempty"      gorm:"not null; default:0; comment:默认评论发布为0，取消后为1"`
	User      User      `json:"user,omitempty"        gorm:"foreignKey:UserId; references:ID; comment:评论所属用户"`
	Video     Video     `json:"video,omitempty"       gorm:"foreignKey:VideoId; references:ID; comment:评论所属视频"`
}
