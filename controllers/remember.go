package controllers

import (
	"github.com/astaxie/beego"
	"rememberApi/models"
	"fmt"
	"strconv"
)

type RememberController struct {
	beego.Controller
}

func (c *RememberController) GetMemoryUserBooksByUid() {
	uid := c.Ctx.Input.Param(":uid")
	intUid, err := strconv.Atoi(uid)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	uintUid := uint32(intUid)
	userBooks := models.GetMemoryUserBooksByUid(uintUid)
	type rememberBooks struct {
		models.UserBooks
		PageNum int
	}
	rbs := []rememberBooks{}
	for _, v := range userBooks{
		pageNum := models.GetBooksPageNum(v.BooksId)
		rbs = append(rbs, rememberBooks{UserBooks: *v, PageNum: pageNum})
	}
	c.Data["json"] = rbs
	c.ServeJSON()
}