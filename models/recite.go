package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

type Recite struct {
	Id          uint32
	UserBooksId uint32
	PostsId     uint32
	Level       int
	LevelTime   int
	Base
}

func GetRecitePostsIdsByUbId(ubId uint32) []uint32 {
	o := orm.NewOrm()
	var recite []*Recite
	_, err := o.Raw(
		"SELECT posts_id FROM recite WHERE user_books_id = ? AND `delete` = 0 AND level_time <> 0 AND (? - update_time) > level_time ORDER BY level",
		ubId,
		time.Now().Unix(),
	).QueryRows(&recite)
	postsIds := []uint32{}
	for _, v := range recite {
		postsIds = append(postsIds, v.PostsId)
	}
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	return postsIds
}

func AddRecite(ubId, postId uint32, level int) (id int64, err error) {
	o := orm.NewOrm()
	recite := Recite{
		UserBooksId: ubId,
		PostsId: postId,
		Level: level,
		LevelTime: LevelMap[level],
	}
	now := int(time.Now().Unix())
	recite.CreateTime = now
	recite.UpdateTime = now
	id, err = o.Insert(&recite)
	return
}