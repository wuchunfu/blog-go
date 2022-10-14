package model

import (
	"reflect"
	"time"
)

type Category struct {
	Universal
	CategoryName string `json:"categoryName"`
}

func (c *Category) IsEmpty() bool {
	return reflect.DeepEqual(c, &Category{})
}

type BackendCategoryDTO struct {
	ID           int       `json:"id"`
	CategoryName string    `json:"categoryName"`
	ArticleCount int       `json:"articleCount"`
	CreateTime   time.Time `json:"createTime"`
}

type CategoryDTO struct {
	ID           int    `json:"id"`
	CategoryName string `json:"categoryName"`
	ArticleCount int    `json:"articleCount"`
}
