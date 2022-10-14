package service

import (
	"myblog/dao"
	"myblog/model"
	"myblog/util/r"
)

type CategoryService struct {
}

func (*CategoryService) GetListBySearch(condition model.Condition) []model.Category {
	return categoryDao.GetCategoryListBySearch(condition)
}

func (*CategoryService) GetBackendList(condition model.Condition) model.PageResult[[]model.BackendCategoryDTO] {
	count := categoryDao.CountBackendCategory(condition)
	if count == 0 {
		return model.PageResult[[]model.BackendCategoryDTO]{}
	}
	categoryList := categoryDao.BackendCategoryList(condition)
	return model.PageResult[[]model.BackendCategoryDTO]{Count: count, RecordList: categoryList}
}

func (*CategoryService) SaveOrUpdate(category model.Category) int {
	return categoryDao.SaveOrUpdate(category)
}

func (*CategoryService) Delete(ids []int) int {
	count := dao.Count(model.Category{}, "category_id in ?", ids)
	if count > 0 {
		return r.CategoryArticleExist
	}
	dao.Delete(model.Category{}, "id in ?", ids)
	return r.SUCCESS
}

func (*CategoryService) GetList() model.PageResult[[]model.CategoryDTO] {
	count := dao.Count(model.Category{}, "")
	if count == 0 {
		return model.PageResult[[]model.CategoryDTO]{}
	}
	categoryList := categoryDao.GetCategoryDTOList()
	return model.PageResult[[]model.CategoryDTO]{count, categoryList}
}
