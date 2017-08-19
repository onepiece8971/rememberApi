package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"rememberApi/models"
	"encoding/json"
	"fmt"
	"time"
	"strings"
)

type English struct {
	Id int
	Name string
	Phsymbol string
	Voice string
	Images string
	Meaning string
	Sentence string
	models.Base
}

func init() {
	// register model
	orm.RegisterModel(new(English))

	// set default database
	orm.RegisterDataBase("spider", "mysql", beego.AppConfig.String("database") + "spider")
}

func getOrm(aliasName string) orm.Ormer {
	o := orm.NewOrm()
	o.Using(aliasName)
	return o
}

func main() {
	o := getOrm("spider")
	english := English{Id: 11}
	err := o.Read(&english)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	var meaning []string
	jsonErr := json.Unmarshal([]byte(english.Meaning), &meaning)
	if jsonErr != nil {
		fmt.Printf("ERR: %v\n", jsonErr)
	}
	var sentence []string
	jsonErr = json.Unmarshal([]byte(english.Sentence), &sentence)
	if jsonErr != nil {
		fmt.Printf("ERR: %v\n", jsonErr)
	}
	post := models.Posts{
		Name: english.Name,
	}
	p := orm.NewOrm()
	if e := p.Read(&post, "name"); e == nil {
		post.UserBooksId = 1
		post.Page = english.Id
		post.Content = "# " + english.Name + "\n" + "#### " + english.Phsymbol + "<s>(" + english.Voice + ")\n" +
			strings.Join(meaning, "\n") + "\n![images](" + english.Images + ")\n列句\n" + strings.Join(sentence, "\n")
		now := int(time.Now().Unix())
		post.UpdateTime = now
		if num, uerr := p.Update(&post); uerr == nil {
			fmt.Println(num)
		}
	}
}
