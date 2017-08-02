package controllers

import (
	"github.com/astaxie/beego"
	"rememberApi/models"
	"fmt"
	"strconv"
)

type PostsController struct {
	beego.Controller
}

func (c *PostsController) GetPostsByUserBooksId() {
	id := c.Ctx.Input.Param(":id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	ubId := uint32(intId)
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
	type result struct {
		models.Posts
		Level int
	}
	// 拼装level字段
	if len(posts) > 0 {
		postsIds := []uint32{}
		for _, v := range posts {
			postsIds = append(postsIds, v.Id)
		}
		recites := models.GetRecitesLevel(ubId, postsIds)
		results := []result{}
		for _, v := range posts {
			results = append(results, result{Posts: *v, Level: recites[v.Id]})
		}
		c.Data["json"] = results
	} else  {
		c.Data["json"] = []result{}
	}
	c.ServeJSON()
}

func (c *PostsController) GetPostById() {
	id := c.Ctx.Input.Param(":postId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	postId := uint32(intId)
	post := models.GetPostById(postId)
	level := models.GetReciteLevel(postId)
	type result struct {
		models.Posts
		Level int
	}
	c.Data["json"] = result{post, level}
	c.ServeJSON()
}