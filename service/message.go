package service

import (
	"github.com/gin-gonic/gin"
	"myblog/dao"
	"myblog/model"
	"myblog/util"
)

type MessageService struct {
}

func (*MessageService) GetBackendList(condition model.Condition, isReview bool) model.PageResult[[]model.Message] {
	count := dao.Count(model.Message{}, "")
	messageList := messageDao.BackendMessageList(condition, isReview)
	return model.PageResult[[]model.Message]{
		Count:      count,
		RecordList: messageList,
	}
}

func (*MessageService) UpdateReview(data model.ReviewVO) {
	messageDao.UpdateMessagesReview(data)
}

func (*MessageService) Delete(ids []int) {
	dao.Delete(model.Message{}, "id in ?", ids)
}

func (*MessageService) GetList() []model.Message {
	return dao.List([]model.Message{}, "*", "", "is_review = 1")
}

func (*MessageService) SaveMessage(c *gin.Context, data model.MessageVO) {
	ipAddress := util.IpUtil.GetIpAddress(c)
	ipSource := util.IpUtil.GetIpSource(ipAddress)
	message := model.Message{
		Nickname:       data.Nickname,
		Avatar:         data.Avatar,
		MessageContent: data.MessageContent,
		IpAddress:      ipAddress,
		IpSource:       ipSource,
		Time:           data.Time,
	}
	if blogInfoService.GetWebsiteConfig().IsMessageReview == 0 {
		message.IsReview = 1
	} else {
		message.IsReview = 0
	}
	dao.Create(&message)
}
