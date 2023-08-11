package controllers

import (
	"github.com/gin-gonic/gin"
)

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

// 评论操作
func CommentAction(c *gin.Context) {

}

// 评论列表
func CommentList(c *gin.Context) {

}
