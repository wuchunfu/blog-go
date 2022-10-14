package model

import "time"

type Comment struct {
	Universal
	UserId         int    `json:"userId"`
	TopicId        int    `json:"topicId"`
	CommentContent string `json:"commentContent"`
	ReplyUserId    int    `json:"replyUserId"`
	ParentId       int    `json:"parentId"`
	Type           int    `json:"type"`
	IsDelete       int    `json:"isDelete"`
	IsReview       int    `json:"isReview"`
}

type BackendCommentDTO struct {
	ID             int       `json:"id"`
	Avatar         string    `json:"avatar"`
	Nickname       string    `json:"nickname"`
	ReplyNickname  string    `json:"replyNickname"`
	ArticleTitle   string    `json:"articleTitle"`
	CommentContent string    `json:"commentContent"`
	Type           int       `json:"type"`
	IsReview       int       `json:"isReview"`
	CreateTime     time.Time `json:"createTime"`
}

type CommentCountDTO struct {
	ID           int `json:"id"`
	CommentCount int `json:"commentCount"`
}

type ReplyCountDTO struct {
	CommentId  int `json:"commentId"`
	ReplyCount int `json:"replyCount"`
}

type CommentVO struct {
	Current        int    `json:"current" form:"current"`
	Size           int    `json:"size" form:"size"`
	ReplyUserId    int    `json:"replyUserId" form:"replyUserId"`
	TopicId        int    `json:"topicId,string" form:"topicId"`
	CommentContent string `json:"commentContent" form:"commentContent"`
	ParentId       int    `json:"parentId" form:"parentId"`
	Type           int    `json:"type" form:"type"`
}

type CommentDTO struct {
	ID             int        `json:"id"`
	UserId         int        `json:"userId"`
	Nickname       string     `json:"nickname"`
	Avatar         string     `json:"avatar"`
	Website        string     `json:"website"`
	CommentContent string     `json:"commentContent"`
	LikeCount      int        `json:"likeCount"`
	CreateTime     time.Time  `json:"createTime"`
	ReplyCount     int        `json:"replyCount"`
	ReplyDTOList   []ReplyDTO `json:"replyDTOList" gorm:"-"`
}

type ReplyDTO struct {
	ID             int       `json:"id"`
	ParentId       int       `json:"parentId"`
	UserId         int       `json:"userId"`
	Nickname       string    `json:"nickname"`
	Avatar         string    `json:"avatar"`
	Website        string    `json:"website"`
	ReplyUserId    int       `json:"replyUserId"`
	ReplyNickname  string    `json:"replyNickname"`
	ReplyWebSite   string    `json:"replyWebSite"`
	CommentContent string    `json:"commentContent"`
	LikeCount      int       `json:"likeCount"`
	CreateTime     time.Time `json:"createTime"`
}
