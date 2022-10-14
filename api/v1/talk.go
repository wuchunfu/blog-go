package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/config"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
	"strconv"
)

type Talk struct {
}

func (Talk) GetBackendList(c *gin.Context) {
	r.Send(c, r.SUCCESS, talkService.GetBackendList(util.BindQuery[model.Condition](c)))
}

func (Talk) GetBackendById(c *gin.Context) {
	talkId, _ := strconv.Atoi(c.Param("talkId"))
	r.Send(c, r.SUCCESS, talkService.GetBackendById(talkId))
}

func (Talk) SaveImages(c *gin.Context) {
	file, _ := c.FormFile("file")
	r.Send(c, r.SUCCESS, util.FileUtil.UploadFile(file, config.TcConf.TalkPath))
}

func (Talk) SaveOrUpdate(c *gin.Context) {
	talkService.SaveOrUpdate(c, util.BindJSON[model.BackendTalkDTO](c))
	r.Send(c, r.SUCCESS)
}

func (Talk) Delete(c *gin.Context) {
	talkService.Delete(util.BindJSON[[]string](c))
	r.Send(c, r.SUCCESS)
}

func (Talk) GetHomeList(c *gin.Context) {
	r.Send(c, r.SUCCESS, talkService.GetHomeList())
}

func (Talk) GetList(c *gin.Context) {
	r.Send(c, r.SUCCESS, talkService.GetList(util.BindQuery[model.Condition](c)))
}

func (Talk) GetById(c *gin.Context) {
	talkId := c.Param("talkId")
	code, talk := talkService.GetById(talkId)
	r.Send(c, code, talk)
}

func (Talk) SaveLike(c *gin.Context) {
	uid, _ := c.Get("userInfoId")
	talkId, _ := strconv.Atoi(c.Param("talkId"))
	talkService.SaveLike(uid.(int), talkId)
	r.Send(c, r.SUCCESS)
}
