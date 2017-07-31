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
	postsIds := models.GetRecitePostsIdsByUbId(uintId)
	c.Data["json"] = postsIds
	c.ServeJSON()
}