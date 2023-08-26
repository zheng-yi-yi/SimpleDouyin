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
