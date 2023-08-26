package setup

import (
	"errors"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// MysqlConfig 用于存储 MySQL 数据库的配置信息
type MysqlConfig struct {
	Host     string `json:"host" yaml:"host"`               // 主机
	Port     string `json:"port" yaml:"port"`               // 端口
	Dbname   string `json:"db_name" mapstructure:"db_name"` // 数据库名
	Username string `json:"username" yaml:"username"`       // 用户名
	Password string `json:"password" yaml:"password"`       // 密码
	Charset  string `json:"charset" yaml:"charset"`         // 字符集
	Prefix   string `json:"prefix" yaml:"prefix"`           // 前缀
}

// 返回一个DSN字符串，它是连接 MySQL 数据库所需的参数字符串
func (m *MysqlConfig) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?charset=" + m.Charset + "&parseTime=True&loc=Local"
}

type ServerConfig struct {
	Port   int    `json:"port,omitempty" yaml:"port"`
	DbType string `json:"db_type,omitempty" yaml:"db_type"`
}

// 包含配置信息的结构体
type Conf struct {
	Mysql  *MysqlConfig  // MysqlConfig 结构体定义了 MySQL 数据库的配置信息
	Server *ServerConfig // ServerConfig 结构体定义了一些服务器配置的信息。
}

var Config *Conf // Config 指向一个包含配置信息的结构体

// 解析根目录下的配置文件 config.yaml，解析后的配置统一通过 config.Config 获取
func InitConfig() {
	// 创建一个新的 Viper 实例 （Viper 是一个用于配置管理的库，用于读取和解析配置文件。）
	v := viper.New()
	// 尝试从根目录下的 "config.yaml" 文件读取配置信息。
	v.SetConfigFile("config.yaml")
	// 设置配置文件的类型为 YAML 格式。
	v.SetConfigType("yaml")
	// 尝试读取配置文件内容到 Viper 实例中，并将可能的错误保存在 err 变量中
	err := v.ReadInConfig()
	if err != nil {
		panic(errors.New("打开配置文件出错，请检查根目录是否存在 config.yaml 文件 "))
	}
	// 尝试将配置文件的内容解析到 config.Config 变量中。
	if err := v.Unmarshal(&Config); err != nil {
		panic(errors.New("配置文件读取失败，请检查配置项与官方配置是否一致! "))
	}
	// 监听配置文件修改 : 在配置文件发生变化时，Viper 将会重新加载配置。
	viper.WatchConfig()
	// 设置一个回调函数
	// 当配置文件发生变化时，回调函数会被调用。
	// 在回调函数内部，同样尝试将配置文件内容解析到 config.Config 变量中
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := v.Unmarshal(&Config); err != nil {
			println(errors.New("配置文件读取失败，请检查配置项与官方配置是否一致! "))
		}
	})
}
