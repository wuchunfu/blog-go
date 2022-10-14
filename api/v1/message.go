package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
	"strings"
)

type Message struct {
}

func (Message) GetBackendList(c *gin.Context) {
	isReview := false
	if index := strings.Index(c.Request.RequestURI, "isReview"); index != -1 {
		isReview = true
	}
	r.Send(c, r.SUCCESS, messageService.GetBackendList(util.BindQuery[model.Condition](c), isReview))
}

func (Message) UpdateReview(c *gin.Context) {
	messageService.UpdateReview(util.BindJSON[model.ReviewVO](c))
	r.Send(c, r.SUCCESS)
}

func (Message) Delete(c *gin.Context) {
	messageService.Delete(util.BindJSON[[]int](c))
	r.Send(c, r.SUCCESS)
}

func (Message) GetList(c *gin.Context) {
	r.Send(c, r.SUCCESS, messageService.GetList())
}

func (Message) SaveMessage(c *gin.Context) {
	messageService.SaveMessage(c, util.BindJSON[model.MessageVO](c))
	r.Send(c, r.SUCCESS)
}
