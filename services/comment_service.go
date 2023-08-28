package services

import (
	"errors"
	"time"

	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

type CommentService struct {
}

// CreateComment: 创建新评论
func (s *CommentService) CreateComment(video_id uint, content string, user_id uint) (response.Comment, error) {
	if len(content) == 0 {
		return response.Comment{}, errors.New("不能插入空评论！")
	}
	comment := models.Comment{
		VideoId:   video_id,
		Content:   content,
		UserId:    user_id,
		CreatedAt: time.Now(),
		// Cancel:    0,
	}
	result := config.Database.Create(&comment)
	if err := models.IncrementCommentCount(uint(video_id)); err != nil {
		return response.Comment{}, err
	}
	userInfo, _ := models.FetchData(user_id)
	NewComment := response.Comment{
		Id: int64(comment.ID), // 评论id
		User: response.User{ // 评论用户
			Id:     int64(userInfo.ID),
			Name:   userInfo.UserName,
			Avatar: userInfo.Avatar,
		},
		Content:    content,                                     // 评论内容
		CreateDate: time.Now().Format(config.SHORT_DATE_FORMAT), // 评论发布日期，格式 mm-dd
	}
	return NewComment, result.Error
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
	result := config.Database.Where("user_id=? and video_id=? and id=?", userId, videoId, commentId).Delete(&models.Comment{})
	if err := models.DecreaseCommentCount(uint(videoId)); err != nil {
		return err
	}
	return result.Error
}
