package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
)

type OperationLog struct {
}

func (OperationLog) GetList(c *gin.Context) {
	r.Send(c, r.SUCCESS, operationLogService.GetList(util.BindQuery[model.Condition](c)))
}

func (OperationLog) Delete(c *gin.Context) {
	operationLogService.Delete(util.BindJSON[[]int](c))
	r.Send(c, r.SUCCESS)
}
