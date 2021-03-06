package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Users struct {
	Uid        uint32 `orm:"pk;column(uid);"`
	Name       string `orm:"size(100)"`
	Email      string
	Password   string
	Head       string
	Base
}

func init() {
	// register model
	orm.RegisterModel(new(Users))
}

func GetUserByUid(uid uint32) Users {
	o := orm.NewOrm()
	user := Users{Uid: uid}
	err := o.Read(&user)

	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}

	return user
}