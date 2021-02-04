package service

import (
	"fmt"
	"time"
)


// 格式化时间 返回 Y-m-d H:i:s
func GetStringTime() string {
	formatTime := fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"))
	return formatTime
}

// 返回时间戳(秒)
func GetIntTime() int64 {
	return  time.Now().Unix()
}