package dao

import (
	"myblog/model"
	"time"
)

type UniqueViewDao struct {
}

func (*UniqueViewDao) UniqueViewList(startTime time.Time, endTime time.Time) (list []model.UniqueViewDTO) {
	db.Select("DATE_FORMAT( create_time, \"%Y-%m-%d\" ) AS day,views_count").
		Table("unique_view").
		Where("create_time > ? AND create_time <= ?", startTime, endTime).
		Order("create_time").Find(&list)
	return
}
