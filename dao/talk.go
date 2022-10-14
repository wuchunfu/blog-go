package dao

import (
	"myblog/model"
)

type TalkDao struct {
}

func (*TalkDao) CountBackendTalkList(condition model.Condition) int {
	var count int64
	tx := db.Model(model.Talk{})
	if condition.Status != 0 {
		tx.Where("status = ?", condition.Status)
	}
	tx.Count(&count)
	return int(count)
}

func (*TalkDao) BackendTalkList(condition model.Condition) (list []model.BackendTalkDTO) {
	offset := (condition.Current - 1) * condition.Size
	tx := db.Select("t.id,nickname,avatar,content,images,t.is_top,t.status,t.create_time").
		Table("talk t").Joins("JOIN user_info ui ON t.user_id = ui.id")
	if condition.Status != 0 {
		tx.Where("status = ?", condition.Status)
	}
	tx.Order("t.is_top,t.id desc").
		Limit(condition.Size).Offset(offset).Find(&list)
	return
}

func (*TalkDao) GetBackendTalkById(talkId int) (talk model.BackendTalkDTO) {
	db.Select("t.id,nickname,avatar,content,images,t.create_time").
		Table("talk t").Joins("JOIN user_info ui ON t.user_id = ui.id").
		Where("t.id = ?", talkId).Find(&talk)
	return
}

func (*TalkDao) SaveOrUpdate(talk model.Talk) {
	if talk.ID != 0 {
		Updates(&talk, "content", "images", "is_top", "status")
	} else {
		Create(&talk)
	}
}

func (*TalkDao) GetTalkList(condition model.Condition) (list []model.TalkDTO) {
	offset := (condition.Current - 1) * condition.Size
	db.Select("t.id,nickname,avatar,content,images,t.is_top,t.create_time").
		Table("talk t").Joins("JOIN user_info ui ON t.user_id = ui.id").
		Where("t.status = 1").Order("t.is_top desc,t.id desc").
		Limit(condition.Size).Offset(offset).Find(&list)
	return
}

func (*TalkDao) GetTalkById(id string) (talk model.TalkDTO) {
	db.Select("t.id,nickname,avatar,content,images,t.create_time").
		Table("talk t").Joins("JOIN user_info ui ON t.user_id = ui.id").
		Where("t.id = ? AND t.status = 1", id).Find(&talk)
	return
}
