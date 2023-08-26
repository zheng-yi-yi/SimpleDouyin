package services

import (
	"errors"

	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
	"gorm.io/gorm"
)

type RelationService struct {
}

// 关注用户
func (s *RelationService) FollowUser(fromUserId uint, toUserId uint) error {
	// 获取数据库连接的引用
	db := config.Database
	// 检查 fromUserId 和 toUserId 是否是有效的用户
	var fromUser, toUser models.User
	if err := db.First(&fromUser, fromUserId).Error; err != nil {
		return errors.New("未找到 fromUserId")
	}
	if err := db.First(&toUser, toUserId).Error; err != nil {
		return errors.New("未找到 toUserId")
	}
	// 查找已存在的关系，包括可能的取消关注情况
	var existingRelation models.Relation
	if err := db.Where("from_user_id = ? AND to_user_id = ?", fromUserId, toUserId).First(&existingRelation).Error; err == nil {
		// 如果已存在关系，检查是否是取消状态，如果是则重新激活
		if existingRelation.Cancel == 1 {
			// 更新为重新关注状态
			if err := db.Model(&existingRelation).Update("cancel", 0).Error; err != nil {
				return err
			}
			return nil
		}
		return errors.New("关系已经存在")
	}
	// 创建一个新的关注关系
	newRelation := models.Relation{
		FromUserId: fromUserId,
		ToUserId:   toUserId,
		Cancel:     0, // 默认：未取消
	}
	// 将新的关系插入到数据库中
	if err := db.Create(&newRelation).Error; err != nil {
		return err
	}
	// 更新关注用户的 FollowCount
	if err := config.Database.Model(&models.User{}).Where("id = ?", fromUserId).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
		return err
	}
	// 更新被关注用户的 FollowerCount
	if err := config.Database.Model(&models.User{}).Where("id = ?", toUserId).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// 取消关注用户
func (s *RelationService) CancelFollowUser(fromUserId, toUserId uint) error {
	// 获取数据库连接的引用
	db := config.Database
	// 查找已存在的关注关系
	var existingRelation models.Relation
	if err := db.Where("from_user_id = ? AND to_user_id = ?", fromUserId, toUserId).First(&existingRelation).Error; err != nil {
		return errors.New("关注关系不存在")
	}
	// 检查关注关系的取消状态
	if existingRelation.Cancel == 1 {
		return errors.New("关注关系已经是取消状态")
	}
	// 更新关注关系为取消状态
	if err := db.Model(&existingRelation).Update("cancel", 1).Error; err != nil {
		return err
	}
	//更新关注用户的 FollowCount
	if err := config.Database.Model(&models.User{}).Where("id = ?", fromUserId).Update("follow_count", gorm.Expr("follow_count - 1")).Error; err != nil {
		return err
	}
	//更新被关注用户的 FollowerCount
	if err := config.Database.Model(&models.User{}).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count - 1")).Error; err != nil {
		return err
	}
	return nil
}

// 获取关注列表
func (s *RelationService) GetFllowList(userId uint) ([]models.User, error) {
	//通过relation表查询用户的所有关注用户Id
	var relation []models.Relation
	result := config.Database.Select("to_user_id").Where("from_user_id = ?", userId).Find(&relation)
	if result.Error != nil {
		return nil, result.Error
	}
	var toUserIds []uint64
	for _, rel := range relation {
		toUserIds = append(toUserIds, uint64(rel.ToUserId))
	}
	//再通过users表返回被关注的用户的详细信息
	var users []models.User
	result = config.Database.Where("id IN (?)", toUserIds).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// 判断是否关注
func (s *RelationService) IsFollow(fromUserId uint, toUserId uint) bool {
	// 获取数据库连接的引用
	db := config.Database
	// 查找存在的关注关系
	var existingRelation models.Relation
	if err := db.Where("from_user_id = ? AND to_user_id = ? AND cancel = 0", fromUserId, toUserId).First(&existingRelation).Error; err != nil {
		return false
	}
	return true
}

// 获取粉丝列表
func (s *RelationService) GetFollowerList(userId uint) ([]models.User, error) {
	// 通过relation表查询关注了该用户的所有用户Id
	var relation []models.Relation
	result := config.Database.Select("from_user_id").Where("to_user_id = ? AND cancel = ?", userId, 0).Find(&relation)
	if result.Error != nil {
		return nil, result.Error
	}
	var followerUserIds []uint64
	for _, rel := range relation {
		followerUserIds = append(followerUserIds, uint64(rel.FromUserId))
	}
	// 再通过users表返回粉丝的详细信息
	var followers []models.User
	result = config.Database.Where("id IN (?)", followerUserIds).Find(&followers)
	if result.Error != nil {
		return nil, result.Error
	}
	return followers, nil
}

// GetFriendsList 通过用户ID获取好友用户列表
func (s *RelationService) GetFriendsList(userID uint) ([]models.User, error) {
	// 通过relation表查询用户关注的好友
	var relations []models.Relation
	result := config.Database.Where("(from_user_id = ? OR to_user_id = ?) AND cancel = ?", userID, userID, 0).Find(&relations)
	if result.Error != nil {
		return nil, result.Error
	}

	friendUserIDs := make(map[uint]bool) // 使用 map 来存储好友的用户ID，避免重复
	for _, rel := range relations {
		if rel.FromUserId == userID {
			friendUserIDs[rel.ToUserId] = true
		} else {
			friendUserIDs[rel.FromUserId] = true
		}
	}

	// 获取好友用户的详细信息
	var friendUserList []models.User
	for friendID := range friendUserIDs {
		var user models.User
		result := config.Database.First(&user, friendID)
		if result.Error != nil {
			return nil, result.Error
		}
		friendUserList = append(friendUserList, user)
	}

	return friendUserList, nil
}
