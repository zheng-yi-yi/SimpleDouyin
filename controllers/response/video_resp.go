package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Video struct {
	Id            int64  `json:"id,omitempty"`   // 视频唯一标识
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverUrl      string `json:"cover_url"`      // 视频封面地址
	PlayUrl       string `json:"play_url"`       // 视频播放地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	Title         string `json:"title"`          // 视频标题
}

type VideoList []Video

// 视频流响应
type FeedResponse struct {
	StatusCode int32   `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`           // 返回状态描述
	VideoList  []Video `json:"video_list,omitempty"` // 视频猎豹
	NextTime   int64   `json:"next_time,omitempty"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

// 获取视频流成功
func GetFeedSuccess(c *gin.Context, nextTime int64, videoList []Video) {
	c.JSON(http.StatusOK, FeedResponse{
		StatusCode: 0,
		StatusMsg:  "成功获取视频流",
		VideoList:  videoList,
		NextTime:   nextTime,
	})
}

// 视频投稿响应
type PostVideoResponse struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

// VideoFileAccessError , 视频文件获取失败
func VideoFileAccessError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, PostVideoResponse{
		StatusCode: 1,
		StatusMsg:  "视频文件获取失败",
	})
}

// VideoFileSaveFailure , 视频文件保存失败
func VideoFileSaveFailure(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, PostVideoResponse{
		StatusCode: 1,
		StatusMsg:  "视频文件保存失败",
	})
}

// PostVideoSuccessful , 视频发布成功
func PostVideoSuccessful(c *gin.Context) {
	c.JSON(http.StatusOK, PostVideoResponse{
		StatusCode: 0,
		StatusMsg:  "视频发布成功",
	})
}

// 发布列表响应
type VideoListResponse struct {
	StatusCode int32   `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`           // 返回状态描述
	VideoList  []Video `json:"video_list,omitempty"` // 用户发布的视频列表
}

// GetPublishListSuccess , 视频列表获取成功
func GetPublishListSuccess(c *gin.Context, videoList []Video) {
	c.JSON(http.StatusOK, VideoListResponse{
		StatusCode: 0,
		StatusMsg:  "视频列表获取成功",
		VideoList:  videoList,
	})
}
