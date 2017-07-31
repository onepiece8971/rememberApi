package controllers

import (
	"github.com/astaxie/beego"
	"rememberApi/models"
)

// 主页
type HomeController struct {
	beego.Controller
}

func (c *HomeController) GetBooks() {
	books := models.GetBooks()
	var out []map[string]interface{}
	var v *models.Books
	for _, v  = range books {
		user := models.GetUserByUid(v.CreateUser)
		book := map[string]interface{}{
			"BookName": v.Name,
			"Cover":    v.Cover,
			"UserName": user.Name,
			"Head":     user.Head,
		}
		out = append(out, book)
	}
	c.Data["json"] = out
	c.ServeJSON()
}