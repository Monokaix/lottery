package datasource

import (
	"github.com/go-xorm/xorm"
	"fmt"
	"myproject/lottery/conf"
	"github.com/lunny/log"
	_"github.com/go-sql-driver/mysql"
	"sync"
)

var masterInstance *xorm.Engine
var dbLock sync.Mutex

func InstanceDbMaster() *xorm.Engine {
	if masterInstance != nil{
		return masterInstance
	}
	dbLock.Lock()
	defer dbLock.Unlock()

	if masterInstance != nil{
		return masterInstance
	}
	return NewDbMaster()
}
func NewDbMaster() *xorm.Engine {
	sourcename := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s/?chartset=utf8",
		conf.DbMaster.User,
		conf.DbMaster.Pwd,
		conf.DbMaster.Host,
		conf.DbMaster.Port,
		conf.DbMaster.Database,
	)

	instance, err := xorm.NewEngine(conf.DriverName, sourcename)
	if err != nil{
		log.Fatal("dbhelper NewMaster NewEngine error", err)
	}
	instance.ShowSQL(true)
	masterInstance = instance
	return masterInstance
}
