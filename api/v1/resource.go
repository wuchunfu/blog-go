package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
	"strconv"
)

type Resource struct {
}

func (Resource) GetList(c *gin.Context) {
	r.Send(c, r.SUCCESS, resourceService.GetList(util.BindQuery[model.Condition](c)))
}

func (Resource) SaveOrUpdate(c *gin.Context) {
	resourceService.SaveOrUpdate(util.BindJSON[model.Resource](c))
	r.Send(c, r.SUCCESS)
}

func (Resource) Delete(c *gin.Context) {
	resourceId, _ := strconv.Atoi(c.Param("resourceId"))
	resourceService.Delete(resourceId)
	r.Send(c, r.SUCCESS)
}

func (Resource) OptionList(c *gin.Context) {
	r.Send(c, r.SUCCESS, resourceService.OptionList())
}
