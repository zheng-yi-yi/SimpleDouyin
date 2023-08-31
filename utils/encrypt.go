package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const CustomCost int = 16

// EncryptPassword 生成密码哈希，使用自定义的哈希迭代成本（cost）
func EncryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), CustomCost)
}

// CheckPasswordValidity 将加密后的密码和用户提供的密码进行比较
func CheckPasswordValidity(encryptedPassword, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(userPassword))
	return err == nil
}
