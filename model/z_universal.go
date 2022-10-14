package model

import (
	"time"
)

type Universal struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	CreateTime time.Time `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"updateTime" gorm:"autoUpdateTime"`
}

type Condition struct { // 绑定query用form tag
	Current    int       `form:"current"`
	Size       int       `form:"size"`
	Keywords   string    `form:"keywords"`
	CategoryId int       `form:"categoryId"`
	TagId      int       `form:"tagId"`
	AlbumId    int       `form:"albumId"`
	LoginType  int       `form:"loginType"`
	Type       int       `form:"type"`
	Status     int       `form:"status"`
	StartTime  time.Time `form:"startTime"`
	EndTime    time.Time `form:"endTime"`
	IsDelete   int       `form:"isDelete"`
	IsReview   int       `form:"isReview"`
}

type PageResult[T any] struct {
	Count      int `json:"count"`
	RecordList T   `json:"recordList"`
}

type About struct {
	AboutContent string `json:"aboutContent"`
}

type BlogHomeInfoDTO struct {
	ArticleCount  int             `json:"articleCount"`
	CategoryCount int             `json:"categoryCount"`
	TagCount      int             `json:"tagCount"`
	ViewsCount    int             `json:"viewsCount"`
	WebsiteConfig WebsiteConfigVO `json:"websiteConfig"`
	PageList      []Page          `json:"pageList"`
}

type BlogBackendInfoDTO struct {
	ViewsCount   int `json:"viewsCount"`
	MessageCount int `json:"messageCount"`
	UserCount    int `json:"userCount"`
	ArticleCount int `json:"articleCount"`
}

type BlogBackInfoDTO struct {
	ViewsCount            int                    `json:"viewsCount"`
	MessageCount          int                    `json:"messageCount"`
	UserCount             int                    `json:"userCount"`
	ArticleCount          int                    `json:"articleCount"`
	CategoryDTOList       []CategoryDTO          `json:"categoryDTOList"`
	TagDTOList            []TagDTO               `json:"tagDTOList"`
	ArticleStatisticsList []ArticleStatisticsDTO `json:"articleStatisticsList"`
	UniqueViewDTOList     []UniqueViewDTO        `json:"uniqueViewDTOList"`
	ArticleRankDTOList    []ArticleRankDTO       `json:"articleRankDTOList"`
}

type DeleteVO struct {
	IdList   []int `json:"idList"`
	IsDelete int   `json:"isDelete"`
}

type LabelOptionDTO struct {
	ID       int              `json:"id"`
	Label    string           `json:"label"`
	Children []LabelOptionDTO `json:"children"`
}

type EmailVO struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
