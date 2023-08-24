package utils

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

// GetVideoDst 根据文件信息和用户ID生成视频文件的目标路径。
func GetVideoDst(file *multipart.FileHeader, userId uint) string {
	videoName := GetVideoName(userId)
	if videoName == "" {
		return ""
	}
	fileExt := filepath.Ext(file.Filename)
	filename := videoName + fileExt
	videoDst := filepath.Join("./public/videos/", filename)
	return videoDst
}

// GetVideoPath 根据文件信息和用户ID生成视频文件的完整本地路径。
func GetVideoPath(file *multipart.FileHeader, userId uint) string {
	pwd, getPwdErr := os.Getwd()
	if getPwdErr != nil {
		fmt.Println(getPwdErr.Error())
		return ""
	}
	parentDir := filepath.Dir(pwd)
	videoDst := GetVideoDst(file, userId)
	videoPath := filepath.Join(parentDir, videoDst)
	return videoPath
}

// GetCoverDst 根据文件信息和用户ID生成封面图片文件的目标路径。
func GetCoverDst(file *multipart.FileHeader, userId uint) string {
	coverName := GetCoverName(userId)
	if coverName == "" {
		return ""
	}
	coverDst := filepath.Join("./public/images/", coverName)
	return coverDst
}

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

// GetCoverName 根据用户ID生成封面图片的文件名
func GetCoverName(userId uint) string {
	VideoCount, err := models.GetVideoCount(config.DB, userId)
	if err != nil {
		log.Printf("无法获取用户的视频个数")
		return ""
	}
	coverName := fmt.Sprintf("%d_%d.png", userId, VideoCount+1)
	return coverName
}
