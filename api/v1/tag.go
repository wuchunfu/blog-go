package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
)

type Tag struct {
}

func (Tag) GetListBySearch(c *gin.Context) {
	r.Send(c, r.SUCCESS, tagService.GetListBySearch(util.BindQuery[model.Condition](c)))
}

func (Tag) GetBackendList(c *gin.Context) {
	r.Send(c, r.SUCCESS, tagService.GetBackendList(util.BindQuery[model.Condition](c)))
}

func (Tag) SaveOrUpdate(c *gin.Context) {
	r.Send(c, tagService.SaveOrUpdate(util.BindJSON[model.Tag](c)))
}

func (Tag) Delete(c *gin.Context) {
	r.Send(c, tagService.Delete(util.BindJSON[[]int](c)))
}

func (Tag) GetList(c *gin.Context) {
	r.Send(c, r.SUCCESS, tagService.GetList())
}
