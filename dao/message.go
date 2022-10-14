package dao

import (
	"myblog/model"
)

type MessageDao struct {
}

func (*MessageDao) BackendMessageList(condition model.Condition, isReview bool) (list []model.Message) {
	offset := (condition.Current - 1) * condition.Size
	tx := db.Table("message")
	if condition.Keywords != "" {
		tx = tx.Where("nickname like ?", "%"+condition.Keywords+"%")
	}
	if isReview {
		tx = tx.Where("is_review = ?", condition.IsReview)
	}
	tx.Order("id desc").Limit(condition.Size).Offset(offset).Find(&list)
	return
}

func (*MessageDao) UpdateMessagesReview(data model.ReviewVO) {
	db.Select("is_review").Where("id in ?", data.IdList).Updates(model.Message{IsReview: data.IsReview})
}
