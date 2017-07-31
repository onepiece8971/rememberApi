package controllers

import (
	"github.com/astaxie/beego"
	"rememberApi/models"
	"fmt"
)

type RememberController struct {
	beego.Controller
}

func (c *RememberController) GetMemoryUserBooksByUid() {
	intUid, err := c.GetInt("uid")
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	uid := uint32(intUid)
	c.Data["json"] = models.GetMemoryUserBooksByUid(uid)
	c.ServeJSON()
}