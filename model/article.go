package model

import (
	"reflect"
	"time"
)

const (
	PUBLIC = iota + 1
	SECRET
	DRAFT
)

type Article struct {
	Universal
	UserId         int    `json:"userId"`
	CategoryId     int    `json:"categoryId"`
	ArticleCover   string `json:"articleCover"`
	ArticleTitle   string `json:"articleTitle"`
	ArticleContent string `json:"articleContent"`
	Type           int    `json:"type"`
	OriginalUrl    string `json:"originalUrl"`
	IsTop          int    `json:"isTop"`
	IsDelete       int    `json:"isDelete"`
	Status         int    `json:"status"`
}

func (a *Article) IsEmpty() bool {
	return reflect.DeepEqual(a, &Article{})
}

type ArticleBackendDTO struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	ArticleCover string    `json:"articleCover"`
	ArticleTitle string    `json:"articleTitle"`
	CreateTime   time.Time `json:"createTime"`
	LikeCount    int       `json:"likeCount"`
	ViewsCount   int       `json:"viewsCount"`
	CategoryName string    `json:"categoryName"`
	Type         int       `json:"type"`
	IsTop        int       `json:"isTop"`
	IsDelete     int       `json:"isDelete"`
	Status       int       `json:"status"`
	// foreignKey 本表ID joinForeignKey 中间表关联字段 reference tag 表ID joinReference  中间表 TagId （在这里可以省略掉除joinForeignKey以外的）
	// 暂时还不知道使用自定义结构体如何进行m2m,先这样后面继续看看怎么做
	Tags []Tag `json:"tagDTOList" gorm:"many2many:article_tag;foreignKey:ID;joinForeignKey:ArticleId;reference:ID;joinReference:TagId;"`
}

type ArticleVO struct {
	ID             int      `json:"id"`
	ArticleTitle   string   `json:"articleTitle"`
	ArticleContent string   `json:"articleContent"`
	ArticleCover   string   `json:"articleCover"`
	CategoryName   string   `json:"categoryName"`
	TagNameList    []string `json:"tagNameList"`
	Type           int      `json:"type"`
	OriginalUrl    string   `json:"originalUrl"`
	IsTop          int      `json:"isTop"`
	Status         int      `json:"status"`
}

type ArticleHomeDTO struct {
	ID             int       `json:"id"`
	ArticleCover   string    `json:"articleCover"`
	ArticleTitle   string    `json:"articleTitle"`
	ArticleContent string    `json:"articleContent"`
	CreateTime     time.Time `json:"createTime"`
	IsTop          int       `json:"isTop"`
	Type           int       `json:"type"`
	CategoryId     int       `json:"categoryId"`
	CategoryName   string    `json:"categoryName"`
	Tags           []Tag     `json:"tagDTOList" gorm:"many2many:article_tag;foreignKey:ID;joinForeignKey:ArticleId;reference:ID;joinReference:TagId;"`
}

type ArticleDTO struct {
	ID                   int                   `json:"id"`
	ArticleCover         string                `json:"articleCover"`
	ArticleTitle         string                `json:"articleTitle"`
	ArticleContent       string                `json:"articleContent"`
	LinkCount            int                   `json:"linkCount" gorm:"-"`
	ViewsCount           int                   `json:"viewsCount" gorm:"-"`
	Type                 int                   `json:"type"`
	OriginalUrl          string                `json:"originalUrl"`
	CreateTime           time.Time             `json:"createTime"`
	UpdateTime           time.Time             `json:"updateTime"`
	CategoryId           int                   `json:"categoryId"`
	CategoryName         string                `json:"categoryName"`
	Tags                 []Tag                 `json:"tagDTOList" gorm:"many2many:article_tag;foreignKey:ID;joinForeignKey:ArticleId;reference:ID;joinReference:TagId;"`
	LastArticle          ArticlePaginationDTO  `json:"lastArticle" gorm:"-"`
	NextArticle          ArticlePaginationDTO  `json:"nextArticle" gorm:"-"`
	RecommendArticleList []RecommendArticleDTO `json:"recommendArticleList" gorm:"-"`
	NewestArticleList    []RecommendArticleDTO `json:"newestArticleList" gorm:"-"`
}

type ArticlePaginationDTO struct {
	ID           int    `json:"id"`
	ArticleCover string `json:"articleCover"`
	ArticleTitle string `json:"articleTitle"`
}

type RecommendArticleDTO struct {
	ID           int       `json:"id"`
	ArticleCover string    `json:"articleCover"`
	ArticleTitle string    `json:"articleTitle"`
	CreateTime   time.Time `json:"createTime"`
}

type ArticleSearchDTO struct {
	ID             int    `json:"id"`
	ArticleTitle   string `json:"articleTitle"`
	ArticleContent string `json:"articleContent"`
	IsDelete       int    `json:"isDelete"`
	Status         int    `json:"status"`
}

type ArchiveDTO struct {
	ID           int       `json:"id"`
	ArticleTitle string    `json:"articleTitle"`
	CreateTime   time.Time `json:"createTime"`
}

type ArticleStatisticsDTO struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type ArticleRankDTO struct {
	ArticleTitle string `json:"articleTitle"`
	ViewsCount   int    `json:"viewsCount"`
}

type ArticlePreviewDTO struct {
	ID           int       `json:"id"`
	ArticleCover string    `json:"articleCover"`
	ArticleTitle string    `json:"articleTitle"`
	CreateTime   time.Time `json:"createTime"`
	CategoryId   int       `json:"categoryId"`
	CategoryName string    `json:"categoryName"`
	Tags         []Tag     `json:"tagDTOList" gorm:"many2many:article_tag;foreignKey:ID;joinForeignKey:ArticleId;"`
}

type ArticlePreviewListDTO struct {
	Name                  string              `json:"name"`
	ArticlePreviewDTOList []ArticlePreviewDTO `json:"articlePreviewDTOList"`
}
