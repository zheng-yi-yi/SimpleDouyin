package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
)

// 评论列表,已删除的评论和已注销用户的评论不会展示
func CommentList(c *gin.Context) {
	token, video_id := c.Query("token"), c.Query("video_id")
	userId := UsersLoginInfo[token].ID
	videoId, parseVideoId := strconv.ParseUint(video_id, 10, 64)
	if parseVideoId != nil {
		Failed(c, parseVideoId.Error())
		return
	}
	comments, err := commentService.GetVideoComment(uint(videoId))
	if err != nil {
		Failed(c, err.Error())
		return
	}
	commentList := make([]Comment, 0, len(comments))
	for _, comment := range comments {
		userInfo, getUserInfoErr := userService.GetUserInfoById(comment.UserId)
		if getUserInfoErr != nil {
			continue
		}
		commentList = append(commentList, Comment{
			Id: int64(comment.ID),
			User: User{
				Id:             int64(userInfo.ID),
				Name:           userInfo.UserName,
				FollowCount:    int64(userInfo.FollowCount),
				FollowerCount:  int64(userInfo.FollowerCount),
				IsFollow:       IsFollow(userId, userInfo.ID),
				Avatar:         userInfo.Avatar,
				Background:     userInfo.BackgroundImage,
				Signature:      userInfo.Signature,
				TotalFavorited: userInfo.TotalFavorited,
				WorkCount:      userInfo.WorkCount,
				FavoriteCount:  userInfo.FavoriteCount,
			},
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format(config.SHORT_DATE_FORMAT),
		})
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		CommentList: commentList,
	})
}
