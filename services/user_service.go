package services

import (
	"errors"

	"github.com/zheng-yi-yi/simpledouyin/config"
	"github.com/zheng-yi-yi/simpledouyin/models"
	"gorm.io/gorm"
)

type UserService struct {
}

// 用户注册服务
func (userService *UserService) Register(username, password string) (models.User, error) {
	if len(password) <= 5 {
		return models.User{}, errors.New("密码长度不能小于5")
	}
	var user models.User
	err := config.Database.Where("user_name = ?", username).First(&user).Error
	if err == nil {
		return models.User{}, errors.New("用户名已注册")
	}
	user = models.User{
		UserName:        username,
		PassWord:        password,
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
