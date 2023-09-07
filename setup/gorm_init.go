package setup

import (
	"github.com/zheng-yi-yi/simpledouyin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 初始化 GORM 配置，返回一个 GORM 配置的指针
func InitGormConfig() *gorm.Config {
	// 配置 GORM 库的行为
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: &schema.NamingStrategy{
			TablePrefix:   "dy_",
			SingularTable: false,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		PrepareStmt: true,
	}
	return gormConfig
}

// 初始化数据库连接：返回一个与 MySQL 数据库连接的 GORM 数据库对象。
func InitGorm() *gorm.DB {
	mysqlconf := Config.Mysql
	mysqlConfig := mysql.Config{
		DriverName:                "mysql",
		ServerVersion:             "",
		DSN:                       mysqlconf.Dsn(),
		Conn:                      nil,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         0,
		DefaultDatetimePrecision:  nil,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   false,
		DontSupportForShareClause: false,
	}
	// 创建数据库连接
	db, err := gorm.Open(mysql.New(mysqlConfig), config.GormConfig)
	if err != nil {
		panic(err)
	}
	return db
}
