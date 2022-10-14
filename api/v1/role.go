package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
)

type Role struct {
}

func (Role) GetUserRoleList(c *gin.Context) {
	r.Send(c, r.SUCCESS, roleService.GetUserRoleList())
}

func (Role) GetList(c *gin.Context) {
	r.Send(c, r.SUCCESS, roleService.GetList(util.BindQuery[model.Condition](c)))
}

func (Role) SaveOrUpdate(c *gin.Context) {
	r.Send(c, roleService.SaveOrUpdate(util.BindJSON[model.RoleVO](c)))
}

func (Role) Delete(c *gin.Context) {
	roleService.Delete(util.BindJSON[[]int](c))
	r.Send(c, r.SUCCESS)
}
