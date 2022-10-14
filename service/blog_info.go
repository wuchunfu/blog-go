package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"myblog/dao"
	"myblog/model"
	"myblog/util"
	"sort"
	"strconv"
	"strings"
)

type BlogInfoService struct {
}

func (*BlogInfoService) Report(c *gin.Context) {
	ipAddress := util.IpUtil.GetIpAddress(c)
	userAgent := util.IpUtil.GetUserAgent(c)
	browser := userAgent.Name + " " + userAgent.Version.String()
	os := userAgent.OS + " " + userAgent.OSVersion.String()
	uuid := util.Generator.MD5(ipAddress + browser + os)
	if !rdb.SIsMember(UniqueVisitor, uuid).Val() {
		ipSource := util.IpUtil.GetIpSource(ipAddress)
		if ipSource != "" {
			address := strings.Split(ipSource, "|")
			ipSource = strings.ReplaceAll(address[2], "省", "")
			rdb.HIncrBy(VisitorArea, ipSource, 1)
		} else {
			rdb.HIncrBy(VisitorArea, "未知", 1)
		}
		rdb.Incr(BlogViewsCount)
		rdb.SAdd(UniqueVisitor, uuid)
	}
}

func (b *BlogInfoService) GetHomeInfo() model.BlogHomeInfoDTO {
	articleCount := dao.Count(model.Article{}, "status = ? AND is_delete = ?", 1, 0)
	categoryCount := dao.Count(model.Category{}, "")
	tagCount := dao.Count(model.Tag{}, "")
	viewsCount, _ := rdb.Get(BlogViewsCount).Int()
	websiteConfigVO := b.GetWebsiteConfig()
	pageList := dao.List([]model.Page{}, "*", "", "")
	return model.BlogHomeInfoDTO{
		ArticleCount:  articleCount,
		CategoryCount: categoryCount,
		TagCount:      tagCount,
		ViewsCount:    viewsCount,
		WebsiteConfig: websiteConfigVO,
		PageList:      pageList,
	}
}

func (*BlogInfoService) GetBackendInfo() model.BlogBackInfoDTO {
	viewsCount, _ := rdb.Get(BlogViewsCount).Int()
	messageCount := dao.Count(model.Message{}, "")
	userCount := dao.Count(model.UserInfo{}, "")
	articleCount := dao.Count(model.Article{}, "is_delete = ?", 0)
	articleStatisticsList := articleDao.ArticleStatisticsList()
	categoryDTOList := categoryDao.GetCategoryDTOList()
	tagDTOList := tagDao.TagDTOList()
	uniqueViewList := uniqueService.UniqueViewList()
	articleMap := rdb.ZRangeWithScores(ArticleViewsCount, 0, 4).Val()
	articleRankDTOList := ArticleListRank(articleMap)
	blogBackInfoDTO := model.BlogBackInfoDTO{
		ViewsCount:            viewsCount,
		MessageCount:          messageCount,
		UserCount:             userCount,
		ArticleCount:          articleCount,
		CategoryDTOList:       categoryDTOList,
		TagDTOList:            tagDTOList,
		ArticleStatisticsList: articleStatisticsList,
		UniqueViewDTOList:     uniqueViewList,
		ArticleRankDTOList:    articleRankDTOList,
	}
	return blogBackInfoDTO
}

func (*BlogInfoService) GetWebsiteConfig() (websiteConfigVO model.WebsiteConfigVO) {
	websiteConfig, err := rdb.Get(WebsiteConfig).Result()
	if websiteConfig == "" || err != nil {
		websiteConfig = dao.GetOne(model.WebsiteConfig{}, "id = ?", 1).Config
		rdb.Set(WebsiteConfig, websiteConfig, 0)
	}
	util.Json.Unmarshal(websiteConfig, &websiteConfigVO)
	return
}

func (*BlogInfoService) UpdateWebsiteConfig(websiteConfigVO model.WebsiteConfigVO) {
	config := util.Json.Marshal(websiteConfigVO)
	websiteConfig := model.WebsiteConfig{
		Universal: model.Universal{ID: 1},
		Config:    config,
	}
	dao.Updates(&websiteConfig, "config")
	rdb.Del(WebsiteConfig)
}

func (*BlogInfoService) GetAbout() string {
	return rdb.Get(About).Val()
}

func (*BlogInfoService) UpdateAbout(data model.About) {
	rdb.Set(About, data.AboutContent, 0)
}

func ArticleListRank(articleMap []redis.Z) (list []model.ArticleRankDTO) {
	var ids []int
	aMap := make(map[int]int)
	for _, article := range articleMap {
		id, _ := strconv.Atoi(article.Member.(string))
		ids = append(ids, id)
		aMap[id] = int(article.Score)
	}
	alist := dao.List([]model.Article{}, "*", "", "id in ?", ids)
	for _, artilce := range alist {
		list = append(list, model.ArticleRankDTO{
			ArticleTitle: artilce.ArticleTitle,
			ViewsCount:   aMap[artilce.ID],
		})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].ViewsCount > list[j].ViewsCount
	})
	return
}
