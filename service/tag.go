package service

import (
	"myblog/dao"
	"myblog/model"
	"myblog/util/r"
)

type TagService struct {
}

func (*TagService) GetListBySearch(condition model.Condition) []model.Tag {
	return tagDao.GetTagListBySearch(condition)
}

func (*TagService) GetBackendList(condition model.Condition) model.PageResult[[]model.BackendTagDTO] {
	count := tagDao.CountBackendTag(condition)
	if count == 0 {
		return model.PageResult[[]model.BackendTagDTO]{}
	}
	tagList := tagDao.BackendTagList(condition)
	return model.PageResult[[]model.BackendTagDTO]{Count: count, RecordList: tagList}
}

func (*TagService) SaveOrUpdate(tag model.Tag) int {
	return tagDao.SaveOrUpdate(tag)
}

func (*TagService) Delete(ids []int) int {
	count := dao.Count(model.ArticleTag{}, "tag_id in ?", ids)
	if count > 0 {
		return r.TagArticleExist
	}
	dao.Delete(model.Tag{}, "id in ?", ids)
	return r.SUCCESS
}

func (*TagService) GetList() model.PageResult[[]model.Tag] {
	count := dao.Count(model.Tag{}, "")
	if count == 0 {
		return model.PageResult[[]model.Tag]{count, []model.Tag{}}
	}
	tagList := dao.List([]model.Tag{}, "*", "", "", "")
	return model.PageResult[[]model.Tag]{count, tagList}
}
