package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/services"
)

// service 包调用
var (
	userService     services.UserService
	videoService    services.VideoService
	favoriteService services.FavoriteService
	relationService services.RelationService
	messageService  services.MessageService
)

// ================= 响应 =================
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func Success(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: msg})
}

func Failed(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: msg})
}

// ================= 视频 =================

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type VideoList []Video

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list"`
	NextTime  int64   `json:"next_time"`
}

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}
