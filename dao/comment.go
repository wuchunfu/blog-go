package dao

import (
	"database/sql"
	"myblog/model"
	"time"
)

type CommentDao struct {
}

func (*CommentDao) CountBackendCommentDTO(condition model.Condition, isReview bool) int {
	var count int64
	tx := db.Select("COUNT(1)").Table("comment c").
		Joins("LEFT JOIN user_info u ON c.user_id = u.id")
	if condition.Type != 0 {
		tx = tx.Where("c.type = ?", condition.Type)
	}
	if isReview {
		tx = tx.Where("c.is_review = ?", condition.IsReview)
	}
	if condition.Keywords != "" {
		tx = tx.Where("u.nickname = ?", "%"+condition.Keywords+"%")
	}
	tx.Count(&count)
	return int(count)
}

func (*CommentDao) BackendCommentDTOList(condition model.Condition, isReview bool) (list []model.BackendCommentDTO) {
	offset := (condition.Current - 1) * condition.Size
	tx := db.Select("c.id,u.avatar,u.nickname,r.nickname AS reply_nickname," +
		"a.article_title,c.comment_content,c.type,c.create_time,c.is_review").
		Table("comment c").
		Joins("LEFT JOIN article a ON c.topic_id = a.id").
		Joins("LEFT JOIN user_info u ON c.user_id = u.id").
		Joins("LEFT JOIN user_info r ON c.reply_user_id = r.id")
	if condition.Type != 0 {
		tx = tx.Where("c.type = ?", condition.Type)
	}
	if isReview {
		tx = tx.Where("c.is_review = ?", condition.IsReview)
	}
	if condition.Keywords != "" {
		tx = tx.Where("u.nickname = ?", "%"+condition.Keywords+"%")
	}
	tx.Order("id desc").Limit(condition.Size).Offset(offset).Find(&list)
	return
}

func (*CommentDao) UpdateCommentsReview(data model.ReviewVO) {
	db.Select("is_review").Where("id in ?", data.IdList).Updates(model.Comment{IsReview: data.IsReview})
}

func (*CommentDao) ListCountByTopicIds(ids []int, type_ int) (list []model.CommentCountDTO) {
	db.Select("topic_id AS id,COUNT( 1 ) AS comment_count").
		Table("comment").Where("topic_id IN ? AND type = ? AND parent_id IS NULL", ids, type_).
		Group("topic_id").Find(&list)
	return
}

func (*CommentDao) CountComment(vo model.CommentVO) int {
	var count int64
	tx := db.Table("comment").Where("type = ? AND is_review = 1 AND parent_id IS NULL", vo.Type)
	if vo.TopicId != 0 {
		tx = tx.Where("topic_id = ?", vo.TopicId)
	}
	tx.Count(&count)
	return int(count)
}

func (*CommentDao) GetCommentList(vo model.CommentVO) (list []model.CommentDTO) {
	offset := (vo.Current - 1) * 10
	tx := db.Select("u.nickname,u.avatar,u.web_site,c.user_id,c.id,c.comment_content,c.create_time").
		Table("comment c").Joins("JOIN user_info u ON c.user_id = u.id").
		Where("type = ? AND c.is_review = 1 AND parent_id IS NULL", vo.Type)
	if vo.TopicId != 0 {
		tx = tx.Where("topic_id = ?", vo.TopicId)
	}
	tx.Order("c.id desc").Limit(10).Offset(offset).Find(&list)
	return
}

func (*CommentDao) GetReplyList(ids []int) (list []model.ReplyDTO) {
	b := db.Select("c.user_id,u.nickname,u.avatar,u.web_site,c.reply_user_id,r.nickname AS reply_nickname,"+
		"r.web_site AS reply_web_site,c.id,c.parent_id,c.comment_content,c.create_time").
		Table("comment c").
		Joins("JOIN user_info u ON c.user_id = u.id").
		Joins("JOIN user_info r ON c.reply_user_id = r.id").
		Where("c.is_review = 1 AND parent_id IN ?", ids).
		Order("c.parent_id,c.create_time asc")
	t := db.Select("if(@mno=b.parent_id,@rank:=@rank+1,@rank:=1) as row_number,@mno:=b.parent_id,b.*").
		Table("(?) b", b)
	db.Table("(?) t", t).Where("row_number < 4").Find(&list)
	return
}

func (*CommentDao) GetReplyCountListByCommentId(ids []int) (list []model.ReplyCountDTO) {
	db.Select("parent_id as comment_id,count(1) AS reply_count").
		Table("comment").Where("is_review = 1 AND parent_id in ?", ids).
		Group("parent_id").Find(&list)
	return
}

func (*CommentDao) GetReplyListByCommentId(id int, condition model.Condition) (list []model.ReplyDTO) {
	offset := (condition.Current - 1) * condition.Size
	db.Select("c.user_id,u.nickname,u.avatar,u.web_site,c.reply_user_id,r.nickname as reply_nickname,"+
		"r.web_site as reply_web_site,c.id,c.parent_id,c.comment_content,c.create_time").
		Table("comment c").
		Joins("JOIN user_info u ON c.user_id = u.id").
		Joins("JOIN user_info r ON c.reply_user_id = r.id").
		Where("c.is_review = 1 AND parent_id = ?", id).
		Order("c.id asc").Limit(condition.Size).Offset(offset).Find(&list)
	return
}

func (*CommentDao) SaveComment(c model.Comment) {
	db.Exec("INSERT INTO"+
		" `comment` (user_id,topic_id,comment_content,reply_user_id,parent_id,type,is_delete,is_review,create_time)"+
		" values(?,?,?,?,?,?,?,?,?)", c.UserId, nullInt64(c.TopicId), c.CommentContent, nullInt64(c.ReplyUserId),
		nullInt64(c.ParentId), c.Type, 0, c.IsReview, time.Now())
}

func nullInt64(i int) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: int64(i),
		Valid: true,
	}
}
