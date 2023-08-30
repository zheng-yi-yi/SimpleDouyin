package services

import (
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
	"gorm.io/gorm"
)

type FavoriteService struct {
}

// IsFavorite ，判断用户是否点赞了视频
func (s *FavoriteService) IsFavorite(userId, videoId uint) bool {
	// 创建一个 Favorite 结构体实例，用于存储查询结果
	var favorite models.Favorite

	// 在数据库中查找匹配的点赞记录
	result := config.Database.Where("user_id = ? AND video_id = ?", userId, videoId).First(&favorite)

	// 检查是否找到匹配的点赞记录
	return result.Error == nil
}

// AddLike ，点赞操作
func (s *FavoriteService) AddLike(userId, videoId uint) error {
	// 获取数据库连接实例
	db := config.Database

	// 新建一条点赞记录
	favorite := models.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	result := db.Create(&favorite)
	if result.Error != nil {
		return result.Error
	}

	// 将该视频的获赞总数加一
	if err := models.IncrementVideoLikeCount(uint(videoId)); err != nil {
		return err
	}

	// 将本次点赞用户的点赞数加一
	if err := models.IncrementUserLikeCount(userId); err != nil {
		return err
	}

	// 根据该视频id来获取其作者id
	author_id, err := models.GetAuthorIDForVideo(uint(videoId))
	if err != nil {
		return err
	}

	// 将视频作者的获赞总数加一
	if err := models.IncrementAuthorTotalFavorited(author_id); err != nil {
		return err
	}

	return nil
}

// CancelLike ，取消点赞操作
func (s *FavoriteService) CancelLike(userId, videoId uint) error {
	// 获取数据库连接实例
	db := config.Database

	// 查找要删除的现有点赞关系
	var existingFavorite models.Favorite
	result := db.Where("user_id = ? AND video_id = ?", userId, videoId).First(&existingFavorite)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// 如果点赞关系不存在，则无错误地返回
			return nil
		}
		return result.Error
	}

	// 删除现有的点赞关系
	result = db.Delete(&existingFavorite)
	if result.Error != nil {
		return result.Error
	}

	// 成功取消点赞后，调用 DecrementUserLikeCount 函数，将用户的点赞数减一
	if err := models.DecrementUserLikeCount(userId); err != nil {
		return err
	}

	// 成功取消点赞后，调用 DecrementVideoLikeCount 函数，将视频的获赞数加一
	if err := models.DecrementVideoLikeCount(uint(videoId)); err != nil {
		return err
	}

	// 获取该取消赞的视频作者id
	author_id, err := models.GetAuthorIDForVideo(uint(videoId))
	if err != nil {
		return err
	}

	// 将视频作者的获赞总数减一
	if err := models.DecrementAuthorTotalFavorited(author_id); err != nil {
		return err
	}

	return nil // 操作成功，返回 nil 表示没有错误
}

// GetFavoriteList ，根据用户ID取出该用户点赞的所有视频ID
func (s *FavoriteService) GetFavoriteList(userId uint) ([]uint, error) {
	var favorites []models.Favorite

	// 查询该用户点赞的所有记录
	result := config.Database.Where("user_id = ?", userId).Find(&favorites)
	if result.Error != nil {
		return nil, result.Error
	}

	// 提取视频ID并放入列表
	var videoIDs []uint
	for _, favorite := range favorites {
		videoIDs = append(videoIDs, favorite.VideoId)
	}

	return videoIDs, nil
}
