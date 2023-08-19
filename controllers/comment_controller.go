package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
)

type Comment struct {
	Id         int64  `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}
type CommentActionResponse struct {
	Response
	Comment `json:"comment,omitempty"`
}

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
