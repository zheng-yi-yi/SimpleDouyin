package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/models"
	"github.com/zheng-yi-yi/simpledouyin/services"
)

// ================= Service =================
var (
	userService     services.UserService
	videoService    services.VideoService
	favoriteService services.FavoriteService
	relationService services.RelationService
	messageService  services.MessageService
	commentService  services.CommentService
)

// ================= Response =================
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

// ================= Video =================
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

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list"`
	NextTime  int64   `json:"next_time"`
}

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// ================= User =================
type User struct {
	Id             int64  `json:"id,omitempty"`               // 用户id
	Name           string `json:"name,omitempty"`             // 用户名称
	FollowCount    int64  `json:"follow_count,omitempty"`     // 关注总数
	FollowerCount  int64  `json:"follower_count,omitempty"`   // 粉丝总数
	IsFollow       bool   `json:"is_follow,omitempty"`        // true-已关注，false-未关注
	Avatar         string `json:"avatar,omitempty"`           // 用户头像
	Background     string `json:"background_image,omitempty"` // 用户个人页顶部大图
	Signature      string `json:"signature,omitempty"`        // 个人简介
	FavoriteCount  int64  `json:"favorite_count"`             // 喜欢数
	TotalFavorited string `json:"total_favorited"`            // 获赞数量
	WorkCount      int64  `json:"work_count"`                 // 作品数
}

type UserList []User

var UsersLoginInfo = map[string]models.User{}

type UserLoginResponse struct {
	Response
	UserId int    `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User `json:"user"`
}

// ================= Comment =================
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

// ================= relation =================
type relationUser struct {
	ID              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Avatar          string `json:"avatar,omitempty"`           //用户头像
	Signature       string `json:"signature,omitempty"`        //个人简介
	FollowCount     int64  `json:"follow_count,omitempty"`     //关注数量
	FollowerCount   int64  `json:"follower_count,omitempty"`   //粉丝数量
	IsFollow        bool   `json:"is_follow,omitempty"`        //是否关注
	BackgroundImage string `json:"background_image,omitempty"` //用户个人页顶部大图
	TotalFavorited  string `json:"total_favorited,omitempty"`  //获赞数量
	WorkCount       int64  `json:"work_count,omitempty"`       //作品数量
	FavoriteCount   int64  `json:"favorite_count,omitempty"`   //点赞数量
}

type UserListResponse struct {
	Response
	UserList []relationUser `json:"user_list"`
}

// ================= ChatResponse =================
type ChatResponse struct {
	Response
	MessageList []models.Message `json:"message_list"`
}
