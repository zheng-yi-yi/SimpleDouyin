package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Comment struct {
	Id         int64  `json:"id"`          // 评论id
	User       User   `json:"user"`        // 评论用户
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
}

// 评论操作响应
type CommentActionResponse struct {
	StatusCode int32   `json:"status_code"`       // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`        // 返回状态描述
	Comment    Comment `json:"comment,omitempty"` // 评论成功返回评论内容，不需要重新拉取整个列表
}

// CommentAddSuccess ， 评论添加成功
func CommentAddSuccess(c *gin.Context, comment Comment) {
	c.JSON(http.StatusOK, CommentActionResponse{
		StatusCode: 0,
		StatusMsg:  "评论添加成功",
		Comment:    comment,
	})
}

// CommentSaveFailed ， 评论保存失败
func CommentSaveFailed(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, CommentActionResponse{
		StatusCode: 1,
		StatusMsg:  "评论保存失败",
	})
}

// CommentDelSuccess ， 评论删除成功
func CommentDelSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, CommentActionResponse{
		StatusCode: 0,
		StatusMsg:  "评论删除成功",
	})
}

// CommentDelFailed ， 评论删除失败
func CommentDelFailed(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, CommentActionResponse{
		StatusCode: 1,
		StatusMsg:  "评论删除失败",
	})
}

// 评论列表响应
type CommentListResponse struct {
	StatusCode  int64     `json:"status_code"`            // 状态码，0-成功，其他值-失败
	StatusMsg   string    `json:"status_msg"`             // 返回状态描述
	CommentList []Comment `json:"comment_list,omitempty"` // 评论列表
}

// GetCommentListSuccess ， 评论列表获取成功
func GetCommentListSuccess(c *gin.Context, commentList []Comment) {
	c.JSON(http.StatusOK, CommentListResponse{
		StatusCode:  0,
		StatusMsg:   "评论列表获取成功",
		CommentList: commentList,
	})
}

// GetCommentListFailed ， 评论列表获取失败
func GetCommentListFailed(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, CommentListResponse{
		StatusCode: 1,
		StatusMsg:  "评论列表获取失败",
	})
}
