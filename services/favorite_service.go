package services

import (
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

type FavoriteService struct {
}

// 判断用户是否点赞了视频
func (s *FavoriteService) IsFavorite(userId, videoId uint) bool {
	var favorite models.Favorite
	err := config.DB.Where("user_id = ? AND video_id = ? AND status = 1", userId, videoId).First(&favorite).Error
	return err == nil
}
