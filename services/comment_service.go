package services

import (
	"errors"
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
	return config.Database.Create(&comment).Error
}

// 获取所有未删除的评论
func (s *CommentService) GetVideoComment(video_id uint) ([]models.Comment, error) {
	var commentList []models.Comment
	if err := config.Database.Where("video_id=?", video_id).Find(&commentList).Error; err != nil {
		return nil, err
	}
	return commentList, nil
}

// 根据相应的评论获取ID
func (s *CommentService) GetCommentById(comment_id int64) models.Comment {
	var comment models.Comment
	config.Database.Where("id = ?", uint(comment_id)).First(&comment)
	return comment
}

// 根据用户ID、视频ID和评论ID，定位并删除对应评论
func (s *CommentService) DeleteCommentById(userId, videoId, commentId uint) error {
	return config.Database.Where("user_id=? and video_id=? and id=?", userId, videoId, commentId).Delete(&models.Comment{}).Error
}
