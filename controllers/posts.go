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
	ReciteId uint32
	Level    int
}

func (c *PostsController) GetPostsByUserBooksId() {
	id := c.Ctx.Input.Param(":id")
	page := c.Ctx.Input.Param(":page")
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("intIdERR: %v\n", err)
	}
	intPage, err := strconv.Atoi(page)
	if err != nil {
		fmt.Printf("intPageERR: %v\n", err)
	}
	ubId := uint32(intId)
	var posts []*models.Posts
	userBook := models.GetUserBookById(ubId)
	if userBook.PagesUptime != 0 {
		posts = models.GetPostsByUserBooksId(ubId, intPage)
		if len(posts) < userBook.UsedPages {
			pages := []int{}
			for _, v := range posts {
				pages = append(pages, v.Page)
			}
			posts2 := models.GetPostsByBooksIdExcludePage(userBook.BooksId, pages, intPage)
			all := make([]*models.Posts, len(posts) + len(posts2))
			copy(all, posts)
			copy(all[len(posts):], posts2)
			posts = all
		}
	} else {
		posts = models.GetPostsByBooksId(userBook.BooksId, intPage)
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
			results = append(results, Review{Posts: *v, ReciteId: recites[v.Id].Id, Level: recites[v.Id].Level})
		}
		c.Data["json"] = results
	} else  {
		c.Data["json"] = []Review{}
	}
	c.ServeJSON()
}

func (c *PostsController) GetPostByUbIdAndPage() {
	ubId, err := c.GetUint32(":ubId")
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	pageId, err := c.GetInt(":pageId")
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	post := models.GetPostByUbIdPage(ubId, pageId)
	reciteId, level := models.GetReciteLevel(ubId, post.Id)
	c.Data["json"] = Review{post, reciteId, level}
	c.ServeJSON()
}

func (c *PostsController) GetPostById() {
	id := c.Ctx.Input.Param(":ubId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	ubId := uint32(intId)
	postId := c.Ctx.Input.Param(":postId")
	intPostId, postIdErr := strconv.Atoi(postId)
	if err != nil {
		fmt.Printf("ERR: %v\n", postIdErr)
	}
	uintPostId := uint32(intPostId)
	post := models.GetPostById(uintPostId)
	reciteId, level := models.GetReciteLevel(ubId, uintPostId)
	c.Data["json"] = Review{post, reciteId, level}
	c.ServeJSON()
}