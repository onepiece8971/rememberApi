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

type Review struct {
	models.Posts
	Level int
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
	// 拼装level字段
	if len(posts) > 0 {
		postsIds := []uint32{}
		for _, v := range posts {
			postsIds = append(postsIds, v.Id)
		}
		recites := models.GetRecitesLevel(ubId, postsIds)
		results := []Review{}
		for _, v := range posts {
			results = append(results, Review{Posts: *v, Level: recites[v.Id]})
		}
		c.Data["json"] = results
	} else  {
		c.Data["json"] = []Review{}
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
	c.Data["json"] = Review{post, level}
	c.ServeJSON()
}