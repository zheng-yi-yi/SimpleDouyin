package models

import "gorm.io/gorm"

//  @Description: 视频表
type Video struct {
	gorm.Model
	UserId        uint   `json:"user_id,omitempty"        gorm:"type: int; not null; comment:作者ID"`
	PlayUrl       string `json:"play_url,omitempty"       gorm:"type: text; not null; comment:视频播放地址"`
	CoverUrl      string `json:"cover_url,omitempty"      gorm:"type: text; not null; comment:视频封面地址"`
	FavoriteCount int64  `json:"favorite_count,omitempty" gorm:"type: int; default: 0; comment:点赞数量"`
	CommentCount  int64  `json:"comment_count,omitempty"  gorm:"type: int; default: 0; comment:评论数量"`
	Description   string `json:"description,omitempty"    gorm:"type: text; comment:视频描述"`
	User          User   `json:"user,omitempty"           gorm:"foreignKey:UserId; references:ID; comment:视频作者"`
	IsFavorite    bool   `json:"is_favorite" gorm:"-"`
}
