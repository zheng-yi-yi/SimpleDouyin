package services

import (
	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
)

type RelationService struct {
}

// 判断是否关注
func (s *RelationService) IsFollow(fromUserId, toUserId uint) bool {
	var relation models.Relation
	err := config.DB.
		Where("(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ? AND is_mutual = 1)",
			fromUserId, toUserId, toUserId, fromUserId).
		First(&relation).Error
	return err == nil
}
