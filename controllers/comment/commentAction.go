package comment

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/services"
)

var CommentService services.CommentService

// CommentAction 处理评论操作的请求
func CommentAction(c *gin.Context) {
	// 获取当前登录的用户id
	userId := c.Value("userID").(uint)
	// 评论所属的视频id
	video_id, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		// 视频 id 参数类型转换失败
		response.VideoIdConversionError(c)
		return
	}
	// 操作类型
	action_type := c.Query("action_type")
	switch action_type {
	case "1":
		// 发布评论
		content := c.Query("comment_text")
		commentResp, err := CommentService.CreateComment(uint(video_id), content, userId)
		if err != nil {
			// 评论保存失败
			response.CommentSaveFailed(c)
			return
		}
		// 评论添加成功
		response.CommentAddSuccess(c, commentResp)
		return
	case "2":
		// 删除评论
		comment_id_str := c.Query("comment_id")
		comment_id, err := strconv.ParseInt(comment_id_str, 10, 64)
		if err != nil {
			// 评论id参数类型转换失败
			response.CommentIdConversionError(c)
			return
		}
		err = CommentService.DeleteCommentById(userId, uint(video_id), uint(comment_id))
		if err != nil {
			// 评论删除失败
			response.CommentDelFailed(c)
			return
		}
		// 评论删除成功
		response.CommentDelSuccess(c)
		return
	default:
		// 非法操作
		response.UnsuccessfulAction(c)
	}
}
