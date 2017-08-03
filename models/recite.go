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

func GetReciteByUbId(ubId uint32) []*Recite {
	o := orm.NewOrm()
	var recites []*Recite
	_, err := o.Raw(
		"SELECT posts_id, level FROM recite WHERE user_books_id = ? AND `delete` = 0 AND level_time <> 0 AND (? - update_time) > level_time ORDER BY level",
		ubId,
		time.Now().Unix(),
	).QueryRows(&recites)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	return recites
}

func GetRecitePostsIdsByUbId(ubId uint32) []uint32 {
	recites := GetReciteByUbId(ubId)
	postsIds := []uint32{}
	for _, v := range recites {
		postsIds = append(postsIds, v.PostsId)
	}
	return postsIds
}

func GetRecitesByUbIdAndPostsIds(ubId uint32, postsIds []uint32) []*Recite {
	o := orm.NewOrm()
	recites := []*Recite{}
	qs := o.QueryTable("recite")
	_, err := qs.Filter("user_books_id", ubId).Filter("posts_id__in", postsIds).All(&recites)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	return recites
}

func GetRecitesLevel(ubId uint32, postsIds []uint32) map[uint32]int {
	recites := GetRecitesByUbIdAndPostsIds(ubId, postsIds)
	levels := map[uint32]int{}
	for _, v := range recites {
		levels[v.PostsId] = v.Level
	}
	return levels
}

func GetReciteLevel(postsId uint32) int {
	o := orm.NewOrm()
	recite := Recite{PostsId: postsId}
	err := o.Read(&recite, "PostsId")
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	return recite.Level
}

func AddRecite(ubId, postId uint32) (id int64, err error) {
	o := orm.NewOrm()
	recite := Recite{
		UserBooksId: ubId,
		PostsId:     postId,
		Level:       1,
		LevelTime:   LevelMap[1],
	}
	now := int(time.Now().Unix())
	recite.CreateTime = now
	recite.UpdateTime = now
	id, err = o.Insert(&recite)
	return
}