package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
)

type FriendLink struct {
}

func (FriendLink) GetBackendList(c *gin.Context) {
	r.Send(c, r.SUCCESS, friendLinkService.GetBackendList(util.BindQuery[model.Condition](c)))
}

func (FriendLink) SaveOrUpdate(c *gin.Context) {
	friendLinkService.SaveOrUpdate(util.BindJSON[model.FriendLink](c))
	r.Send(c, r.SUCCESS)
}

func (FriendLink) Delete(c *gin.Context) {
	friendLinkService.Delete(util.BindJSON[[]int](c))
	r.Send(c, r.SUCCESS)
}

func (FriendLink) GetList(c *gin.Context) {
	r.Send(c, r.SUCCESS, friendLinkService.GetList())
}
