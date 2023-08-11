package utils

import (
	"strconv"
	"time"

	"github.com/zheng-yi-yi/simpledouyin/config"
)

// 返回一个格式化的时间字符串。
func CalculateStartTime(lastTimestamp string) string {
	if lastTimestamp != "" {
		timestamp, err := strconv.ParseInt(lastTimestamp, 10, 64)
		if err == nil {
			if timestamp > 1000000000000 {
				timestamp /= 1000
			}
			return time.Unix(timestamp, 0).Format(config.DateTime)
		}
	}
	return time.Now().Format(config.DateTime)
}
