package services

import (
	"time"

	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/controllers/response"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

type CommentService struct {
}

// CreateComment: 创建新评论
func (s *CommentService) CreateComment(video_id uint, content string, user_id uint) (response.Comment, error) {
	// 新建一条新评论记录
	comment := models.Comment{
		VideoId:   video_id,
		Content:   content,
		UserId:    user_id,
		CreatedAt: time.Now(),
	}
	result := config.Database.Create(&comment)
	if result.Error != nil {
		return response.Comment{}, result.Error
	}

	// 如果添加评论成功，则增加该视频的评论总数
	if err := models.IncrementCommentCount(uint(video_id)); err != nil {
		return response.Comment{}, err
	}

	// 评论成功则返回评论内容（响应）
	userInfo, _ := models.FetchData(user_id)
	NewComment := response.Comment{
		Id: int64(comment.ID),
		User: response.User{
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
	// 删除评论
	result := config.Database.Where("user_id=? and video_id=? and id=?", userId, videoId, commentId).Delete(&models.Comment{})
	if result.Error != nil {
		return result.Error
	}
	// 成功删除评论后，就减少该视频的评论总数
	if err := models.DecreaseCommentCount(uint(videoId)); err != nil {
		return err
	}
	return nil
}
