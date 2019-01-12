package conf

import "time"

const SysTimeform = "2006-01-02 15:04:05"
const SysTimeformShort = "2006-01-02"

var SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")
var SignSecret = "zxcvbnm"
var CookieSecret = "hellolottery"