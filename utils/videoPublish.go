package utils

import (
	"fmt"
	"log"

	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

// GetVideoName 根据userId_用户发布的视频数量+1
func GetVideoName(userId uint) string {
	VideoCount, err := models.GetVideoCount(config.DB, userId)
	if err != nil {
		log.Printf("无法获取用户的视频个数")
		return ""
	}
	filename := fmt.Sprintf("%d_%d", userId, VideoCount+1)
	return filename
}
