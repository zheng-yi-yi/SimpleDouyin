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

// IncrementCommentCount 增加指定视频的评论数（加一）
func IncrementCommentCount(db *gorm.DB, videoID uint) error {
	result := db.Model(&Video{}).Where("id = ?", videoID).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DecreaseCommentCount 减少指定视频的评论数（减一）
func DecreaseCommentCount(db *gorm.DB, videoID uint) error {
	result := db.Model(&Video{}).Where("id = ?", videoID).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAuthorIDForVideo 根据视频ID获取视频的作者ID
func GetAuthorIDForVideo(db *gorm.DB, videoID uint) (uint, error) {
	var authorID uint
	if err := db.Model(&Video{}).Select("user_id").Where("id = ?", videoID).Scan(&authorID).Error; err != nil {
		return 0, err
	}
	return authorID, nil
}

// IncrementVideoLikeCount 增加视频获赞数
func IncrementVideoLikeCount(db *gorm.DB, videoID uint) error {
	result := db.Model(&Video{}).Where("id = ?", videoID).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DecrementVideoLikeCount 减少视频获赞数
func DecrementVideoLikeCount(db *gorm.DB, videoID uint) error {
	result := db.Model(&Video{}).Where("id = ?", videoID).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}
