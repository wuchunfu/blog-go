package model

type ArticleTag struct {
	ID        int `json:"id" gorm:"primaryKey"`
	ArticleId int `json:"articleId" gorm:"primaryKey"`
	TagId     int `json:"tagId" gorm:"primaryKey"`
}
