package controllers

import (
	"github.com/astaxie/beego"
	"rememberApi/models"
	"fmt"
	"strconv"
)

type ReciteController struct {
	beego.Controller
}

func (c *ReciteController) GetRecitesByUserBooksId() {
	id := c.Ctx.Input.Param(":id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	uintId := uint32(intId)
	recites := models.GetReciteByUbId(uintId)
	results := []Review{}
	if len(recites) > 0 {
		postsIds := []uint32{}
		levels := map[uint32]models.ReciteLevel{}
		for _, v := range recites {
			postsIds = append(postsIds, v.PostsId)
			levels[v.PostsId] = models.ReciteLevel{Id: v.Id, Level: v.Level}
		}
		posts := models.GetPostsByIds(postsIds)
		for _, v := range posts {
			results = append(results, Review{Posts: *v, ReciteId: levels[v.Id].Id, Level: levels[v.Id].Level})
		}
	}
	c.Data["json"] = results
	c.ServeJSON()
}

func (c *ReciteController) AddRecite() {
	ubId := c.Ctx.Input.Param(":ubId")
	intUbId, err := strconv.Atoi(ubId)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	uintUbId := uint32(intUbId)
	postId := c.Ctx.Input.Param(":postId")
	intPostId, postIdErr := strconv.Atoi(postId)
	if postIdErr != nil {
		fmt.Printf("ERR: %v\n", postIdErr)
	}
	uintPostId := uint32(intPostId)
	id, addErr := models.AddRecite(uintUbId, uintPostId)
	if addErr != nil {
		fmt.Printf("ERR: %v\n", addErr)
	}
	c.Data["json"] = id
	c.ServeJSON()
}

func (c *ReciteController) upLevel(isForget bool) {
	intId, err := c.GetInt("id")
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	uintId := uint32(intId)
	n := models.UpLevelById(uintId, isForget)
	c.Data["json"] = n
	c.ServeJSON()
}

func (c *ReciteController) Remember() {
	c.upLevel(false)
}

func (c *ReciteController) Forget() {
	c.upLevel(true)
}