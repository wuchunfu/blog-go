package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
	"strconv"
)

type Page struct {
}

func (Page) GetList(c *gin.Context) {
	r.Send(c, r.SUCCESS, pageService.GetList())
}

func (Page) SaveOrUpdate(c *gin.Context) {
	pageService.SaveOrUpdate(util.BindJSON[model.Page](c))
	r.Send(c, r.SUCCESS)
}

func (Page) Delete(c *gin.Context) {
	pageId, _ := strconv.Atoi(c.Param("pageId"))
	pageService.Delete(pageId)
	r.Send(c, r.SUCCESS)
}
