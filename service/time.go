package service

import (
	"fmt"
	"time"
)

type TimeService struct {

}

// 格式化时间 返回 Y-m-d H:i:s
func (timeService *TimeService) GetStringTime() string {
	formatTime := fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"))
	return formatTime
}

// 返回时间戳(秒)
func (timeService TimeService) GetIntTime() int64 {
	return  time.Now().Unix()
}