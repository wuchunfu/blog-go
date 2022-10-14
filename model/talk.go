package model

import (
	"reflect"
	"time"
)

type Talk struct {
	Universal
	UserId  int    `json:"userId"`
	Content string `json:"content"`
	Images  string `json:"images"`
	IsTop   int    `json:"isTop"`
	Status  int    `json:"status"`
}

type BackendTalkDTO struct {
	ID         int       `json:"id"`
	Nickname   string    `json:"nickname"`
	Avatar     string    `json:"avatar"`
	Content    string    `json:"content"`
	Images     string    `json:"images"`
	ImgList    []string  `json:"imgList"`
	IsTop      int       `json:"isTop"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
}

type TalkDTO struct {
	ID           int       `json:"id"`
	Nickname     string    `json:"nickname"`
	Avatar       string    `json:"avatar"`
	Content      string    `json:"content"`
	Images       string    `json:"images"`
	ImgList      []string  `json:"imgList" gorm:"-"`
	IsTop        int       `json:"isTop"`
	LikeCount    int       `json:"likeCount"`
	CommentCount int       `json:"commentCount"`
	CreateTime   time.Time `json:"createTime"`
}

func (td *TalkDTO) IsEmpty() bool {
	return reflect.DeepEqual(td, &TalkDTO{})
}
