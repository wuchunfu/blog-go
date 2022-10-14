package dao

import "myblog/model"

type FriendLinkDao struct {
}

func (*FriendLinkDao) CountFriendLink(condition model.Condition) int {
	var count int64
	tx := db.Model(model.FriendLink{})
	if condition.Keywords != "" {
		tx = tx.Where("link_name like ?", "%"+condition.Keywords+"%")
	}
	tx.Count(&count)
	return int(count)
}

func (*FriendLinkDao) BackendFriendLinkList(condition model.Condition) (list []model.FriendLink) {
	offset := (condition.Current - 1) * condition.Size
	tx := db.Table("friend_link")
	if condition.Keywords != "" {
		tx = tx.Where("link_name like ?", "%"+condition.Keywords+"%")
	}
	tx.Limit(condition.Size).Offset(offset).Find(&list)
	return
}

func (*FriendLinkDao) SaveOrUpdate(data model.FriendLink) {
	if data.ID != 0 {
		Updates(&data)
	} else {
		Create(&data)
	}
}
