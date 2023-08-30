package services

import (
	"log"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
	"github.com/zheng-yi-yi/simpledouyin/utils"
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

// 获取指定用户发布的视频列表
func (videoService *VideoService) UserPublishList(userId uint) []*models.Video {
	var videoList []*models.Video
	config.Database.
		Where("user_id = ?", userId).
		Preload("User").
		Find(&videoList)
	return videoList
}

// GetVideoInfoByIds ，根据点赞过的视频ID ，取出所有对应的视频信息
func (videoService *VideoService) GetVideoInfoByIds(videoIds []uint) []*models.Video {
	var videoList []*models.Video
	config.Database.Where("id IN ?", videoIds).Preload("User").Find(&videoList)
	return videoList
}

// 用户投稿视频
func (videoService *VideoService) SaveVideoFile(c *gin.Context, file *multipart.FileHeader, userId uint, title string) error {
	// 获取视频文件目标路径
	videoDst := utils.GetVideoDst(file, userId)
	// 保存上传的视频文件到目标路径
	if err := c.SaveUploadedFile(file, videoDst); err != nil {
		// 视频文件保存失败
		log.Printf("视频文件保存失败: %s", err)
		return err
	}

	// 获取视频文件和封面图片文件的本地路径
	videoPath := utils.GetVideoPath(file, userId)
	coverPath := utils.GetCoverPath(file, userId)

	// 使用 Ffmpeg 函数生成封面图片
	utils.Ffmpeg(videoPath, coverPath)

	// 生成 playUrl 与 coverUrl
	fileExt := filepath.Ext(file.Filename)
	playUrl := "videos/" + utils.GetVideoName(userId) + fileExt
	coverUrl := "images/" + utils.GetCoverName(userId)
	// 创建视频记录
	video := models.Video{
		UserId:      userId,
		PlayUrl:     playUrl,
		CoverUrl:    coverUrl,
		Description: title,
	}
	err := config.Database.Create(&video).Error
	if err != nil {
		log.Printf("视频数据保存失败: %s", err)
		return err
	}
	// 成功创建视频后，调用 IncrementWorkCount 函数
	if err := models.IncrementWorkCount(userId); err != nil {
		log.Printf("用户作品数添加失败: %s", err)
		return err
	}
	return nil
}
