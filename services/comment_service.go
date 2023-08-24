package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

type CommentService struct {
}

// CreateComment: 创建新评论
func (s *CommentService) CreateComment(video_id uint, content string, user_id uint) error {
	if len(content) == 0 {
		return errors.New("不能插入空评论！")
	}
	comment := models.Comment{
		VideoId:   video_id,
		Content:   content,
		UserId:    user_id,
		CreatedAt: time.Now(),
		Cancel:    0,
	}
	return config.DB.Create(&comment).Error
}

// 获取所有未删除的评论
func (s *CommentService) GetVideoComment(video_id string) ([]models.Comment, error) {

	id, err := strconv.ParseInt(video_id, 10, 64)
	var Comments *[]models.Comment
	if err != nil {
		return *Comments, errors.New("视频不存在")
	}
	config.DB.Where("video_id = ? AND deleted_at is NULL", uint(id)).Find(&Comments)
	return *Comments, nil
}

// 根据相应的评论获取ID
func (s *CommentService) GetCommentById(comment_id int64) models.Comment {
	var comment models.Comment
	config.DB.Where("id = ?", uint(comment_id)).First(&comment)
	return comment
}

// 根据用户ID、视频ID和评论ID，定位并删除对应评论
func (s *CommentService) DeleteCommentById(userId, videoId, commentId uint) error {
	return config.DB.Where("user_id=? and video_id=? and id=?", userId, videoId, commentId).Delete(&models.Comment{}).Error
}
