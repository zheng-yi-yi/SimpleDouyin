package services

import (
	"errors"
	"strconv"

	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

type CommentService struct {
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

// 往DB中插入相应的评论。另外to_user_id似乎没用到，所以没设置
func (s *CommentService) CreateComment(video_id uint, content string, user_id uint) (models.Comment, error) {
	if len(content) == 0 { //不允许空评论
		return models.Comment{}, errors.New("插入空评论！")
	}
	comment := models.Comment{
		Content: content,
		VideoId: video_id,
		UserId:  user_id,
	}
	config.DB.Create(&comment)
	return comment, nil
}

// 根据相应的评论获取ID
func (s *CommentService) GetCommentById(comment_id int64) models.Comment {
	var comment models.Comment
	config.DB.Where("id = ?", uint(comment_id)).First(&comment)
	return comment
}

func (s *CommentService) DeleteCommentById(user_id uint, comment_id uint) error {
	var comment models.Comment
	comment = s.GetCommentById(int64(comment_id))
	if comment.UserId == user_id {
		config.DB.Where("id = ?", comment_id).Delete(&comment)
		return nil
	}
	return errors.New("无权限删除该评论!")
}
