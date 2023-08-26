package services

import (
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

type FavoriteService struct {
}

// IsFavorite ，判断用户是否点赞了视频
func (s *FavoriteService) IsFavorite(userId, videoId uint) bool {
	// 创建一个 Favorite 结构体实例，用于存储查询结果
	var favorite models.Favorite

	// 在数据库中查找匹配的点赞记录
	result := config.DB.Where("user_id = ? AND video_id = ? AND status = 1", userId, videoId).First(&favorite)

	// 检查是否找到匹配的点赞记录
	return result.Error == nil
}

// AddLike ，判断用户是否点赞了该视频
func (s *FavoriteService) AddLike(userId, videoId uint) error {
	// 获取数据库连接实例
	db := config.DB

	// 检查之前是否有记录存在
	existingLike := models.Favorite{}
	result := db.Where("user_id = ? AND video_id = ?", userId, videoId).First(&existingLike)
	if result.Error == nil { // 找到现有记录
		if existingLike.Status == 0 { // 之前是取消点赞状态
			// 更新状态为点赞
			result := db.Model(&existingLike).Update("status", 1)
			if result.Error != nil {
				return result.Error
			}
			return nil
		}
		// 已经存在点赞记录，不需要再次添加
		return nil
	}

	// 创建点赞记录实例
	newLike := models.Favorite{
		UserId:  userId,
		VideoId: videoId,
		Status:  1, // 点赞状态
	}

	// 插入新的点赞记录
	if err := db.Create(&newLike).Error; err != nil {
		return err
	}

	// 获取该点赞视频的用户id
	author_id, err := models.GetAuthorIDForVideo(db, uint(videoId))
	if err != nil {
		return err
	}
	// 成功点赞后，调用 IncrementUserLikeCount 函数，将用户的点赞数加一
	if err := models.IncrementUserLikeCount(db, author_id); err != nil {
		return err
	}
	// 成功点赞后，调用 IncrementVideoLikeCount 函数，将视频的获赞数加一
	if err := models.IncrementVideoLikeCount(db, uint(videoId)); err != nil {
		return err
	}

	return nil
}

// CancelLike ，取消点赞操作
func (s *FavoriteService) CancelLike(userId, videoId uint) error {
	// 获取数据库连接实例
	db := config.DB

	// 创建 Favorite 对象用于存储查询结果
	var favorite models.Favorite

	// 在数据库中查找匹配的点赞记录
	err := db.Where("user_id = ? AND video_id = ? AND status = 1", userId, videoId).First(&favorite).Error
	if err != nil {
		return err // 如果发生错误，返回错误信息
	}

	// 将点赞状态设为 0，表示取消点赞
	favorite.Status = 0

	// 更新数据库中的点赞记录
	err = db.Save(&favorite).Error
	if err != nil {
		return err // 如果更新发生错误，返回错误信息
	}

	// 获取该点赞视频的用户id
	author_id, err := models.GetAuthorIDForVideo(db, uint(videoId))
	if err != nil {
		return err
	}

	// 成功取消点赞后，调用 DecrementUserLikeCount 函数，将用户的点赞数减一
	if err := models.DecrementUserLikeCount(db, author_id); err != nil {
		return err
	}

	// 成功取消点赞后，调用 IncrementVideoLikeCount 函数，将视频的获赞数加一
	if err := models.DecrementVideoLikeCount(db, uint(videoId)); err != nil {
		return err
	}

	return nil // 操作成功，返回 nil 表示没有错误
}
