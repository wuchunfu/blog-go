package service

import (
	"myblog/dao"
	"myblog/model"
)

type OperationLogService struct {
}

func (*OperationLogService) GetList(condition model.Condition) model.PageResult[[]model.OperationLog] {
	count := dao.Count(model.OperationLog{}, "")
	if count == 0 {
		return model.PageResult[[]model.OperationLog]{}
	}
	operationLogList := operationLogDao.GetList(condition)
	return model.PageResult[[]model.OperationLog]{count, operationLogList}
}

func (*OperationLogService) Delete(ids []int) {
	dao.Delete(model.OperationLog{}, "id in ?", ids)
}
