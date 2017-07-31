package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Posts struct {
	Id          uint32
	Name        string
	Content     string
	UserBooksId uint32
	Page        int
	Share       string
	Base
}

type BooksHasPosts struct {
	BooksId uint32 `orm:"pk"`
	PostsId uint32
}

func GetPostsByUserBooksId(ubId uint32) []*Posts {
	o := orm.NewOrm()
	var posts []*Posts
	qs := o.QueryTable("posts")
	_, err := qs.Filter("delete", 0).
		Filter("UserBooksId", ubId).
		All(&posts)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	return posts
}

func GetPostsIdsByBooksId(bid uint32) (postsIds []uint32) {
	o := orm.NewOrm()
	var bps []*BooksHasPosts
	qs := o.QueryTable("books_has_posts")
	_, err := qs.Filter("books_id", bid).All(&bps)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	for _, v := range bps {
		postsIds = append(postsIds, v.PostsId)
	}
	return
}

func GetPostsByBooksId(bid uint32) []*Posts {
	o := orm.NewOrm()
	postsIds := GetPostsIdsByBooksId(bid)
	var posts []*Posts
	qs := o.QueryTable("posts")
	_, err := qs.Filter("id__in", postsIds).All(&posts)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	return posts
}

func GetPostsByBooksIdExcludePage(bid uint32, pages []int) []*Posts {
	o := orm.NewOrm()
	postsIds := GetPostsIdsByBooksId(bid)
	var posts []*Posts
	qs := o.QueryTable("posts")
	_, err := qs.Exclude("page__in", pages).Filter("id__in", postsIds).All(&posts)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	return posts
}

func GetPostById(id uint32) Posts {
	o := orm.NewOrm()
	p := Posts{Id: id}
	err := o.Read(&p)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	return p
}

func GetPostsByIds(ids []uint32) []*Posts {
	o := orm.NewOrm()
	var posts []*Posts
	qs := o.QueryTable("posts")
	_, err := qs.Filter("id__in", ids).All(&posts)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	return posts
}