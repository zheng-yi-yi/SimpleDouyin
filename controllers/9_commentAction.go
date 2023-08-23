package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
)

// 评论操作，不允许发送空评论,删除操作仅自身可以删除或者视频作者可以删除
func CommentAction(c *gin.Context) {
	token, action_type := c.Query("token"), c.Query("action_type")
	video_id, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	user := UsersLoginInfo[token]
	if action_type == "1" {
		content := c.Query("comment_text")
		comment, err := commentService.CreateComment(uint(video_id), content, uint(user.ID))
		if err != nil {
			fmt.Print(111)
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{
					StatusCode: 1,
					StatusMsg:  "不允许发送空评论!",
				},
				Comment: Comment{},
			})
		}
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			Comment: Comment{
				Id: int64(comment.ID),
				User: User{
					Id:     int64(user.ID),
					Name:   user.UserName,
					Avatar: config.AvatarURL,
				},
				Content:    comment.Content,
				CreateDate: comment.CreatedAt.Format("01-02"),
			},
		})
	} else {
		comment_id, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		err := commentService.DeleteCommentById(user.ID, uint(comment_id))
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{
					StatusCode: 1,
					StatusMsg:  "无权限删除该评论",
				},
			})
		}
	}
}
