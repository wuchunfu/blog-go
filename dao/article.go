package dao

import (
	"gorm.io/gorm"
	"myblog/model"
	"myblog/util"
)

type ArticleDao struct {
}

func (*ArticleDao) CountBackendArticle(condition model.Condition) int {
	var count int64
	whereMap := make(map[string]any)
	whereMap["is_delete"] = condition.IsDelete
	if condition.Keywords != "" {
		whereMap["article_title like"] = "%" + condition.Keywords + "%"
	}
	if condition.Status != 0 {
		whereMap["status"] = condition.Status
	}
	if condition.CategoryId != 0 {
		whereMap["category_id"] = condition.CategoryId
	}
	if condition.Type != 0 {
		whereMap["type"] = condition.Type
	}
	if condition.TagId != 0 {
		whereMap["at.tag_id"] = condition.TagId
	}
	query, args := util.GetCondition(whereMap)
	db.Table("article a").
		Joins("LEFT JOIN article_tag at on  a.id = at.article_id").
		Where(query, args...).
		Distinct("a.id").Count(&count)
	return int(count)
}

func (*ArticleDao) BackendArticleList(condition model.Condition) (list []model.ArticleBackendDTO) {
	whereMap := make(map[string]any)
	whereMap["is_delete"] = condition.IsDelete
	if condition.Keywords != "" {
		whereMap["article_title like"] = "%" + condition.Keywords + "%"
	}
	if condition.Status != 0 {
		whereMap["status"] = condition.Status
	}
	if condition.CategoryId != 0 {
		whereMap["category_id"] = condition.CategoryId
	}
	if condition.Type != 0 {
		whereMap["type"] = condition.Type
	}
	offset := (condition.Current - 1) * condition.Size
	query, args := util.GetCondition(whereMap)

	table := db.Table("article").Select("id", "article_cover", "article_title",
		"type", "is_top", "is_delete", "status", "create_time", "category_id").
		Order("is_top,id desc").Where(query, args...).Limit(condition.Size).Offset(offset)

	db.Table("(?) a", table).Select("a.id", "article_cover", "article_title",
		"type", "is_top", "a.is_delete", "a.status", "a.create_time", "category_name").
		Preload("Tags", func(db2 *gorm.DB) *gorm.DB {
			return db2.Select("id", "tag_name")
		}). // 尤其要注意这里是属性名字而不是属性类型的名字！！！
		Joins("LEFT JOIN category c ON a.category_id = c.id").
		Order("is_top,a.id desc").
		Find(&list)

	return
}

func (*ArticleDao) SaveOrUpdateArticle(article *model.Article) {
	if article.ID == 0 {
		Create(article)
		return
	}
	Updates(article)
}

func (*ArticleDao) ArticleList(current int) (list []model.ArticleHomeDTO) {
	offset := (current - 1) * 10

	table := db.Table("article").
		Select("id,article_cover,article_title,article_content,type,is_top,create_time,category_id").
		Where("is_delete = 0 AND status = 1").
		Order("is_top,id desc").
		Limit(10).Offset(offset)

	db.Table("(?) a", table).
		Select("a.id,article_cover,article_title,SUBSTR( article_content, 1, 500 ) AS article_content,"+
			"a.create_time,a.type,a.is_top,a.category_id,category_name").
		Preload("Tags", func(db2 *gorm.DB) *gorm.DB {
			return db2.Select("id", "tag_name")
		}).
		Joins("LEFT JOIN category c ON a.category_id = c.id").
		Order("a.is_top,a.id desc").
		Find(&list)

	return
}

func (*ArticleDao) GetArticleById(articleId int) (article model.ArticleDTO) {
	db.Table("article a").
		Select("a.id,article_cover,article_title,article_content,a.type,a.original_url,"+
			"a.create_time,a.update_time,a.category_id,category_name").
		Preload("Tags", func(db2 *gorm.DB) *gorm.DB {
			return db2.Select("id", "tag_name")
		}).
		Joins("JOIN category c ON a.category_id = c.id").
		Where("a.id = ? AND a.is_delete = 0 AND a.status = 1", articleId).
		Find(&article)
	return
}

func (*ArticleDao) RecommendArticleList(articleId int) (list []model.RecommendArticleDTO) {
	table1 := db.
		Table("article_tag").
		Select("tag_id").
		Where("article_id = ?", articleId)
	table2 := db.
		Table("(?) t", table1).
		Select("DISTINCT article_id").
		Joins("JOIN article_tag t1 ON t.tag_id = t1.tag_id").
		Where("article_id != ?", articleId)
	db.Table("(?) t2", table2).
		Select("id,article_title,article_cover,create_time").
		Joins("JOIN article a ON t2.article_id = a.id").
		Where("a.is_delete = 0").
		Order("is_top,id desc").
		Limit(6).Find(&list)
	return
}

func (*ArticleDao) NewestArticleList() (list []model.RecommendArticleDTO) {
	db.Select("id,article_title,article_cover,create_time").
		Table("article").
		Where("is_delete = 0 AND status = 1").
		Order("create_time desc, id asc").
		Limit(5).Find(&list)
	return
}

func (*ArticleDao) LastArticle(articleId int) (article model.ArticlePaginationDTO) {
	db.Select("id,article_title,article_cover").
		Table("article").
		Where("is_delete = 0 AND status = 1 AND id < ?", articleId).
		Limit(1).Find(&article)
	return
}

func (*ArticleDao) NextArticle(articleId int) (article model.ArticlePaginationDTO) {
	db.Select("id,article_title,article_cover").
		Table("article").
		Where("is_delete = 0 AND status = 1 AND id > ?", articleId).
		Limit(1).Find(&article)
	return
}

func (*ArticleDao) ArticleStatisticsList() (list []model.ArticleStatisticsDTO) {
	db.Select("DATE_FORMAT( create_time, \"%Y-%m-%d\" ) AS date,COUNT( 1 ) AS count").
		Table("article").Group("date").Order("date desc").Find(&list)
	return
}

func (*ArticleDao) GetArticleListByCondition(condition model.Condition) (list []model.ArticlePreviewDTO) {
	offset := (condition.Current - 1) * 9
	table := db.Select("id,article_cover,article_title,article_content,create_time,category_id").
		Table("article").Where("is_delete = 0").Where("status = 1")
	if condition.CategoryId != 0 {
		table = table.Where("category_id = ?", condition.CategoryId)
	}
	if condition.TagId != 0 {
		table = table.Where("id IN (SELECT article_id FROM article_tag WHERE tag_id = ?)", condition.TagId)
	}
	table = table.Order("id desc").Limit(condition.Size).Offset(offset)
	db.Select("a.id,article_cover,article_title,a.create_time,a.category_id,category_name").
		Table("(?) a", table).Preload("Tags").
		Joins("JOIN category c ON a.category_id = c.id").
		Find(&list)
	return
}
