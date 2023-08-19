package services

import (
	"github.com/zheng-yi-yi/simpledouyin/models"
	"gorm.io/gorm"
)

type SocializeService struct {
	db *gorm.DB
}

// GetFollowersByUserID unGetFollowerByUserID 通过用户ID获取粉丝用户列表
func (s *SocializeService) GetFollowersByUserID(userID uint) ([]models.User, error) {
	// 使用UserID从数据库中获取粉丝用户列表
	var followers []models.Relation

	err := s.db.Model(&models.Relation{}).Where("to_user_id = ?", userID).Find(&followers).Error
	if err != nil {
		return nil, err
	}

	followerIDs := make([]uint, len(followers))
	for i, follower := range followers {
		followerIDs[i] = follower.FromUserId
	}

	var followerUsers []models.User
	err = s.db.Where("id IN (?)", followerIDs).Find(&followerUsers).Error
	if err != nil {
		return nil, err
	}

	return followerUsers, nil
}

// GetFriendsByUserID 通过用户ID获取好友用户列表
func (s *SocializeService) GetFriendsByUserID(userID uint) ([]models.User, error) {
	// 使用UserID从数据库中获取好友用户列表
	var relations []models.Relation

	// Find relations where both from_user_id and to_user_id are involved
	err := s.db.Where("from_user_id = ? OR to_user_id = ?", userID, userID).Find(&relations).Error
	if err != nil {
		return nil, err
	}

	friendIDs := make([]uint, 0)
	for _, relation := range relations {
		// Ensure the other user's ID is added to the friendIDs list
		if relation.FromUserId == userID {
			friendIDs = append(friendIDs, relation.ToUserId)
		} else {
			friendIDs = append(friendIDs, relation.FromUserId)
		}
	}

	var friendUsers []models.User
	err = s.db.Where("id IN (?)", friendIDs).Find(&friendUsers).Error
	if err != nil {
		return nil, err
	}

	return friendUsers, nil
}
