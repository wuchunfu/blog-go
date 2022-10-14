package v1

import (
	"github.com/gin-gonic/gin"
	"myblog/config"
	"myblog/model"
	"myblog/util"
	"myblog/util/r"
	"strconv"
)

type Photo struct {
}

func (Photo) GetBackendAlbumList(c *gin.Context) {
	r.Send(c, r.SUCCESS, photoService.GetBackendAlbumList(util.BindQuery[model.Condition](c)))
}

func (Photo) SaveAlbumCover(c *gin.Context) {
	file, _ := c.FormFile("file")
	r.Send(c, r.SUCCESS, util.FileUtil.UploadFile(file, config.TcConf.AlbumPath))
}

func (Photo) SaveOrUpdateAlbum(c *gin.Context) {
	r.Send(c, photoService.SaveOrUpdateAlbum(util.BindJSON[model.PhotoAlbum](c)))
}

func (Photo) DeleteAlbum(c *gin.Context) {
	albumId, _ := strconv.Atoi(c.Param("albumId"))
	photoService.DeleteAlbum(albumId)
	r.Send(c, r.SUCCESS)
}

func (Photo) GetBackendAlbumInfoList(c *gin.Context) {
	r.Send(c, r.SUCCESS, photoService.GetBackendAlbumInfoList())
}

func (Photo) GetBackendAlbumById(c *gin.Context) {
	albumId, _ := strconv.Atoi(c.Param("albumId"))
	r.Send(c, r.SUCCESS, photoService.GetBackendAlbumById(albumId))
}

func (Photo) GetList(c *gin.Context) {
	r.Send(c, r.SUCCESS, photoService.GetList(util.BindQuery[model.Condition](c)))
}

func (Photo) Save(c *gin.Context) {
	photoService.Save(util.BindJSON[model.PhotoVO](c))
	r.Send(c, r.SUCCESS)
}

func (Photo) Update(c *gin.Context) {
	photoService.Update(util.BindJSON[model.Photo](c))
	r.Send(c, r.SUCCESS)
}

func (Photo) UpdatePhotosAlbum(c *gin.Context) {
	photoService.UpdatePhotosAlbum(util.BindJSON[model.PhotoVO](c))
	r.Send(c, r.SUCCESS)
}

func (Photo) UpdateDelete(c *gin.Context) {
	photoService.UpdateDelete(util.BindJSON[model.DeleteVO](c))
	r.Send(c, r.SUCCESS)
}

func (Photo) Delete(c *gin.Context) {
	photoService.Delete(util.BindJSON[[]int](c))
	r.Send(c, r.SUCCESS)
}

func (Photo) GetAlbumList(c *gin.Context) {
	r.Send(c, r.SUCCESS, photoService.GetAlbumList())
}

func (Photo) GetListByAlbumId(c *gin.Context) {
	albumId, _ := strconv.Atoi(c.Param("albumId"))
	code, photos := photoService.GetListByAlbumId(albumId, util.BindQuery[model.Condition](c))
	r.Send(c, code, photos)
}
