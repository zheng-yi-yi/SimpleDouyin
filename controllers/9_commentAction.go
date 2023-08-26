package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

// CommentAction 处理评论操作的请求。
func CommentAction(c *gin.Context) {
	userId := c.GetUint("userID")
	if userId == 0 {
		Failed(c, "用户不存在")
		return
	}
	action_type := c.Query("action_type")
	video_id, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		Failed(c, err.Error())
		return
	}
	// 发布评论
	if action_type == "1" {
		content := c.Query("comment_text")
		if err := commentService.CreateComment(uint(video_id), content, userId); err != nil {
			Failed(c, err.Error())
			return
		}
		if err := models.IncrementCommentCount(config.Database, uint(video_id)); err != nil {
			Failed(c, err.Error())
			return
		}
		Success(c, "成功添加一条评论")
		return
	}
	// 删除评论
	comment_id, parseCommentIdErr := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	if parseCommentIdErr != nil {
		Failed(c, parseCommentIdErr.Error())
		return
	}
	if err := commentService.DeleteCommentById(userId, uint(video_id), uint(comment_id)); err != nil {
		Failed(c, err.Error())
		return
	}
	if err := models.DecreaseCommentCount(config.Database, uint(video_id)); err != nil {
		Failed(c, err.Error())
		return
	}
	Success(c, "成功删除一条评论")
}
