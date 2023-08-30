package models

import (
	"github.com/zheng-yi-yi/simpledouyin/config"
	"gorm.io/gorm"
)

// 判断登录用户是否关注了视频作者
func IsFollow(fromUserId, toUserId uint) bool {
	// 查找存在的关注关系
	var existingRelation Relation
	if err := config.Database.Where("from_user_id = ? AND to_user_id = ? AND cancel = 0", fromUserId, toUserId).First(&existingRelation).Error; err != nil {
		return false
	}
	return true
}

// 判断用户是否点赞当前视频
func IsFavorite(userId, videoId uint) bool {
	// 创建一个 Favorite 结构体实例，用于存储查询结果
	var favorite Favorite
	// 在数据库中查找匹配的点赞记录
	result := config.Database.Where("user_id = ? AND video_id = ? AND status = 1", userId, videoId).First(&favorite)
	// 检查是否找到匹配的点赞记录
	return result.Error == nil
}

// IncrementWorkCount 增加用户的作品数。
func IncrementWorkCount(userID uint) error {
	return config.Database.Model(&User{}).Where("id = ?", userID).Update("work_count", gorm.Expr("work_count + ?", 1)).Error
}

// IncrementUserLikeCount 增加用户的点赞数。
func IncrementUserLikeCount(userID uint) error {
	return config.Database.Model(&User{}).Where("id = ?", userID).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
}

// IncrementAuthorTotalFavorited 根据视频作者的id来将其获赞总数加一
func IncrementAuthorTotalFavorited(userID uint) error {
	return config.Database.Model(&User{}).Where("id = ?", userID).Update("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error
}

// DecrementAuthorTotalFavorited 根据视频作者的id来将其获赞总数减一
func DecrementAuthorTotalFavorited(userID uint) error {
	return config.Database.Model(&User{}).Where("id = ?", userID).Update("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error
}

// DecrementUserLikeCount 减少用户的点赞数。
func DecrementUserLikeCount(userID uint) error {
	return config.Database.Model(&User{}).Where("id = ?", userID).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
}

// GetVideoCount 根据用户ID查询其拥有的视频数量。
func GetVideoCount(userId uint) (int64, error) {
	var videoCount int64
	result := config.Database.Model(&Video{}).Where("user_id = ?", userId).Count(&videoCount)
	if result.Error != nil {
		return 0, result.Error
	}
	return videoCount, nil
}

// IncrementCommentCount 增加指定视频的评论数（加一）
func IncrementCommentCount(videoID uint) error {
	result := config.Database.Model(&Video{}).Where("id = ?", videoID).Update("comment_count", gorm.Expr("comment_count + ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DecreaseCommentCount 减少指定视频的评论数（减一）
func DecreaseCommentCount(videoID uint) error {
	result := config.Database.Model(&Video{}).Where("id = ?", videoID).Update("comment_count", gorm.Expr("comment_count - ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAuthorIDForVideo 根据视频ID获取视频的作者ID
func GetAuthorIDForVideo(videoID uint) (uint, error) {
	var authorID uint
	if err := config.Database.Model(&Video{}).Select("user_id").Where("id = ?", videoID).Scan(&authorID).Error; err != nil {
		return 0, err
	}
	return authorID, nil
}

// IncrementVideoLikeCount 增加视频获赞数
func IncrementVideoLikeCount(videoID uint) error {
	result := config.Database.Model(&Video{}).Where("id = ?", videoID).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DecrementVideoLikeCount 减少视频获赞数
func DecrementVideoLikeCount(videoID uint) error {
	result := config.Database.Model(&Video{}).Where("id = ?", videoID).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FetchData , 通过用户ID获取用户信息
func FetchData(userId uint) (User, error) {
	var user User
	err := config.Database.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
