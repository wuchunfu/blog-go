package middleware

import (
	"github.com/gin-gonic/gin"
	"myblog/dao"
	"myblog/util/r"
)

func RBAC() gin.HandlerFunc {
	e := dao.ReturnEnforcer()
	e.LoadPolicy()
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		//log.Println("==================> 权限管理:", role, c.FullPath(), c.Request.Method)
		if has, err := e.Enforce(role, c.FullPath()[4:], c.Request.Method); err != nil || !has {
			r.Send(c, r.PermissionDenied)
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
