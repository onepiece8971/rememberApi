package controllers

import (
	"github.com/astaxie/beego"
	"rememberApi/models"
	"fmt"
)

type PostsController struct {
	beego.Controller
}

func (c *PostsController) GetPostsByUserBooksId() {
	id, err := c.GetInt("id")
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	ubId := uint32(id)
	var posts []*models.Posts
	userBook := models.GetUserBookById(ubId)
	if userBook.PagesUptime != 0 {
		posts = models.GetPostsByUserBooksId(ubId)
		if len(posts) < userBook.UsedPages {
			pages := []int{}
			for _, v := range posts {
				pages = append(pages, v.Page)
			}
			posts2 := models.GetPostsByBooksIdExcludePage(userBook.BooksId, pages)
			all := make([]*models.Posts, len(posts) + len(posts2))
			copy(all, posts)
			copy(all[len(posts):], posts2)
			posts = all
		}
	} else {
		posts = models.GetPostsByBooksId(userBook.BooksId)
	}
	c.Data["json"] = posts
	c.ServeJSON()
}