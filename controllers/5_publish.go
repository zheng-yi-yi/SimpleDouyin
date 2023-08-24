package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/utils"
)

// Publish 处理视频投稿发布请求。
func Publish(c *gin.Context) {
	// 获取用户ID
	userId := c.GetUint("userID")
	if userId == 0 {
		Failed(c, "不存在该用户...")
		return
	}
	// 获取视频标题
	title := c.PostForm("title")
	// 获取上传的视频文件
	file, getFileErr := c.FormFile("data")
	if getFileErr != nil {
		fmt.Println(getFileErr.Error())
		Failed(c, "获取视频文件失败...")
		return
	}
	// 获取视频文件目标路径
	videoDst := utils.GetVideoDst(file, userId)
	if videoDst == "" {
		Failed(c, "获取视频保存路径失败...")
		return
	}
	// 保存上传的视频文件到目标路径
	if err := c.SaveUploadedFile(file, videoDst); err != nil {
		fmt.Println(err.Error())
		Failed(c, "视频文件保存失败...")
		return
	}
	// 获取视频文件和封面图片文件的本地路径
	videoPath := utils.GetVideoPath(file, userId)
	coverPath := utils.GetCoverPath(file, userId)
	// 使用 Ffmpeg 函数生成封面图片
	if err := utils.Ffmpeg(videoPath, coverPath); err != nil {
		fmt.Println(err.Error())
		Failed(c, "视频封面生成失败...")
		return
	}
	// 获取封面图片目标路径
	coverDst := utils.GetCoverDst(file, userId)
	// 创建视频记录
	if _, createVideoErr := videoService.Create(videoDst, coverDst, title, userId); createVideoErr != nil {
		fmt.Println(createVideoErr.Error())
		Failed(c, "数据保存失败")
		return
	}
	// 返回一个成功的响应
	Success(c, "视频发布成功")
}
