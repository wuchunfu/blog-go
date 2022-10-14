package dao

import (
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
)

type TagDao struct {
}

func (*TagDao) SaveTag(tag *model.Tag) {
	db.Create(&tag)
}

func (*TagDao) GetTagListBySearch(condition model.Condition) (list []model.Tag) {
	whereMap := make(map[string]any)
	if condition.Keywords != "" {
		whereMap["tag_name like"] = "%" + condition.Keywords + "%"
	}
	query, args := util.GetCondition(whereMap)
	return List(list, "*", "id desc", query, args)
}

func (*TagDao) TagNameListByArticleId(articleId int) []string {
	var list []string
	db.Table("tag t").Select("tag_name").
		Joins("JOIN article_tag at ON t.id = at.tag_id").
		Where("article_id", articleId).Find(&list)
	return list
}

func (*TagDao) CountBackendTag(condition model.Condition) int {
	var count int64
	tx := db.Model(&model.Tag{})
	if condition.Keywords != "" {
		tx = tx.Where("tag_name like ?", "%"+condition.Keywords+"%")
	}
	tx.Count(&count)
	return int(count)
}

func (*TagDao) BackendTagList(condition model.Condition) (list []model.BackendTagDTO) {
	offset := (condition.Current - 1) * condition.Size
	tx := db.Table("tag t").
		Select("t.id,tag_name,COUNT(at.article_id) AS article_count,t.create_time").
		Joins("LEFT JOIN article_tag at ON t.id = at.tag_id")
	if condition.Keywords != "" {
		tx = tx.Where("tag_name like ?", "%"+condition.Keywords+"%")
	}
	tx.Group("t.id").Order("t.id desc").
		Limit(condition.Size).Offset(offset).Find(&list)
	return
}

func (*TagDao) SaveOrUpdate(tag model.Tag) int {
	existTag := GetOne(model.Tag{}, "tag_name = ?", tag.TagName)
	if !existTag.IsEmpty() && existTag.ID != tag.ID {
		return r.TagExist
	}
	if tag.ID != 0 {
		Updates(&tag, "tag_name")
	} else {
		db.Create(&tag)
	}
	return r.SUCCESS
}

func (*TagDao) TagDTOList() (list []model.TagDTO) {
	db.Select("id,tag_name").Table("tag").Find(&list)
	return
}
