package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
	"strconv"
)

type Menu struct {
}

func (Menu) GetUserMenuList(c *gin.Context) {
	userInfoId, _ := c.Get("userInfoId")
	r.Send(c, r.SUCCESS, menuService.GetUserMenuList(userInfoId.(int)))
}
func (Menu) GetList(c *gin.Context) {
	r.Send(c, r.SUCCESS, menuService.GetList(util.BindQuery[model.Condition](c)))
}

func (Menu) SaveOrUpdate(c *gin.Context) {
	menuService.SaveOrUpdate(util.BindJSON[model.Menu](c))
	r.Send(c, r.SUCCESS)
}

func (Menu) Delete(c *gin.Context) {
	menuId, _ := strconv.Atoi(c.Param("menuId"))
	menuService.Delete(menuId)
	r.Send(c, r.SUCCESS)
}

func (Menu) OptionList(c *gin.Context) {
	r.Send(c, r.SUCCESS, menuService.OptionList())
}
