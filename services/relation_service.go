package services

import (
	"errors"

	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
	"gorm.io/gorm"
)

type RelationService struct {
}

// 判断是否关注
func (s *RelationService) IsFollow(fromUserId uint, toUserId uint) bool {
	var relation models.Relation
	err := config.DB.
		Where("(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ? AND is_mutual = 1)",
			fromUserId, toUserId, toUserId, fromUserId).
		First(&relation).Error
	return err == nil
}

// 判断是否被关注
func (s *RelationService) IsFollower(fromUserId uint, toUserId uint) bool {
	var relation models.Relation
	err := config.DB.
		Where("from_user_id = ? AND to_user_id = ?", toUserId, fromUserId).
		First(&relation).Error
	return err == nil
}

// 关注用户
func (s *RelationService) FollowUser(fromUserId uint, toUserId uint) error {

	// 判断是否被关注，如果被关注，则需要将is_mutual修改成互相关注（1）
	isFollower := s.IsFollower(fromUserId, toUserId)

	isFollew := s.IsFollow(fromUserId, toUserId)

	if isFollew {
		//已经关注过，无需重复操作
		return errors.New("已关注该用户")
	}

	//如果要关注对象没有关注自己
	isMutual := uint(0)
	if isFollower {
		isMutual = 1
	}

	relation := models.Relation{
		FromUserId: fromUserId,
		ToUserId:   toUserId,
		IsMutual:   isMutual,
	}

	err := config.DB.Create(&relation).Error
	if err != nil {
		// 处理创建记录错误
		return err
	}

	//更新关注用户的 FollowCount
	err = config.DB.Model(&models.User{}).Where("id = ?", fromUserId).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error
	if err != nil {
		// 处理更新错误
		return err
	}

	//更新被关注用户的 FollowerCount
	err = config.DB.Model(&models.User{}).Where("id = ?", toUserId).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error
	if err != nil {
		// 处理更新错误
		return err
	}

	return nil
}

// 取消关注用户
func (s *RelationService) CancelFollowUser(fromUserId, toUserId uint) error {
	isFollew := s.IsFollow(fromUserId, toUserId)
	if !isFollew {
		//表示没有关注过，无需取关操作
		return errors.New("未关注该用户")
	}

	//取关用户
	err := config.DB.Where("(is_mutual = 0 OR is_mutual = 1) AND from_user_id = ? AND to_user_id = ?", fromUserId, toUserId).Delete(&models.Relation{}).Error
	if err != nil {
		// 处理删除错误
		return err
	}

	//判断toUser是否关注fromUser，并且is_mutual为1，是的话需要修改为0
	// 在关联关系表中查询数据
	var relation models.Relation
	err = config.DB.Where("from_user_id = ? AND to_user_id = ? AND is_mutual = 1", toUserId, fromUserId).First(&relation).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有关联关系数据，不需要修改 is_mutual
		} else {
			// 处理其他查询错误
			return err
		}
	} else {
		// 修改 is_mutual 值为 0
		relation.IsMutual = 0
		err = config.DB.Save(&relation).Error
		if err != nil {
			// 处理保存错误
			return err
		}
	}

	//更新关注用户的 FollowCount
	err = config.DB.Model(&models.User{}).Where("id = ?", fromUserId).Update("follow_count", gorm.Expr("follow_count - 1")).Error
	if err != nil {
		// 处理更新错误
		return err
	}

	//更新被关注用户的 FollowerCount
	err = config.DB.Model(&models.User{}).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count - 1")).Error
	if err != nil {
		// 处理更新错误
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
