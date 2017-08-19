package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

type UserBooks struct {
	Id          uint32
	Uid         uint32
	BooksId     uint32
	Name        string
	Cover       string
	Info        string
	IsMemory    string
	UsedPages   int
	PagesUptime int `orm:"size(10)"`
	Base
}

func init() {
	// register model
	orm.RegisterModel(new(UserBooks))
}

func GetUserBooksByUid(uid uint32, isAll bool) []*UserBooks {
	o := orm.NewOrm()
	var userBooks []*UserBooks
	qs := o.QueryTable("user_books")
	qs = qs.Filter("delete", 0).Filter("uid", uid)
	var err error
	if isAll {
		_, err = qs.All(&userBooks)
	} else  {
		_, err = qs.Filter("is_memory", true).All(&userBooks)
	}
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	return userBooks
}

func GetUserBookById(id uint32) UserBooks {
	o := orm.NewOrm()
	userBook := UserBooks{Id: id}
	err := o.Read(&userBook)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	return userBook
}

func AddUserBooks(uid, booksId uint32) (id int64, err error) {
	o := orm.NewOrm()
	userBook := UserBooks{
		Uid: uid,
		BooksId: booksId,
	}
	now := int(time.Now().Unix())
	userBook.CreateTime = now
	userBook.UpdateTime = now
	id, err = o.Insert(&userBook)
	return
}