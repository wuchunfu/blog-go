package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
)

type Category struct {
}

// GetListBySearch 根据分类名模糊查询分类
func (Category) GetListBySearch(c *gin.Context) {
	r.Send(c, r.SUCCESS, categoryService.GetListBySearch(util.BindQuery[model.Condition](c)))
}

// GetBackendList 获取后台分类列表
func (Category) GetBackendList(c *gin.Context) {
	r.Send(c, r.SUCCESS, categoryService.GetBackendList(util.BindQuery[model.Condition](c)))
}

// SaveOrUpdate 新增或更新分类
func (Category) SaveOrUpdate(c *gin.Context) {
	r.Send(c, categoryService.SaveOrUpdate(util.BindJSON[model.Category](c)))
}

// Delete 删除分类
func (Category) Delete(c *gin.Context) {
	r.Send(c, categoryService.Delete(util.BindJSON[[]int](c)))
}

func (Category) GetList(c *gin.Context) {
	r.Send(c, r.SUCCESS, categoryService.GetList())
}
