package dao

import (
	"myblog/model"
)

type OperationLogDao struct {
}

func (*OperationLogDao) GetList(condition model.Condition) (list []model.OperationLog) {
	offset := (condition.Current - 1) * condition.Size
	tx := db.Table("operation_log")
	if condition.Keywords != "" {
		tx = tx.Where("opt_module like ?", "%"+condition.Keywords+"%").
			Or("opt_desc like ?", "%"+condition.Keywords+"%")
	}
	tx.Order("id desc").Limit(condition.Size).Offset(offset).Find(&list)
	return
}
