package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 点赞操作响应
type FavoriteResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

// LikeActionSucceeded , 点赞成功
func LikeActionSucceeded(c *gin.Context) {
	c.JSON(http.StatusOK, FavoriteResponse{
		StatusCode: 0,
		StatusMsg:  "点赞成功",
	})
}

// LikeActionError ， 点赞失败
func LikeActionError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, FavoriteResponse{
		StatusCode: 1,
		StatusMsg:  "点赞失败",
	})
}

// UnlikeActionSucceeded , 取消赞成功
func UnlikeActionSucceeded(c *gin.Context) {
	c.JSON(http.StatusOK, FavoriteResponse{
		StatusCode: 0,
		StatusMsg:  "取消赞成功",
	})
}

// CancelLikeError , 取消赞失败
func CancelLikeError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, FavoriteResponse{
		StatusCode: 1,
		StatusMsg:  "取消赞失败",
	})
}

// 喜欢列表响应
type FavoriteListResponse struct {
	StatusCode string  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`           // 返回状态描述
	VideoList  []Video `json:"video_list,omitempty"` // 用户点赞视频列表
}

// GetFavoriteListSucceeded , 视频列表获取成功
func GetFavoriteListSucceeded(c *gin.Context, videoList []Video) {
	c.JSON(http.StatusOK, FavoriteListResponse{
		StatusCode: "0",
		StatusMsg:  "视频列表获取成功",
		VideoList:  videoList,
	})
}
