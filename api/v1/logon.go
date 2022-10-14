package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/util/r"
)

type Logon struct {
}

func (Logon) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	userInfoDTO, code := userService.Login(c, username, password)
	r.Send(c, code, userInfoDTO)
}

func (Logon) Logout(c *gin.Context) {
	r.Send(c, r.SUCCESS)
}
