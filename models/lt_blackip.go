package models

type LtBlackip struct {
	Id         int    `xorm:"not null pk autoincr comment('主键') INT(11)"`
	Ip         string `xorm:"comment('ip地址') VARCHAR(50)"`
	Blacktime  int    `xorm:"comment('黑名单限制到期时间') INT(11)"`
	SysCreated int    `xorm:"comment('创建时间') INT(11)"`
	SysUpdated int    `xorm:"comment('更新时间') INT(11)"`
}
