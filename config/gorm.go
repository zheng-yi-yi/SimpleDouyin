package config

import (
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
