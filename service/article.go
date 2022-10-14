package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"myblog/config"
	"myblog/dao"
	"myblog/internal/set"
	"myblog/model"
	"myblog/util"
	"strconv"
	"strings"
)

type ArticleService struct {
}

func (*ArticleService) GetBackendList(condition model.Condition) model.PageResult[[]model.ArticleBackendDTO] {
	count := articleDao.CountBackendArticle(condition)
	if count == 0 {
		return model.PageResult[[]model.ArticleBackendDTO]{}
	}
	articleBackendDTOList := articleDao.BackendArticleList(condition)
	// [待] 查询文章点赞量和浏览量
	viewsCountZ := rdb.ZRangeWithScores(ArticleViewsCount, 0, -1).Val()
	likeCountMap := rdb.HGetAll(ArticleLikeCount).Val()
	viewCountMap := getViewCountMap(viewsCountZ)
	for i, article := range articleBackendDTOList {
		articleBackendDTOList[i].ViewsCount = viewCountMap[article.ID]
		articleBackendDTOList[i].LikeCount, _ = strconv.Atoi(likeCountMap[strconv.Itoa(article.ID)])
	}
	return model.PageResult[[]model.ArticleBackendDTO]{Count: count, RecordList: articleBackendDTOList}
}

func (*ArticleService) SaveOrUpdate(articleVO model.ArticleVO, id int) {
	websiteConfig := blogInfoService.GetWebsiteConfig()
	category := saveArticleCategory(articleVO)
	article := model.Article{
		Universal:      model.Universal{ID: articleVO.ID},
		ArticleCover:   articleVO.ArticleCover,
		ArticleTitle:   articleVO.ArticleTitle,
		ArticleContent: articleVO.ArticleContent,
		Type:           articleVO.Type,
		OriginalUrl:    articleVO.OriginalUrl,
		IsTop:          articleVO.IsTop,
		Status:         articleVO.Status,
	}
	if !category.IsEmpty() {
		article.CategoryId = category.ID
	}
	if articleVO.ArticleCover == "" {
		articleVO.ArticleCover = websiteConfig.ArticleCover
	}
	article.UserId = id
	articleDao.SaveOrUpdateArticle(&article)
	saveArticleTag(articleVO, article.ID)
}

func (*ArticleService) UpdateTop(article model.Article) {
	dao.Updates(&article, "is_top")
}

func (*ArticleService) GetBackendById(articleId int) model.ArticleVO {
	article := dao.GetOne(model.Article{}, "id = ?", articleId)
	category := dao.GetOne(model.Category{}, "id = ?", article.CategoryId)
	tagNameList := tagDao.TagNameListByArticleId(articleId)
	return model.ArticleVO{
		ID:             articleId,
		ArticleTitle:   article.ArticleTitle,
		ArticleContent: article.ArticleContent,
		ArticleCover:   article.ArticleCover,
		CategoryName:   category.CategoryName,
		TagNameList:    tagNameList,
		Type:           article.Type,
		OriginalUrl:    article.OriginalUrl,
		IsTop:          article.IsTop,
		Status:         article.Status,
	}
}

func (*ArticleService) UpdateDelete(deleteVO model.DeleteVO) {
	for _, id := range deleteVO.IdList {
		article := model.Article{
			Universal: model.Universal{ID: id},
			IsTop:     0,
			IsDelete:  deleteVO.IsDelete,
		}
		dao.Updates(&article, "is_top", "is_delete")
	}
}

func (s *ArticleService) Delete(ids []int) {
	dao.Delete(model.ArticleTag{}, "article_id in ?", ids)
	dao.Delete(model.Article{}, "id", ids)
}

func (*ArticleService) Export(ids []int) (urls []string) {
	articleList := dao.List([]model.Article{}, "article_title,article_content", "", "id in ?", ids)
	for _, article := range articleList {
		util.FileUtil.WriteFile(article.ArticleTitle+".md", config.GinConf.MDPath, article.ArticleContent)
		urls = append(urls, config.GinConf.MDUrlPath+article.ArticleTitle+".md")
	}
	return urls
}

func (*ArticleService) Import(article model.Article) {
	articleDao.SaveOrUpdateArticle(&article)
}

func (*ArticleService) GetList(current int) []model.ArticleHomeDTO {
	return articleDao.ArticleList(current)
}

func (s *ArticleService) GetById(c *gin.Context, articleId int) (article model.ArticleDTO) {
	article = articleDao.GetArticleById(articleId)
	// 查询推荐文章
	article.RecommendArticleList = articleDao.RecommendArticleList(articleId)
	// 查询最新文章
	article.NewestArticleList = articleDao.NewestArticleList()
	// 更新文章浏览量
	updateArticleViewsCount(c, articleId)
	article.LastArticle = articleDao.LastArticle(articleId)
	article.NextArticle = articleDao.NextArticle(articleId)
	// 封装点赞量和浏览量
	score := rdb.ZScore(ArticleViewsCount, strconv.Itoa(articleId)).Val()
	article.ViewsCount = int(score)
	like, _ := rdb.HGet(ArticleLikeCount, strconv.Itoa(articleId)).Int()
	article.LinkCount = like
	return
}

