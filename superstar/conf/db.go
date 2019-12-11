package conf

type DbConf struct {
	Host     string
	Port     int
	UserName string
	Password string
	DataBase string
}

const (
	DriverName = "mysql"
	Loc        = "Asia/Shanghai"
	ConnStr    = "%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=%s&parseTime=true"
)

var MasterDbConfig DbConf = DbConf{
	Host:     "127.0.0.1",
	Port:     3306,
	UserName: "root",
	Password: "123456",
	DataBase: "yay",
}

var SlaveDbConfig DbConf = DbConf{
	Host:     "127.0.0.1",
	Port:     3306,
	UserName: "root",
	Password: "123456",
	DataBase: "yay",
}
