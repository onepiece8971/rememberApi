package models

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Books struct {
	Id         uint32
	Name       string
	Info       string
	Cover      string
	Types      int
	CreateUser uint32
	Base
}

func GetBooks() []*Books {
	o := orm.NewOrm()
	var books []*Books
	qs := o.QueryTable("books")
	_, err := qs.Filter("Delete", 0).All(&books)

	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	return books
}