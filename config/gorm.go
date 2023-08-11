package config

import (
	"github.com/zheng-yi-yi/simpledouyin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB // DB : 全局变量，用于保存数据库连接实例。

var GormConfig *gorm.Config // GormConfig : 全局变量，用于保存 GORM 库的配置选项。

// 初始化 GORM 配置，返回一个 GORM 配置的指针。
func InitGormConfig() *gorm.Config {
	tablePrefix := Config.Mysql.Prefix
	// 配置 GORM 库的行为
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: &schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: false,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		PrepareStmt: true,
	}
	// 返回了创建好的 gormConfig 结构体实例
	return gormConfig
}

// 初始化数据库连接：返回一个与 MySQL 数据库连接的 GORM 数据库对象。
func InitGorm() *gorm.DB {
	// 创建了一个指向 MysqlConfig 类型对象的指针 mysqlconf，它存储了 MySQL 数据库的配置信息。
	mysqlconf := Config.Mysql
	// 创建 mysqlConfig 变量，用于存储 MySQL 数据库连接的各种配置选项。
	mysqlConfig := mysql.Config{
		DriverName:                "mysql",         // 数据库驱动
		ServerVersion:             "",              // 数据库服务器的版本号
		DSN:                       mysqlconf.Dsn(), // 连接数据库所需的信息字符串，通常包括用户名、密码、主机地址、端口号和数据库名。
		Conn:                      nil,             // 如果已经有一个现有的数据库连接，可以将它赋值给这个选项，否则设置为 nil。
		SkipInitializeWithVersion: false,           // 是否跳过根据数据库版本进行初始化。
		DefaultStringSize:         0,               // 默认的字符串长度
		DefaultDatetimePrecision:  nil,             // 默认的日期时间精度
		DisableDatetimePrecision:  false,           // 是否禁用日期时间精度
		DontSupportRenameIndex:    false,           // 是否禁用日期时间精度
		DontSupportRenameColumn:   false,           // 是否不支持重命名索引
		DontSupportForShareClause: false,           // 是否不支持重命名列
	}
	// 创建数据库连接
	// gorm.Open 函数被调用来创建一个 GORM 数据库连接。
	db, err := gorm.Open(mysql.New(mysqlConfig), GormConfig)
	if err != nil {
		panic(err)
	}
	// 最后，函数返回创建的 GORM 数据库连接对象 db。
	return db
}

// CreateTable ：自动创建数据库表格或更新已有表格的结构。
// AutoMigrate 接受一个或多个模型对象作为参数，它会根据这些模型的定义来生成对应的数据库表格。
func CreateTable(db *gorm.DB) error {
	err := db.AutoMigrate(
		models.User{},
		models.Video{},
		models.Favorite{},
		models.Comment{},
		models.Relation{},
	)
	if err != nil {
		return err
	}
	// 初始化用户表：
	var count_User int64
	db.Model(&models.User{}).Count(&count_User)
	if count_User == 0 {
		user1 := models.User{
			UserName:      "User1",
			PassWord:      "123456",
			Status:        1,
			FollowCount:   0,
			FollowerCount: 1,
		}
		db.Create(&user1)
		user2 := models.User{
			UserName:      "User2",
			PassWord:      "123456",
			Status:        1,
			FollowCount:   1,
			FollowerCount: 0,
		}
		db.Create(&user2)
	}
	// 初始化视频表：
	var count_Video int64
	db.Model(&models.Video{}).Count(&count_Video)
	if count_Video == 0 {
		videos := []*models.Video{
			{UserId: 1, PlayUrl: "videos/1.mp4", CoverUrl: "images/1.jpg", FavoriteCount: 1, Description: "邀你观看2023稀土开发者大会!11场论坛，60位技术大咖，涵盖AIGC与大模型、大前端、音视频等前沿技术资讯"},
			{UserId: 1, PlayUrl: "videos/2.mp4", CoverUrl: "images/2.jpg", FavoriteCount: 1, Description: "那些仅凭半句就封神的诗句"},
			{UserId: 1, PlayUrl: "videos/3.mp4", CoverUrl: "images/3.jpg", FavoriteCount: 1, Description: "登昆仑兮食玉英，与天地兮同寿与日月兮齐光。#汉服之美在华夏"},
			{UserId: 1, PlayUrl: "videos/4.mp4", CoverUrl: "images/4.jpg", FavoriteCount: 1, Description: "落日沉溺于橘色的海，晚风沦陷于赤城的爱"},
			{UserId: 1, PlayUrl: "videos/5.mp4", CoverUrl: "images/5.jpg", FavoriteCount: 1, Description: "五档 启动！ #太阳神尼卡登场"},
			{UserId: 1, PlayUrl: "videos/6.mp4", CoverUrl: "images/6.jpg", FavoriteCount: 1, Description: "好像做了一场短暂的山水梦"},
		}
		db.Create(&videos)
	}
	// 初始化点赞表：
	var count_Favorite int64
	db.Model(&models.Favorite{}).Count(&count_Favorite)
	if count_Favorite == 0 {
		favorite := []*models.Favorite{
			{UserId: 1, VideoId: 1, Status: 1},
			{UserId: 1, VideoId: 2, Status: 1},
			{UserId: 1, VideoId: 3, Status: 1},
			{UserId: 1, VideoId: 4, Status: 1},
			{UserId: 1, VideoId: 5, Status: 1},
		}
		db.Create(&favorite)
	}
	// 初始化评论表：
	var count_Comment int64
	db.Model(&models.Comment{}).Count(&count_Comment)
	if count_Comment == 0 {
		comment := []*models.Comment{
			{VideoId: 1, Content: "来学习了", ToUserId: 1, UserId: 2},
			{VideoId: 2, Content: "为天地立心，为生民立命，为往圣继绝学，为万世开太平", ToUserId: 1, UserId: 2},
			{VideoId: 3, Content: "变装那一瞬间好高级", ToUserId: 1, UserId: 2},
			{VideoId: 4, Content: "我对着日落许愿，希望你永远快乐", ToUserId: 1, UserId: 2},
			{VideoId: 5, Content: "年少的梦终将绽放于盛夏，解放之鼓必将响彻整个夏天", ToUserId: 1, UserId: 2},
			{VideoId: 5, Content: "说走就走的旅行...", ToUserId: 1, UserId: 2},
		}
		db.Create(&comment)
	}
	// 初始化关注关系表：
	var count_Relation int64
	db.Model(&models.Relation{}).Count(&count_Relation)
	if count_Relation == 0 {
		comment := []*models.Relation{
			{FromUserId: 2, ToUserId: 1, IsMutual: 0},
		}
		db.Create(&comment)
	}
	return nil
}
