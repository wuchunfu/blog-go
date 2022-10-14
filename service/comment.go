package service

import (
	"myblog/dao"
	"myblog/model"
	"myblog/util"
	"strconv"
)

type CommentService struct {
}

func (*CommentService) GetBackendList(condition model.Condition, isReview bool) model.PageResult[[]model.BackendCommentDTO] {
	count := commentDao.CountBackendCommentDTO(condition, isReview)
	if count == 0 {
		return model.PageResult[[]model.BackendCommentDTO]{0, []model.BackendCommentDTO{}}
	}
	backendCommentDTOList := commentDao.BackendCommentDTOList(condition, isReview)
	return model.PageResult[[]model.BackendCommentDTO]{count, backendCommentDTOList}
}

func (*CommentService) UpdateReview(data model.ReviewVO) {
	commentDao.UpdateCommentsReview(data)
}

func (*CommentService) Delete(ids []int) {
	dao.Delete(model.Comment{}, "id in ?", ids)
}

func (*CommentService) GetList(commonVO model.CommentVO) model.PageResult[[]model.CommentDTO] {
	commentCount := commentDao.CountComment(commonVO)
	if commentCount == 0 {
		return model.PageResult[[]model.CommentDTO]{0, []model.CommentDTO{}}
	}
	commentList := commentDao.GetCommentList(commonVO)
	likeCountMap := rdb.HGetAll(CommentLikeCount).Val()
	var commentIds []int
	for _, comment := range commentList {
		commentIds = append(commentIds, comment.ID)
	}
	replyList := commentDao.GetReplyList(commentIds)
	replyMap := make(map[int][]model.ReplyDTO)
	for i, reply := range replyList {
		replyList[i].LikeCount, _ = strconv.Atoi(likeCountMap[strconv.Itoa(reply.ID)])
		replyMap[reply.ParentId] = append(replyMap[reply.ParentId], reply)
	}
	replyCountList := commentDao.GetReplyCountListByCommentId(commentIds)
	replyCountMap := make(map[int]int)
	for _, reply := range replyCountList {
		replyCountMap[reply.CommentId] = reply.ReplyCount
	}
	for i, comment := range commentList {
		commentList[i].LikeCount, _ = strconv.Atoi(likeCountMap[strconv.Itoa(comment.ID)])
		commentList[i].ReplyDTOList = replyMap[comment.ID]
		commentList[i].ReplyCount = replyCountMap[comment.ID]
	}
	return model.PageResult[[]model.CommentDTO]{commentCount, commentList}
}

func (*CommentService) GetReplyListByCommentId(id int, condition model.Condition) []model.ReplyDTO {
	replyList := commentDao.GetReplyListByCommentId(id, condition)
	likeCountMap := rdb.HGetAll(CommentLikeCount).Val()
	for i, reply := range replyList {
		replyList[i].LikeCount, _ = strconv.Atoi(likeCountMap[strconv.Itoa(reply.ID)])
	}
	return replyList
}

func (*CommentService) SaveLike(uid int, commentId int) {
	commentLikeKey := CommentUserLike + strconv.Itoa(uid)
	if rdb.SIsMember(commentLikeKey, commentId).Val() {
		rdb.SRem(commentLikeKey, commentId)
		rdb.HIncrBy(CommentLikeCount, strconv.Itoa(commentId), -1)
	} else {
		rdb.SAdd(commentLikeKey, commentId)
		rdb.HIncrBy(CommentLikeCount, strconv.Itoa(commentId), 1)
	}
}

func (*CommentService) SaveComment(uid int, commentVO model.CommentVO) {
	commentVO.CommentContent = util.HTMLUtil.Filter(commentVO.CommentContent)
	comment := model.Comment{
		UserId:         uid,
		TopicId:        commentVO.TopicId,
		CommentContent: commentVO.CommentContent,
		ReplyUserId:    commentVO.ReplyUserId,
		ParentId:       commentVO.ParentId,
		Type:           commentVO.Type,
	}
	websiteConfig := blogInfoService.GetWebsiteConfig()
	if websiteConfig.IsCommentReview == 0 { // 0 就是不用审核，数据就要设置为已审核状态
		comment.IsReview = 1
	} else {
		comment.IsReview = 0
	}
	commentDao.SaveComment(comment)
	// [待]判断是否开启邮箱通知,通知用户
}
