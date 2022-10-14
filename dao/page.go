package dao

import (
	"myblog/model"
)

type PageDao struct {
}

func (*PageDao) SaveOrUpdate(data model.Page) {
	if data.ID != 0 {
		Updates(&data)
	} else {
		Create(&data)
	}
}
