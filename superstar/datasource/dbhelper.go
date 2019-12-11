/**
 * 获取数据库连接
 */
package datasource

import (
	"firstGo/superstar/conf"
	"fmt"
	"log"
	"net/url"
	"sync"
	"xorm.io/xorm"

	_ "github.com/go-sql-driver/mysql"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

// 主库，单例
func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}
	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf(conf.ConnStr, c.UserName, c.Password, c.Host, c.Port, c.DataBase, url.QueryEscape(conf.Loc))
	engine, err := xorm.NewEngine(conf.DriverName, driveSource)
	if err != nil {
		log.Fatal("dbHelper.DbInstanceMaster,", err)
		return nil
	}
	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(true)
	engine.SetTZLocation(conf.SysTimeLocation)

	// 性能优化的时候才考虑，加上本机的SQL缓存
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)

	masterEngine = engine
	return engine
}

// 从库，单例
func InstanceSlave() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}
	c := conf.SlaveDbConfig
	engine, err := xorm.NewEngine(conf.DriverName,
		fmt.Sprintf(conf.ConnStr, c.UserName, c.Password, c.Host, c.Port, c.DataBase, url.QueryEscape(conf.Loc)))
	if err != nil {
		log.Fatal("dbHelper", "DbInstanceMaster", err)
		return nil
	}
	engine.SetTZLocation(conf.SysTimeLocation)
	slaveEngine = engine
	return engine
}
