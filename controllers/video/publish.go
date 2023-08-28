package video

import (
	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
)

// Publish , 处理视频投稿发布请求。
func Publish(c *gin.Context) {
	// 获取上传的视频文件
	file, err := c.FormFile("data")
	if err != nil {
		response.VideoFileAccessError(c) // 视频获取失败
		return
	}
	// 当前用户id
	userId := c.Value("userID").(uint)
	// 获取视频标题
	title := c.PostForm("title")
	// 保存视频
	err = VideoService.SaveVideoFile(c, file, userId, title)
	if err != nil {
		response.VideoFileSaveFailure(c) // 视频保存失败
		return
	}
	// 视频发布成功
	response.PostVideoSuccessful(c)
}
