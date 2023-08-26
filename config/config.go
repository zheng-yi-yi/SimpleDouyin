package config

import "gorm.io/gorm"

const (
	LOCAL_IP_ADDRESS        = "xxx.xxx.xxx.xxx"     // 填入本机 IP 地址
	VIDEO_STREAM_BATCH_SIZE = 30                    // 每次获取视频流的数量限制
	DATETIME_FORMAT         = "2006-01-02 15:04:05" // 固定的时间格式
	SHORT_DATE_FORMAT       = "01-02"               // 短日期格式的字符串
)

var (
	Database   *gorm.DB     // Database: 全局变量，用于保存数据库连接实例
	GormConfig *gorm.Config // GormConfig: 全局变量，用于保存 GORM 库的配置选项
)
