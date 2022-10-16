package service

import (
	"myblog/dao"
	"myblog/model"
	"time"
)

type UniqueViewService struct {
}

func (*UniqueViewService) UniqueViewList() []model.UniqueViewDTO {
	startTime := time.Now().AddDate(0, 0, -7)
	endTime := time.Now()
	return uniqueViewDao.UniqueViewList(startTime, endTime)
}

// 定时任务，统计每天用户量保存到mysql
func (*UniqueViewService) saveUniqueView() {
	count := rdb.SCard(UniqueVisitor).Val()
	dao.Create(&model.UniqueView{
		ViewsCount: int(count),
		CreateTime: time.Now().AddDate(0, 0, -1), // 主动保存时间
	})
}

// 定时任务，清空每天用户量和游客分布区域
func (*UniqueViewService) clear() {
	rdb.Del(UniqueVisitor)
	rdb.Del(VisitorArea)
}
