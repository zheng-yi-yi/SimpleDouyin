package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 关注操作响应
type RelationActionResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

// RelationActionSucceeded , 关注成功
func RelationActionSucceeded(c *gin.Context) {
	c.JSON(http.StatusOK, RelationActionResponse{
		StatusCode: 0,
		StatusMsg:  "关注成功",
	})
}

// RelationActionError ， 关注失败
func RelationActionError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, RelationActionResponse{
		StatusCode: 1,
		StatusMsg:  "关注失败",
	})
}

// CancelRelationSucceeded , 取关成功
func CancelRelationSucceeded(c *gin.Context) {
	c.JSON(http.StatusOK, RelationActionResponse{
		StatusCode: 0,
		StatusMsg:  "取关成功",
	})
}

// CancelRelationError ， 取关失败
func CancelRelationError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, RelationActionResponse{
		StatusCode: 1,
		StatusMsg:  "取关失败",
	})
}

// 关注列表响应
type FollowListResponse struct {
	StatusCode string `json:"status_code"`         // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`          // 返回状态描述
	UserList   []User `json:"user_list,omitempty"` // 用户信息列表
}

// GetFollowListSucceeded , 关注列表获取成功
func GetFollowListSucceeded(c *gin.Context, userList []User) {
	c.JSON(http.StatusOK, FollowListResponse{
		StatusCode: "0",
		StatusMsg:  "关注列表获取成功",
		UserList:   userList,
	})
}

// GetFollowListError ， 关注列表获取失败
func GetFollowListError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, FollowListResponse{
		StatusCode: "1",
		StatusMsg:  "关注列表获取失败",
	})
}

// 粉丝列表响应
type FollowerListResponse struct {
	StatusCode string `json:"status_code"`         // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`          // 返回状态描述
	UserList   []User `json:"user_list,omitempty"` // 用户列表
}

// GetFollowerListSucceeded , 粉丝列表获取成功
func GetFollowerListSucceeded(c *gin.Context, userList []User) {
	c.JSON(http.StatusOK, FollowerListResponse{
		StatusCode: "0",
		StatusMsg:  "粉丝列表获取成功",
		UserList:   userList,
	})
}

// GetFollowerListError ， 粉丝列表获取失败
func GetFollowerListError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, FollowerListResponse{
		StatusCode: "1",
		StatusMsg:  "粉丝列表获取失败",
	})
}

// 好友列表响应
type FriendListResponse struct {
	StatusCode string `json:"status_code"`         // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`          // 返回状态描述
	UserList   []User `json:"user_list,omitempty"` // 用户列表
}

// GetFriendListSucceeded , 好友列表获取成功
func GetFriendListSucceeded(c *gin.Context, userList []User) {
	c.JSON(http.StatusOK, FriendListResponse{
		StatusCode: "0",
		StatusMsg:  "粉丝列表获取成功",
		UserList:   userList,
	})
}

// GetFriendListError ， 好友列表获取失败
func GetFriendListError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, FriendListResponse{
		StatusCode: "1",
		StatusMsg:  "粉丝列表获取失败",
	})
}
