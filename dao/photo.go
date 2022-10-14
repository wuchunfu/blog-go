package dao

import (
	"myblog/model"
)

type PhotoDao struct {
}

func (*PhotoDao) CountBackendPhotoAlbum(condition model.Condition) int {
	var count int64
	tx := db.Model(&model.PhotoAlbum{}).Where("is_delete = 0")
	if condition.Keywords != "" {
		tx = tx.Where("album_name like ?", "%"+condition.Keywords+"%")
	}
	tx.Count(&count)
	return int(count)
}

func (*PhotoDao) BackendPhotoAlbumList(condition model.Condition) (list []model.BackendPhotoAlbumDTO) {
	offset := (condition.Current - 1) * condition.Size
	table := db.Select("id,album_name,album_desc,album_cover,status").
		Table("photo_album").Where("is_delete = 0")
	if condition.Keywords != "" {
		table = table.Where("album_name like ?", "%"+condition.Keywords+"%")
	}
	table = table.Order("id desc").Limit(condition.Size).Offset(offset)
	db.Select("pa.id,album_name,album_desc,album_cover,COUNT( a.id ) AS photo_count,status").
		Table("(?) pa", table).
		Joins("LEFT JOIN (SELECT id,album_id FROM photo WHERE is_delete = 0) a ON pa.id = a.album_id").
		Group("pa.id").Find(&list)
	return
}

func (*PhotoDao) SaveOrUpdatePhotoAlbum(album model.PhotoAlbum) {
	if album.ID != 0 {
		Updates(&album)
	} else {
		Create(&album)
	}
}

func (*PhotoDao) DeletePhotoAlbum(albumId int) {
	Updates(&model.PhotoAlbum{
		Universal: model.Universal{ID: albumId},
		IsDelete:  1,
	}, "is_delete")
	db.Where("album_id = ?", albumId).Updates(model.Photo{IsDelete: 1})
}

func (d *PhotoDao) CountPhoto(condition model.Condition) int {
	var count int64
	tx := db.Model(model.Photo{}).Where("is_delete = ?", condition.IsDelete)
	if condition.AlbumId != 0 {
		tx = tx.Where("album_id = ?", condition.AlbumId)
	}
	tx.Count(&count)
	return int(count)
}

func (*PhotoDao) PhotoList(condition model.Condition) (list []model.Photo) {
	offset := (condition.Current - 1) * condition.Size
	tx := db.Model(model.Photo{}).Where("is_delete = ?", condition.IsDelete)
	if condition.AlbumId != 0 {
		tx = tx.Where("album_id = ?", condition.AlbumId)
	}
	tx.Order("id,update_time desc").
		Limit(condition.Size).Offset(offset).Find(&list)
	return
}

func (*PhotoDao) UpdatePhotosAlbum(data model.PhotoVO) {
	db.Select("album_id").Where("id in ?", data.PhotoIdList).Updates(model.Photo{AlbumId: data.AlbumId})
}

func (*PhotoDao) UpdatePhotoDelete(data model.DeleteVO) {
	db.Select("is_delete").Where("id in ?", data.IdList).Updates(model.Photo{IsDelete: data.IsDelete})
	if data.IsDelete == 0 { // 照片恢复->若相册是软删除状态则恢复相册
		photo := GetOne(model.Photo{}, "id = ?", data.IdList[0])
		db.Select("is_delete").Where("id = ?", photo.AlbumId).Updates(&model.PhotoAlbum{IsDelete: 0})
	}
}
