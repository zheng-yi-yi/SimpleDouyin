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
	db := config.DB
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
	if err := config.DB.Model(&models.User{}).Where("id = ?", fromUserId).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
		return err
	}
	// 更新被关注用户的 FollowerCount
	if err := config.DB.Model(&models.User{}).Where("id = ?", toUserId).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

// 取消关注用户
func (s *RelationService) CancelFollowUser(fromUserId, toUserId uint) error {
	// 获取数据库连接的引用
	db := config.DB
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
	if err := config.DB.Model(&models.User{}).Where("id = ?", fromUserId).Update("follow_count", gorm.Expr("follow_count - 1")).Error; err != nil {
		return err
	}
	//更新被关注用户的 FollowerCount
	if err := config.DB.Model(&models.User{}).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count - 1")).Error; err != nil {
		return err
	}
	return nil
}

// 获取关注列表
func (s *RelationService) GetFllowList(userId uint) ([]models.User, error) {
	//通过relation表查询用户的所有关注用户Id
	var relation []models.Relation
	result := config.DB.Select("to_user_id").Where("from_user_id = ?", userId).Find(&relation)
	if result.Error != nil {
		return nil, result.Error
	}
	var toUserIds []uint64
	for _, rel := range relation {
		toUserIds = append(toUserIds, uint64(rel.ToUserId))
	}
	//再通过users表返回被关注的用户的详细信息
	var users []models.User
	result = config.DB.Where("id IN (?)", toUserIds).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// 判断是否关注
func (s *RelationService) IsFollow(fromUserId uint, toUserId uint) bool {
	// 获取数据库连接的引用
	db := config.DB
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
	result := config.DB.Select("from_user_id").Where("to_user_id = ? AND cancel = ?", userId, 0).Find(&relation)
	if result.Error != nil {
		return nil, result.Error
	}
	var followerUserIds []uint64
	for _, rel := range relation {
		followerUserIds = append(followerUserIds, uint64(rel.FromUserId))
	}
	// 再通过users表返回粉丝的详细信息
	var followers []models.User
	result = config.DB.Where("id IN (?)", followerUserIds).Find(&followers)
	if result.Error != nil {
		return nil, result.Error
	}
	return followers, nil
}
