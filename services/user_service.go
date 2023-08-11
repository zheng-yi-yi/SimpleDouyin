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
	err := config.DB.Where("username = ?", username).First(&user).Error
	if err == nil {
		return models.User{}, errors.New("用户名已注册")
	}
	user = models.User{
		UserName: username,
		PassWord: password,
	}
	err = config.DB.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// 用户登录服务
func (userService *UserService) Login(username, password string) (models.User, error) {
	var user models.User
	err := config.DB.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return models.User{}, errors.New("用户名或密码错误")
	} else if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// 通过用户ID获取用户信息
func (userService *UserService) GetUserInfoById(userId uint) (models.User, error) {
	var user models.User
	err := config.DB.Where("id = ?", userId).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return models.User{}, errors.New("用户不存在")
	} else if err != nil {
		return models.User{}, err
	}
	return user, nil
}
