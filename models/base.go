package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type Base struct {
	CreateTime int
	UpdateTime int
	Delete     string
}

var LevelMap = map[int]int{
	1: 300,
}

func init() {
	// register model
	orm.RegisterModel(new(Books), new(UserBooks), new(Users), new(Posts), new(BooksHasPosts))

	// set default database
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("database"))
}