package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/config"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
)

type BlogInfo struct {
}

// GetHomeInfo 获取博客信息
func (BlogInfo) GetHomeInfo(c *gin.Context) {
	r.Send(c, r.SUCCESS, blogInfoService.GetHomeInfo())
}

// GetBackendInfo 获取后台的博客信息
func (BlogInfo) GetBackendInfo(c *gin.Context) {
	r.Send(c, r.SUCCESS, blogInfoService.GetBackendInfo())
}

// GetWebsiteConfig 获取博客配置信息
func (BlogInfo) GetWebsiteConfig(c *gin.Context) {
	r.Send(c, r.SUCCESS, blogInfoService.GetWebsiteConfig())
}

// SaveConfigImage 上传配置图片
func (BlogInfo) SaveConfigImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	imgUrl := util.FileUtil.UploadFile(file, config.TcConf.UserPath)
	r.Send(c, r.SUCCESS, imgUrl)
}

// UpdateWebsiteConfig 更新博客配置信息
func (BlogInfo) UpdateWebsiteConfig(c *gin.Context) {
	blogInfoService.UpdateWebsiteConfig(util.BindJSON[model.WebsiteConfigVO](c))
	r.Send(c, r.SUCCESS)
}

// GetAbout 获取关于我
func (BlogInfo) GetAbout(c *gin.Context) {
	r.Send(c, r.SUCCESS, blogInfoService.GetAbout())
}

// UpdateAbout 更新关于我
func (BlogInfo) UpdateAbout(c *gin.Context) {
	blogInfoService.UpdateAbout(util.BindJSON[model.About](c))
	r.Send(c, r.SUCCESS)
}

// Report 进入时传递相关信息，如ip等信息
func (BlogInfo) Report(c *gin.Context) {
	blogInfoService.Report(c)
	r.Send(c, r.SUCCESS)
}
