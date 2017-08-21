package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"rememberApi/models"
	"encoding/json"
	"fmt"
	"time"
	"strings"
	"math"
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
	qs := o.QueryTable("english")
	count, _ := qs.Count()
	page := 100
	pages := math.Ceil(float64(count) / float64(page))
	p := orm.NewOrm()
	for i := 0; i < int(pages); i++ {
		var englishes []*English
		qs.Limit(page, i * page).All(&englishes)
		for _, english := range englishes {
			var meaning []string
			jsonErr := json.Unmarshal([]byte(english.Meaning), &meaning)
			if jsonErr != nil {
				fmt.Printf("jsonErr1: %v\n", jsonErr)
			}
			var sentence []string
			jsonErr = json.Unmarshal([]byte(english.Sentence), &sentence)
			if jsonErr != nil {
				fmt.Printf("jsonErr2: %v\n", jsonErr)
			}
			post := models.Posts{
				Name: english.Name,
				UserBooksId: 1,
				Page: english.Id,
				// todo 列句
				Content: "# " + english.Name + "\n" + "#### " + english.Phsymbol + "<s>(" + english.Voice + ")\n<hide>" +
					strings.Join(meaning, "\n") + "\n![images](" + english.Images + ")\n列句\n" + strings.Join(sentence, "\n") +
					"</hide>",
			}
			now := int(time.Now().Unix())
			post.CreateTime = now
			post.UpdateTime = now
			if postId, uerr := p.Insert(&post); uerr == nil {
				booksHasPosts := models.BooksHasPosts{
					BooksId: 1,
					PostsId: uint32(postId),
				}
				_, hasErr := p.Insert(&booksHasPosts)
				if hasErr != nil {
					fmt.Printf("HasERR: %v\n", hasErr)
				}
				fmt.Printf("Success: %d\n", postId)
			} else {
				fmt.Printf("UERR: %v\n", uerr)
			}
		}
	}
}
