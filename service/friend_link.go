package service

import (
	"myblog/dao"
	"myblog/model"
)

type FriendLinkService struct {
}

func (*FriendLinkService) GetBackendList(condition model.Condition) model.PageResult[[]model.FriendLink] {
	count := friendLinkDao.CountFriendLink(condition)
	if count == 0 {
		return model.PageResult[[]model.FriendLink]{}
	}
	friendLink := friendLinkDao.BackendFriendLinkList(condition)
	return model.PageResult[[]model.FriendLink]{
		Count:      count,
		RecordList: friendLink,
	}
}

func (*FriendLinkService) SaveOrUpdate(data model.FriendLink) {
	friendLinkDao.SaveOrUpdate(data)
}

func (*FriendLinkService) Delete(data []int) {
	dao.Delete(model.FriendLink{}, "id in ?", data)
}

func (*FriendLinkService) GetList() []model.FriendLink {
	return dao.List([]model.FriendLink{}, "*", "", "")
}
