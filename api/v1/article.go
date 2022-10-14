package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/config"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
	"strconv"
)

type Article struct {
}

// GetBackendList 获取后台文章列表
func (Article) GetBackendList(c *gin.Context) {
	r.Send(c, r.SUCCESS, articleService.GetBackendList(util.BindQuery[model.Condition](c)))
}

// UploadImages 上传文章封面图片
func (Article) UploadImages(c *gin.Context) {
	file, _ := c.FormFile("file")
	r.Send(c, r.SUCCESS, util.FileUtil.UploadFile(file, config.TcConf.ArticlePath))
}

// SaveOrUpdate 更新或修改文章
func (Article) SaveOrUpdate(c *gin.Context) {
	id, _ := c.Get("userInfoId")
	articleService.SaveOrUpdate(util.BindJSON[model.ArticleVO](c), id.(int))
	r.Send(c, r.SUCCESS)
}

// UpdateTop 更新文章置顶信息
func (Article) UpdateTop(c *gin.Context) {
	article := util.BindJSON[model.ArticleVO](c)
	articleService.UpdateTop(model.Article{
		Universal: model.Universal{ID: article.ID},
		IsTop:     article.IsTop,
	})
	r.Send(c, r.SUCCESS)
}

// GetBackendById 根据文章id获取后台单个文章
func (Article) GetBackendById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r.Send(c, r.SUCCESS, articleService.GetBackendById(id))
}

// UpdateDelete 软删除文章
func (Article) UpdateDelete(c *gin.Context) {
	articleService.UpdateDelete(util.BindJSON[model.DeleteVO](c))
	r.Send(c, r.SUCCESS)
}

// Delete 物理删除文章
func (Article) Delete(c *gin.Context) {
	articleService.Delete(util.BindJSON[[]int](c))
	r.Send(c, r.SUCCESS)
}

// Export 导出文章
func (Article) Export(c *gin.Context) {
	r.Send(c, r.SUCCESS, articleService.Export(util.BindJSON[[]int](c)))
}

// Import 导入文章
func (Article) Import(c *gin.Context) {
	file, _ := c.FormFile("file")
	content := util.FileUtil.ReadFile(file)
	article := model.Article{
		ArticleTitle:   file.Filename,
		ArticleContent: content,
		Status:         model.DRAFT,
	}
	articleService.Import(article)
	r.Send(c, r.SUCCESS)
}

// =====================================

// GetList 获取博客文章列表
func (Article) GetList(c *gin.Context) {
	current, _ := strconv.Atoi(c.Query("current"))
	r.Send(c, r.SUCCESS, articleService.GetList(current))
}

// GetById 根据文章id获取文章具体内容
func (Article) GetById(c *gin.Context) {
	articleId, _ := strconv.Atoi(c.Param("articleId"))
	r.Send(c, r.SUCCESS, articleService.GetById(c, articleId))
}

// GetListBySearch 查找有相关内容的所有文章
func (Article) GetListBySearch(c *gin.Context) {
	r.Send(c, r.SUCCESS, articleService.GetListBySearch(util.BindQuery[model.Condition](c)))
}

// GetArchiveList 获取文章归档
func (Article) GetArchiveList(c *gin.Context) {
	current, _ := strconv.Atoi(c.Query("current"))
	r.Send(c, r.SUCCESS, articleService.GetArchiveList(current))
}

func (Article) GetListByCondition(c *gin.Context) {
	r.Send(c, r.SUCCESS, articleService.GetListByCondition(util.BindQuery[model.Condition](c)))
}

func (Article) SaveLike(c *gin.Context) {
	uid, _ := c.Get("userInfoId")
	articleId, _ := strconv.Atoi(c.Param("articleId"))
	articleService.SaveLike(uid.(int), articleId)
	r.Send(c, r.SUCCESS)
}
