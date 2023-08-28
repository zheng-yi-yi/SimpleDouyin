package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

// UnsuccessfulAction ，操作失败
func UnsuccessfulAction(c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		StatusCode: 1,
		StatusMsg:  "，操作失败",
	})
}

// VideoIdConversionError ， 视频id参数类型转换失败
func VideoIdConversionError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		StatusCode: 1,
		StatusMsg:  "视频id参数类型转换失败",
	})
}

// CommentIdConversionError ，评论id参数类型转换失败
func CommentIdConversionError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		StatusCode: 1,
		StatusMsg:  "评论id参数类型转换失败",
	})
}

// UserIdConversionError ， 用户id参数类型转换失败
func UserIdConversionError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		StatusCode: 1,
		StatusMsg:  "用户id参数类型转换失败",
	})
}

// ToUserIdConversionError ， 对方用户id参数类型转换失败
func ToUserIdConversionError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		StatusCode: 1,
		StatusMsg:  "对方用户id参数类型转换失败",
	})
}
