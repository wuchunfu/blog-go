package v1

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"myblog/config"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
)

type User struct {
}

func (User) GetUserAreas(c *gin.Context) {
	r.Send(c, r.SUCCESS, userService.GetUserAreas(util.BindQuery[model.Condition](c)))
}

func (User) UpdateInfo(c *gin.Context) {
	userInfo := util.BindJSON[model.UserInfo](c)
	id, _ := c.Get("userInfoId")
	userInfo.ID = id.(int)
	userService.UpdateInfo(userInfo)
	r.Send(c, r.SUCCESS)
}

func (User) UpdateAvatar(c *gin.Context) {
	file, _ := c.FormFile("file")
	avatar := util.FileUtil.UploadFile(file, config.TcConf.UserPath)
	id, ok := c.Get("userInfoId")
	userInfo := model.UserInfo{Avatar: avatar}
	if ok {
		userInfo.ID = id.(int)
	} else {
		session := sessions.Default(c)
		userInfo.ID = session.Get("userInfoId").(int)
	}
	userService.UpdateAvatar(userInfo)
	r.Send(c, r.SUCCESS, avatar)
}

func (User) UpdateAdminPassword(c *gin.Context) {
	id, _ := c.Get("userInfoId")
	r.Send(c, userService.UpdateAdminPassword(util.BindJSON[model.PasswordVO](c), id.(int)))
}

func (User) GetList(c *gin.Context) {
	r.Send(c, r.SUCCESS, userService.GetList(util.BindQuery[model.Condition](c)))
}

func (User) UpdateRole(c *gin.Context) {
	userService.UpdateRole(util.BindJSON[model.UserRoleVO](c))
	r.Send(c, r.SUCCESS)
}

func (User) UpdateDisable(c *gin.Context) {
	userService.UpdateDisable(util.BindJSON[model.UserInfo](c))
	r.Send(c, r.SUCCESS)
}

func (User) GetOnlineList(c *gin.Context) {
	r.Send(c, r.SUCCESS, userService.UserOnlineList())
}

func (User) ForceOffline(c *gin.Context) {
	userService.ForceOffline(util.BindJSON[model.UserDetailDTO](c))
	r.Send(c, r.SUCCESS)
}

func (User) SendCode(c *gin.Context) {
	r.Send(c, userService.SendCode(c.Query("username")))

}

func (User) SaveEmail(c *gin.Context) {
	r.Send(c, userService.SaveEmail(c, util.BindJSON[model.EmailVO](c)))
}

func (User) Register(c *gin.Context) {
	r.Send(c, userService.Register(util.BindJSON[model.UserVO](c)))
}

func (User) UpdatePassword(c *gin.Context) {
	r.Send(c, userService.UpdatePassword(util.BindJSON[model.UserVO](c)))
}
