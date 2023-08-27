package models

import "gorm.io/gorm"

// IncrementWorkCount 增加用户的作品数。
func IncrementWorkCount(db *gorm.DB, userID uint) error {
	return db.Model(&User{}).Where("id = ?", userID).Update("work_count", gorm.Expr("work_count + ?", 1)).Error
}

// IncrementUserLikeCount 增加用户的点赞数。
func IncrementUserLikeCount(db *gorm.DB, userID uint) error {
	return db.Model(&User{}).Where("id = ?", userID).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
}

// DecrementUserLikeCount 减少用户的点赞数。
func DecrementUserLikeCount(db *gorm.DB, userID uint) error {
	return db.Model(&User{}).Where("id = ?", userID).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
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
