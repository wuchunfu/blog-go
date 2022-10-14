package service

import (
	"encoding/json"
	"myblog/dao"
	"myblog/model"
	"myblog/util"
)

type PageService struct {
}

func (*PageService) GetList() (list []model.Page) {
	pageList, err := rdb.Get(PageCover).Result()
	if pageList == "" || err != nil {
		list = dao.List([]model.Page{}, "id,page_name,page_label,page_cover", "", "")
		listJson, _ := json.Marshal(list)
		rdb.Set(PageCover, listJson, 0)
	} else {
		util.Json.Unmarshal(pageList, &list)
	}
	return
}

func (*PageService) SaveOrUpdate(data model.Page) {
	pageDao.SaveOrUpdate(data)
	rdb.Del(PageCover)
}

func (*PageService) Delete(pageId int) {
	dao.Delete(model.Page{}, "id = ?", pageId)
	rdb.Del(PageCover)
}
