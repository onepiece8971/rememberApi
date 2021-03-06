package models

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type Base struct {
	CreateTime int
	UpdateTime int
	Delete     string
}

var LevelMap = map[int]int{
	1: 30,
	2: 60,
	3: 300,
	4: 1800,
	5: 43200,
	6: 86400,
	7: 172800,
	8: 345600,
	9: 604800,
	10: 1296000,
	11: 2592000,
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("database") + beego.AppConfig.String("dataname"))
}