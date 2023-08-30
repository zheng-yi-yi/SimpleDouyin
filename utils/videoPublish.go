package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/zheng-yi-yi/simpledouyin/models"
)

// GetVideoDst ， 根据视频文件和用户ID获取视频目标路径
func GetVideoDst(file *multipart.FileHeader, userId uint) string {
	// 获取视频名称
	videoName := GetVideoName(userId)

	// 获取文件扩展名
	fileExt := filepath.Ext(file.Filename)

	// 根据视频名称和文件扩展名生成新的文件名
	filename := videoName + fileExt

	// 将文件名与视频文件目录路径合并得到完整的视频目标路径
	videoDst := filepath.Join("./public/videos/", filename)

	// 返回视频目标路径
	return videoDst
}

// GetVideoPath ， 根据视频文件和用户ID获取视频文件路径
func GetVideoPath(file *multipart.FileHeader, userId uint) string {
	// 获取当前工作目录
	pwd, getPwdErr := os.Getwd()
	if getPwdErr != nil {
		fmt.Println(getPwdErr.Error())
		return ""
	}

	// 获取父目录
	parentDir := filepath.Dir(pwd)

	// 获取视频目标路径
	videoDst := GetVideoDst(file, userId)

	// 将父目录与视频目标路径合并得到完整的视频文件路径
	videoPath := filepath.Join(parentDir, videoDst)

	// 返回视频文件路径
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

// GetCoverPath 根据文件信息和用户ID生成封面图片文件的完整本地路径。
func GetCoverPath(file *multipart.FileHeader, userId uint) string {
	pwd, getPwdErr := os.Getwd()
	if getPwdErr != nil {
		fmt.Println(getPwdErr.Error())
		return ""
	}
	parentDir := filepath.Dir(pwd)
	coverDst := GetCoverDst(file, userId)
	coverPath := filepath.Join(parentDir, coverDst)
	return coverPath
}

// GetVideoName ， 根据用户ID获取视频名称
func GetVideoName(userId uint) string {
	// 调用 models 包中的 GetVideoCount 函数，获取用户的视频数量和可能的错误
	VideoCount, err := models.GetVideoCount(userId)

	// 如果获取视频数量时发生错误
	if err != nil {
		VideoCount = 0 // 将视频数量设置为0
	}

	// 使用格式化字符串生成视频文件名，格式为 "用户ID_视频数量+1"
	filename := fmt.Sprintf("%d_%d", userId, VideoCount+1)

	// 返回生成的视频文件名
	return filename
}

// GetCoverName ， 根据用户ID生成封面图片名称
func GetCoverName(userId uint) string {
	// 调用 models 包中的 GetVideoCount 函数，获取用户的视频数量和可能的错误
	VideoCount, err := models.GetVideoCount(userId)

	// 如果获取视频数量时发生错误
	if err != nil {
		VideoCount = 0
	}

	// 使用格式化字符串生成封面图片文件名，格式为 "用户ID_视频数量+1.png"
	coverName := fmt.Sprintf("%d_%d.png", userId, VideoCount+1)

	// 返回生成的封面图片文件名
	return coverName
}