func (*ArticleService) GetListBySearch(condition model.Condition) (list []model.ArticleSearchDTO) {
	if condition.Keywords == "" {
		return []model.ArticleSearchDTO{}
	}
	articleList := dao.List([]model.Article{}, "*", "",
		"is_delete = 0 AND status = 1 AND (article_title like ? OR article_content like ?)",
		"%"+condition.Keywords+"%", "%"+condition.Keywords+"%")
	for _, article := range articleList {
		articleContent := article.ArticleContent
		index := strings.Index(articleContent, condition.Keywords)
		if index != -1 {
			preIndex, postIndex := 0, 0
			if index > 25 {
				preIndex = index - 25
			}
			preText := articleContent[preIndex:index]
			last := index + len(condition.Keywords)
			postLength := len(articleContent) - last
			if postLength > 175 {
				postIndex = last + 175
			} else {
				postIndex = last + postLength
			}
			postText := articleContent[index:postIndex]
			articleContent = strings.ReplaceAll(preText+postText, condition.Keywords,
				"<span style='color:#f47466'>"+condition.Keywords+"</span>")
		}
		articleTitle := strings.ReplaceAll(article.ArticleTitle, condition.Keywords,
			"<span style='color:#f47466'>"+condition.Keywords+"</span>")
		list = append(list, model.ArticleSearchDTO{
			ID:             article.ID,
			ArticleTitle:   articleTitle,
			ArticleContent: articleContent,
		})
	}
	return
}

func (*ArticleService) GetArchiveList(current int) model.PageResult[[]model.ArchiveDTO] {
	count := dao.Count(model.Article{}, "")
	articles := articleDao.ArticleList(current)
	var archives []model.ArchiveDTO
	for _, article := range articles {
		archives = append(archives, model.ArchiveDTO{
			ID:           article.ID,
			ArticleTitle: article.ArticleTitle,
			CreateTime:   article.CreateTime,
		})
	}
	return model.PageResult[[]model.ArchiveDTO]{
		Count:      count,
		RecordList: archives,
	}
}

func (*ArticleService) GetListByCondition(condition model.Condition) model.ArticlePreviewListDTO {
	articlePreviewDTOList := articleDao.GetArticleListByCondition(condition)
	var name string
	if condition.CategoryId != 0 {
		name = dao.GetOne(model.Category{}, "id = ?", condition.CategoryId).CategoryName
	} else {
		name = dao.GetOne(model.Tag{}, "id = ?", condition.TagId).TagName
	}
	if len(articlePreviewDTOList) == 0 {
		articlePreviewDTOList = []model.ArticlePreviewDTO{}
	}
	return model.ArticlePreviewListDTO{Name: name, ArticlePreviewDTOList: articlePreviewDTOList}
}

func (*ArticleService) SaveLike(uid int, articleId int) {
	articleLikeKey := ArticleUserLike + strconv.Itoa(uid)
	if rdb.SIsMember(articleLikeKey, articleId).Val() {
		rdb.SRem(articleLikeKey, articleId)
		rdb.HIncrBy(ArticleLikeCount, strconv.Itoa(articleId), -1)
	} else {
		rdb.SAdd(articleLikeKey, articleId)
		rdb.HIncrBy(ArticleLikeCount, strconv.Itoa(articleId), 1)
	}
}

func saveArticleCategory(article model.ArticleVO) model.Category {
	category := dao.GetOne(model.Category{}, "category_name", article.CategoryName)
	if category.IsEmpty() && article.Status != model.DRAFT {
		category.CategoryName = article.CategoryName
		dao.Create(&category)
	}
	return category
}

func saveArticleTag(article model.ArticleVO, articleId int) {
	if article.ID != 0 {
		dao.Delete(model.ArticleTag{}, "article_id", article.ID)
	}
	tagNameList := article.TagNameList
	var articleTagList []model.ArticleTag
	for _, tagName := range tagNameList {
		tag := dao.GetOne(model.Tag{}, "tag_name = ?", tagName)
		if tag.IsEmpty() {
			tag.TagName = tagName
			dao.Create(&tag)
		}
		articleTagList = append(articleTagList, model.ArticleTag{ArticleId: articleId, TagId: tag.ID})
	}
	dao.Create(&articleTagList)
}

func updateArticleViewsCount(c *gin.Context, articleId int) {
	session := sessions.Default(c)
	if session.Get("articleSet") == nil {
		var articleSet set.Set
		articleSet.Init()
		articleSet.Add(articleId)
		session.Set("articleSet", articleSet)
		rdb.ZIncrBy(ArticleViewsCount, 1, strconv.Itoa(articleId))
	} else {
		articleSet := session.Get("articleSet").(set.Set)
		if !articleSet.Exist(articleId) {
			articleSet.Add(articleId)
			session.Set("articleSet", articleSet)
			rdb.ZIncrBy(ArticleViewsCount, 1, strconv.Itoa(articleId))
		}
	}
}

func getViewCountMap(rz []redis.Z) map[int]int {
	m := make(map[int]int)
	for _, article := range rz {
		id, _ := strconv.Atoi(article.Member.(string))
		m[id] = int(article.Score)
	}
	return m
}
