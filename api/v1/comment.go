package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
	"strconv"
	"strings"
)

type Comment struct {
}

func (Comment) GetBackendList(c *gin.Context) {
	isReview := false
	if index := strings.Index(c.Request.RequestURI, "isReview"); index != -1 {
		isReview = true
	}
	r.Send(c, r.SUCCESS, commentService.GetBackendList(util.BindQuery[model.Condition](c), isReview))
}

func (Comment) UpdateReview(c *gin.Context) {
	commentService.UpdateReview(util.BindJSON[model.ReviewVO](c))
	r.Send(c, r.SUCCESS)
}

func (Comment) Delete(c *gin.Context) {
	commentService.Delete(util.BindJSON[[]int](c))
	r.Send(c, r.SUCCESS)
}

func (Comment) GetList(c *gin.Context) {
	r.Send(c, r.SUCCESS, commentService.GetList(util.BindQuery[model.CommentVO](c)))
}

func (Comment) GetReplyListByCommentId(c *gin.Context) {
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	r.Send(c, r.SUCCESS, commentService.GetReplyListByCommentId(commentId, util.BindQuery[model.Condition](c)))
}

func (Comment) SaveComment(c *gin.Context) {
	uid, _ := c.Get("userInfoId")
	commentService.SaveComment(uid.(int), util.BindJSON[model.CommentVO](c))
	r.Send(c, r.SUCCESS)
}

func (Comment) SaveLike(c *gin.Context) {
	uid, _ := c.Get("userInfoId")
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	commentService.SaveLike(uid.(int), commentId)
	r.Send(c, r.SUCCESS)
}
