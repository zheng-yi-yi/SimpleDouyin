package services

import (
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

type VideoService struct {
}

// 获取视频Feed
func (videoService *VideoService) Feed(startTime string) *[]models.Video {
	var videoList *[]models.Video
	config.Database.
		Where("created_at <= ?", startTime).
		Preload("User").
		Order("created_at DESC").
		Limit(config.VIDEO_STREAM_BATCH_SIZE).
		Find(&videoList)
	return videoList
}

// Create ： 在数据库中创建视频记录
func (videoService *VideoService) Create(playUrl, coverUrl, desc string, userId uint) (models.Video, error) {
	video := models.Video{
		UserId:      userId,
		PlayUrl:     playUrl,
		CoverUrl:    coverUrl,
		Description: desc,
	}
	return video, config.Database.Create(&video).Error
}

// 获取指定用户发布的视频列表
func (videoService *VideoService) UserPublishList(userId uint) []*models.Video {
	var videoList []*models.Video
	config.Database.
		Where("user_id = ?", userId).
		Preload("User").
		Find(&videoList)
	return videoList
}
