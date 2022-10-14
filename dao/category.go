package dao

import (
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
)

type CategoryDao struct {
}

func (*CategoryDao) GetCategoryListBySearch(condition model.Condition) (list []model.Category) {
	whereMap := make(map[string]any)
	if condition.Keywords != "" {
		whereMap["category_name like"] = "%" + condition.Keywords + "%"
	}
	query, args := util.GetCondition(whereMap)
	return List(list, "*", "id desc", query, args)
}

func (*CategoryDao) CountBackendCategory(condition model.Condition) int {
	var count int64
	tx := db.Model(model.Category{})
	if condition.Keywords != "" {
		tx = tx.Where("category_name like ?", "%"+condition.Keywords+"%")
	}
	tx.Count(&count)
	return int(count)
}

func (*CategoryDao) BackendCategoryList(condition model.Condition) (list []model.BackendCategoryDTO) {
	offset := (condition.Current - 1) * condition.Size
	tx := db.Table("category c").
		Select("c.id", "c.category_name", "COUNT(a.id) AS article_count", "c.create_time").
		Joins("LEFT JOIN article a ON c.id = a.category_id")
	if condition.Keywords != "" {
		tx = tx.Where("category_name like ?", "%"+condition.Keywords+"%")
	}
	tx.Group("c.id").Order("c.id desc").
		Limit(condition.Size).Offset(offset).Find(&list)
	return
}

func (*CategoryDao) SaveOrUpdate(category model.Category) int {
	existCategory := GetOne(model.Category{}, "category_name = ?", category.CategoryName)
	if !existCategory.IsEmpty() && existCategory.ID != category.ID { // 同名存在 且 存在的id不等于当前要更新的id -> 重复
		return r.CategoryExist
	}
	if category.ID != 0 { // 存在 更新
		Updates(&category, "category_name")
	} else { // 不存在 插入
		db.Create(&category)
	}
	return r.SUCCESS
}

func (*CategoryDao) GetCategoryDTOList() (list []model.CategoryDTO) {
	db.Select("c.id,c.category_name,COUNT(a.id) AS article_count").
		Table("category c").
		Joins("LEFT JOIN (" +
			" SELECT id, category_id FROM article WHERE is_delete = 0 AND status = 1 " +
			") a ON c.id = a.category_id").
		Group("c.id").Find(&list)
	return
}
