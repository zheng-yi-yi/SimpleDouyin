package comment

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

// 评论列表,已删除的评论和已注销用户的评论不会展示
func CommentList(c *gin.Context) {
	// 当前登录的用户
	userId := c.Value("userID").(uint)
	// 要查询的视频id
	videoId, err := strconv.ParseUint(c.Query("video_id"), 10, 64)
	if err != nil {
		// 视频id参数类型转换失败
		response.VideoIdConversionError(c)
		return
	}
	comments, err := CommentService.GetVideoComment(uint(videoId))
	if err != nil {
		// 评论列表获取失败
		response.GetCommentListFailed(c)
		return
	}
	commentList := make([]response.Comment, 0, len(comments))
	for _, comment := range comments {
		userInfo, getUserInfoErr := models.FetchData(comment.UserId)
		if getUserInfoErr != nil {
			continue
		}
		commentList = append(commentList, response.Comment{
			Id: int64(comment.ID),
			User: response.User{
				Id:             int64(userInfo.ID),
				Name:           userInfo.UserName,
				FollowCount:    int64(userInfo.FollowCount),
				FollowerCount:  int64(userInfo.FollowerCount),
				IsFollow:       models.IsFollow(userId, userInfo.ID),
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
	// 评论列表获取成功
	response.GetCommentListSuccess(c, commentList)
}
