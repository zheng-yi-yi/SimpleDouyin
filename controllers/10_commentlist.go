package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
)

// 评论列表,已删除的评论和已注销用户的评论不会展示
func CommentList(c *gin.Context) {
	token, video_id := c.Query("token"), c.Query("video_id")
	//fmt.Println(video_id)
	Comments, err := commentService.GetVideoComment(video_id)
	if err != nil { //如果视频id非法，返回
		fmt.Println(1000)
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "找不到该视频!",
			},
			CommentList: []Comment{},
		})
	}
	CommentsList := make([]Comment, 0, len(Comments))
	for _, comment := range Comments {
		userInfo, err := userService.GetUserInfoById(comment.UserId) //如果评论用户信息不存在或已注销，不计该评论
		if err != nil {
			continue
		}
		CommentsList = append(CommentsList, Comment{
			Id: int64(comment.ID),
			User: User{
				Id:         int64(userInfo.ID),
				Name:       userInfo.UserName,
				IsFollow:   IsFollow(UsersLoginInfo[token].ID, userInfo.ID),
				Background: config.BackgroundURL,
				Avatar:     config.AvatarURL,
			},
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("01-02"),
		})
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		CommentList: CommentsList,
	})

}
