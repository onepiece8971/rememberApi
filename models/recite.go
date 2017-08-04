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

type ReciteLevel struct {
	Id    uint32
	Level int
}

func GetReciteByUbId(ubId uint32) []*Recite {
	o := orm.NewOrm()
	var recites []*Recite
	_, err := o.Raw(
		"SELECT id, posts_id, level FROM recite WHERE user_books_id = ? AND `delete` = 0 AND level_time <> 0 AND (? - update_time) > level_time ORDER BY level",
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

func GetRecitesLevel(ubId uint32, postsIds []uint32) (re map[uint32]ReciteLevel) {
	recites := GetRecitesByUbIdAndPostsIds(ubId, postsIds)
	re = map[uint32]ReciteLevel{}
	for _, v := range recites {
		re[v.PostsId] = ReciteLevel{v.Id, v.Level}
	}
	return
}

func GetReciteLevel(ubId uint32, postsId uint32) (reciteId uint32, level int) {
	o := orm.NewOrm()
	recite := Recite{}
	qs := o.QueryTable("recite")
	err := qs.Filter("user_books_id", ubId).Filter("posts_id", postsId).One(&recite)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	reciteId = recite.Id
	level = recite.Level
	return
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

func UpLevelById(id uint32, isForget bool) int64 {
	o := orm.NewOrm()
	recite := Recite{Id: id}
	err := o.Read(&recite)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	} else {
		if isForget {
			if recite.Level == 2 {
				return 1
			}
			recite.Level = 2
		} else {
			if recite.Level >= 9 {
				return 1
			}
			recite.Level = recite.Level + 1
			if recite.Level > 9 {
				recite.Level = 9
			}
		}
		recite.LevelTime = LevelMap[recite.Level]
		recite.UpdateTime = int(time.Now().Unix())
		if num, err := o.Update(&recite, "Level", "LevelTime", "UpdateTime"); err == nil {
			return num
		}
	}
	return 0
}