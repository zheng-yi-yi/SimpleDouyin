package setup

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// MysqlConfig , 存储 MySQL 数据库的配置信息
type MysqlConfig struct {
	Host     string `yaml:"host"`            // 主机
	Port     string `yaml:"port"`            // 端口
	Username string `yaml:"username"`        // 用户名
	Password string `yaml:"password"`        // 密码
	Dbname   string `mapstructure:"db_name"` // 数据库名
}

// Dsn , 返回数据库所需的参数字符串
func (m *MysqlConfig) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
}

type Conf struct {
	Mysql *MysqlConfig // MysqlConfig 结构体定义了 MySQL 数据库的配置信息
}

var Config *Conf // Config 指向一个包含配置信息的结构体

// 解析根目录下的配置文件 config.yaml，解析后的配置统一通过 setupConfig 获取
func InitConfig() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Error opening the configuration file: %v", err)
		return
	}
	// 尝试将配置文件的内容解析到 config.Config 变量中。
	if err := v.Unmarshal(&Config); err != nil {
		log.Printf("Configuration file reading failed: %v", err)
		return
	}
	fmt.Printf("配置文件读取结果：%v", *Config)
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := v.Unmarshal(&Config); err != nil {
			log.Printf("Configuration file reading failed: %v", err)
			return
		}
	})
}
