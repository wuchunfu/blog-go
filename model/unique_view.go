package model

import "time"

type UniqueView struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime" gorm:"autoUpdateTime"`
	ViewsCount int       `json:"viewsCount"`
}

type UniqueViewDTO struct {
	Day        string `json:"day"`
	ViewsCount int    `json:"viewsCount"`
}
