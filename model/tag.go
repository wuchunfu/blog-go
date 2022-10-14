package model

import (
	"reflect"
	"time"
)

type Tag struct {
	Universal
	TagName string `json:"tagName"`
}

func (t *Tag) IsEmpty() bool {
	return reflect.DeepEqual(t, &Tag{})
}

// TagDTO 先留着，看看以后怎么将这种自定义非数据库中对应结构进行m2m的查询
type TagDTO struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	TagName string `json:"tagName"`
}

func (TagDTO) TableName() string {
	return "tag"
}

type BackendTagDTO struct {
	ID           int       `json:"id"`
	TagName      string    `json:"tagName"`
	ArticleCount int       `json:"articleCount"`
	CreateTime   time.Time `json:"createTime"`
}
