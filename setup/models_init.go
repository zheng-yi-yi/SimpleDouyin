package setup

import (
	"github.com/zheng-yi-yi/simpledouyin/models"
	"gorm.io/gorm"
)

// CreateTable ：自动创建（或更新）数据库表格，并初始化表格数据
func CreateTable(db *gorm.DB) error {
	err := db.AutoMigrate(
		models.User{},
		models.Video{},
		models.Favorite{},
		models.Comment{},
		models.Relation{},
		models.Message{},
	)
	if err != nil {
		return err
	}
	tableInit(db)
	return nil
}

// 表格-样例数据初始化
func tableInit(db *gorm.DB) {
	initUsers(db)
	initVideos(db)
	initFavorites(db)
	initComments(db)
	initRelations(db)
	initMessages(db)
}
