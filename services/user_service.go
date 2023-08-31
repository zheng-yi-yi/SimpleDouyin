package services

import (
	"errors"

	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
	"github.com/zheng-yi-yi/simpledouyin/utils"
	"gorm.io/gorm"
)

type UserService struct {
}

// 用户注册服务
func (userService *UserService) Register(username, password string) (models.User, error) {
	var user models.User
	err := config.Database.Where("user_name = ?", username).First(&user).Error
	if err == nil {
		return models.User{}, errors.New("用户名已注册")
	}

	// 加密用户密码
	encryptedPassword, err := utils.EncryptPassword(password)
	if err != nil {
		return models.User{}, err
	}

	user = models.User{
		UserName:        username,
		PassWord:        string(encryptedPassword), // 存储加密后的密码
		FollowCount:     0,
		FollowerCount:   0,
		FavoriteCount:   0,
		Avatar:          config.DEFAULT_USER_AVATAR_URL,
		BackgroundImage: config.DEFAULT_USER_BG_IMAGE_URL,
		Signature:       config.DEFAULT_USER_BIO,
		TotalFavorited:  "0",
		WorkCount:       0,
	}
	err = config.Database.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// 用户登录服务
func (userService *UserService) Login(username, password string) (models.User, error) {
	var user models.User
	err := config.Database.Where("user_name = ? AND pass_word = ?", username, password).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return models.User{}, errors.New("用户名或密码错误")
	} else if err != nil {
		return models.User{}, err
	}
	return user, nil
}
