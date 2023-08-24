package models

import (
	"time"

	"gorm.io/gorm"
)

// @Description: 视频表
type Video struct {
	ID            uint   `json:"id"                       gorm:"primarykey"`
	UserId        uint   `json:"user_id,omitempty"        gorm:"type: int; not null; comment:作者ID"`
	PlayUrl       string `json:"play_url,omitempty"       gorm:"type: text; not null; comment:视频播放地址"`
	CoverUrl      string `json:"cover_url,omitempty"      gorm:"type: text; not null; comment:视频封面地址"`
	FavoriteCount int64  `json:"favorite_count,omitempty" gorm:"type: int; default: 0; comment:点赞数量"`
	CommentCount  int64  `json:"comment_count,omitempty"  gorm:"type: int; default: 0; comment:视频的评论总数"`
	Description   string `json:"description,omitempty"    gorm:"type: text; comment:视频描述"`
	IsFavorite    bool   `json:"is_favorite"              gorm:"-"`
	User          User   `json:"author,omitempty"         gorm:"foreignKey:UserId; references:ID; comment:视频作者信息"`
	CreatedAt     time.Time
}

// GetVideoCount 根据用户ID查询其拥有的视频数量。
func GetVideoCount(db *gorm.DB, userId uint) (int64, error) {
	var videoCount int64
	result := db.Model(&Video{}).Where("user_id = ?", userId).Count(&videoCount)
	if result.Error != nil {
		return 0, result.Error
	}
	return videoCount, nil
}

// IncrementWorkCount 增加用户的作品数。
func IncrementWorkCount(db *gorm.DB, userID uint) error {
	return db.Model(&User{}).Where("id = ?", userID).Update("work_count", gorm.Expr("work_count + ?", 1)).Error
}
