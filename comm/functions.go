package comm

import (
	"time"
	"myproject/lottery/conf"
)

//当前时间戳
func NowUnix() int {
	return int(time.Now().In(conf.SysTimeLocation).Unix())
}
