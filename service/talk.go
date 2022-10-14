package service

import (
	"github.com/gin-gonic/gin"
	"myblog/dao"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
	"strconv"
)

type TalkService struct {
}

func (*TalkService) GetBackendList(condition model.Condition) model.PageResult[[]model.BackendTalkDTO] {
	count := talkDao.CountBackendTalkList(condition)
	if count == 0 {
		return model.PageResult[[]model.BackendTalkDTO]{}
	}
	talkList := talkDao.BackendTalkList(condition)
	for i, talk := range talkList {
		var imgList []string
		util.Json.Unmarshal(talk.Images, &imgList)
		talkList[i].ImgList = imgList
	}
	return model.PageResult[[]model.BackendTalkDTO]{
		Count:      count,
		RecordList: talkList,
	}
}

func (*TalkService) GetBackendById(talkId int) (talk model.BackendTalkDTO) {
	talk = talkDao.GetBackendTalkById(talkId)
	var imgList []string
	util.Json.Unmarshal(talk.Images, &imgList)
	talk.ImgList = imgList
	return
}

func (*TalkService) SaveOrUpdate(c *gin.Context, data model.BackendTalkDTO) {
	userInfoId, _ := c.Get("userInfoId")
	talk := model.Talk{
		Universal: model.Universal{ID: data.ID},
		UserId:    userInfoId.(int),
		Content:   data.Content,
		Images:    data.Images,
		IsTop:     data.IsTop,
		Status:    data.Status,
	}
	talkDao.SaveOrUpdate(talk)
}

func (*TalkService) Delete(talkIds []string) {
	id, _ := strconv.Atoi(talkIds[0])
	dao.Delete(model.Talk{}, "id = ?", id)
}

func (*TalkService) GetHomeList() (list []string) {
	talkList := dao.List([]model.Talk{}, "id,content,images,is_top,status",
		"is_top desc, id desc", "status = 1")
	for _, talk := range talkList {
		list = append(list, talk.Content)
	}
	return
}

func (*TalkService) GetList(condition model.Condition) model.PageResult[[]model.TalkDTO] {
	count := dao.Count(model.Talk{}, "status = 1")
	if count == 0 {
		return model.PageResult[[]model.TalkDTO]{}
	}
	talkList := talkDao.GetTalkList(condition)
	// 查询点赞数量和评论数量
	var ids []int
	for _, talk := range talkList {
		ids = append(ids, talk.ID)
	}
	commentCountList := commentDao.ListCountByTopicIds(ids, 3)
	commentCountMap := make(map[int]int)
	for _, commentCount := range commentCountList {
		commentCountMap[commentCount.ID] = commentCount.CommentCount
	}
	likeCountMap := rdb.HGetAll(TalkLikeCount).Val()
	for i, talk := range talkList {
		var imgList []string
		util.Json.Unmarshal(talk.Images, &imgList)
		talkList[i].ImgList = imgList
		talkList[i].CommentCount = commentCountMap[talk.ID]
		talkList[i].LikeCount, _ = strconv.Atoi(likeCountMap[strconv.Itoa(talk.ID)])
	}
	return model.PageResult[[]model.TalkDTO]{count, talkList}
}

func (*TalkService) GetById(talkId string) (int, model.TalkDTO) {
	talk := talkDao.GetTalkById(talkId)
	if talk.IsEmpty() {
		return r.TalkNotExist, talk
	}
	// 说说点赞数量
	talk.LikeCount, _ = rdb.HGet(TalkLikeCount, talkId).Int()
	var imgList []string
	util.Json.Unmarshal(talk.Images, &imgList)
	talk.ImgList = imgList
	return r.SUCCESS, talk
}

func (*TalkService) SaveLike(uid int, talkId int) {
	talkLikeKey := TalkUserLike + strconv.Itoa(uid)
	if rdb.SIsMember(talkLikeKey, talkId).Val() {
		rdb.SRem(talkLikeKey, talkId)
		rdb.HIncrBy(TalkLikeCount, strconv.Itoa(talkId), -1)
	} else {
		rdb.SAdd(talkLikeKey, talkId)
		rdb.HIncrBy(TalkLikeCount, strconv.Itoa(talkId), 1)
	}
}
