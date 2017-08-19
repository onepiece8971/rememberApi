package models

import (
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
	PageNum    int
	Base
}

func init() {
	// register model
	orm.RegisterModel(new(Books))
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

func GetBooksPageNum(id uint32) int {
	o := orm.NewOrm()
	book := Books{Id: id}
	err := o.Read(&book)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	return book.PageNum
}