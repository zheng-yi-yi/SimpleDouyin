package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zheng-yi-yi/simpledouyin/config"
)

type relationUser struct {
	ID              int64  `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Avatar          string `json:"avatar,omitempty"`           //用户头像
	Signature       string `json:"signature,omitempty"`        //个人简介
	FollowCount     int64  `json:"follow_count,omitempty"`     //关注数量
	FollowerCount   int64  `json:"follower_count,omitempty"`   //粉丝数量
	IsFollow        bool   `json:"is_follow,omitempty"`        //是否关注
	BackgroundImage string `json:"background_image,omitempty"` //用户个人页顶部大图
	TotalFavorited  int64  `json:"total_favorited,omitempty"`  //获赞数量
	WorkCount       int64  `json:"work_count,omitempty"`       //作品数量
	FavoriteCount   int64  `json:"favorite_count,omitempty"`   //点赞数量
}

type UserListResponse struct {
	Response
	UserList []relationUser `json:"user_list"`
}

// 关注操作
func RelationAction(c *gin.Context) {

	toUserIdStr := c.Query("to_user_id")
	actionType := c.Query("action_type")
	//userId := c.GetUint("UserID")
	token := c.Query("token")
	userId := UsersLoginInfo[token].ID

	if userId == 0 {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "不存在该用户",
			},
			UserList: []relationUser{},
		})
		return
	}

	//获取存储到上下文的用户id
	formUserId := userId

	//获取请求参数中的被关注用户id

	var toUserId uint64

	_toUserId, err := strconv.ParseUint(toUserIdStr, 10, 64)
	if err != nil {
		Failed(c, err.Error())
		return
	} else {
		toUserId = _toUserId
	}

	switch actionType {
	case "1":
		//关注操作
		err := relationService.FollowUser(uint(formUserId), uint(toUserId))
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 3, StatusMsg: "关注失败"})
		}

	case "2":
		//取消关注操作
		err := relationService.CancelFollowUser(uint(formUserId), uint(toUserId))
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 3, StatusMsg: err.Error()})

		}
	default:
		c.JSON(http.StatusBadRequest, Response{StatusCode: 3, StatusMsg: "无效操作"})
	}
}

// 关注列表
func FollowList(c *gin.Context) {
	userIdStr := c.Query("user_id")

	var formUserId uint64

	_formUserId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		Failed(c, err.Error())
		return
	} else {
		formUserId = _formUserId
	}

	//找用户关注的所有用户
	users, err := relationService.GetFllowList(uint(formUserId))
	if err != nil {
		Failed(c, err.Error())
		return
	}

	var relationUsers []relationUser
	//返回的用户信息到时候再完善一下TotalFavorited、WorkCount、FavoriteCount
	for _, user := range users {
		isFollow := relationService.IsFollow(uint(formUserId), user.ID)
		relationUser := relationUser{
			ID:              int64(user.ID),
			Name:            user.UserName,
			Avatar:          config.AvatarURL,
			Signature:       config.SignatureStr,
			FollowCount:     int64(user.FollowCount),
			FollowerCount:   int64(user.FollowerCount),
			IsFollow:        isFollow,
			BackgroundImage: config.BackgroundURL,
			TotalFavorited:  1,
			WorkCount:       1,
			FavoriteCount:   1,
		}
		relationUsers = append(relationUsers, relationUser)
	}

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{StatusCode: 0},
		UserList: relationUsers,
	})
}
