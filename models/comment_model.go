package models

import "gorm.io/gorm"

//  @Description: 评论表
type Comment struct {
	gorm.Model
	Content  string `json:"content,omitempty"    gorm:"type:text; not null; comment:评论内容"`
	VideoId  uint   `json:"video_id,omitempty"   gorm:"type:int; not null; comment:评论所属视频ID"`
	ToUserId uint   `json:"to_user_id,omitempty" gorm:"type:int; comment:评论回复用户ID"`
	UserId   uint   `json:"user_id,omitempty"    gorm:"type:int; not null; comment:评论所属用户ID"`
	User     User   `json:"user,omitempty"       gorm:"foreignKey:UserId; references:ID; comment:评论所属用户"`
	Video    Video  `json:"video,omitempty"      gorm:"foreignKey:VideoId; references:ID; comment:评论所属视频"`
}
