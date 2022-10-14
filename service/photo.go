package service

import (
	"myblog/dao"
	"myblog/model"
	"myblog/util/r"
	"strings"
)

type PhotoService struct {
}

func (*PhotoService) GetBackendAlbumList(condition model.Condition) model.PageResult[[]model.BackendPhotoAlbumDTO] {
	count := photoDao.CountBackendPhotoAlbum(condition)
	if count == 0 {
		return model.PageResult[[]model.BackendPhotoAlbumDTO]{}
	}
	backendPhotoAlbumList := photoDao.BackendPhotoAlbumList(condition)
	return model.PageResult[[]model.BackendPhotoAlbumDTO]{
		Count:      count,
		RecordList: backendPhotoAlbumList,
	}
}

func (*PhotoService) SaveOrUpdateAlbum(data model.PhotoAlbum) int {
	album := dao.GetOne(model.PhotoAlbum{}, "album_name = ?", data.AlbumName)
	if !album.IsEmpty() && album.ID != data.ID {
		return r.AlbumExist
	}
	photoDao.SaveOrUpdatePhotoAlbum(data)
	return r.SUCCESS
}

func (*PhotoService) DeleteAlbum(albumId int) {
	count := dao.Count(model.Photo{}, "album_id = ?", albumId)
	if count > 0 {
		photoDao.DeletePhotoAlbum(albumId)
	} else {
		dao.Delete(model.PhotoAlbum{}, "id = ?", albumId)
	}
}

func (*PhotoService) GetBackendAlbumInfoList() []model.PhotoAlbum {
	return dao.List([]model.PhotoAlbum{}, "*", "", "is_delete = ?", 0)
}

func (*PhotoService) GetBackendAlbumById(albumId int) model.BackendPhotoAlbumDTO {
	photoAlbum := dao.GetOne(model.PhotoAlbum{}, "id = ?", albumId)
	count := dao.Count(model.Photo{}, "album_id = ? AND is_delete = 0", albumId)
	return model.BackendPhotoAlbumDTO{
		ID:         photoAlbum.ID,
		AlbumName:  photoAlbum.AlbumName,
		AlbumDesc:  photoAlbum.AlbumDesc,
		AlbumCover: photoAlbum.AlbumCover,
		PhotoCount: count,
		Status:     photoAlbum.Status,
	}
}

func (*PhotoService) GetList(condition model.Condition) model.PageResult[[]model.Photo] {
	count := photoDao.CountPhoto(condition)
	if count == 0 {
		return model.PageResult[[]model.Photo]{count, []model.Photo{}}
	}
	photoList := photoDao.PhotoList(condition)
	return model.PageResult[[]model.Photo]{
		Count:      count,
		RecordList: photoList,
	}
}

func (*PhotoService) Save(data model.PhotoVO) {
	var photoList []model.Photo
	for _, url := range data.PhotoUrlList {
		photoList = append(photoList, model.Photo{
			AlbumId:   data.AlbumId,
			PhotoName: url[strings.LastIndex(url, "/")+1:],
			PhotoSrc:  url,
		})
	}
	dao.Create(&photoList)
}

func (*PhotoService) Update(photo model.Photo) {
	dao.Updates(&photo, "photo_desc")
}

func (s *PhotoService) UpdatePhotosAlbum(data model.PhotoVO) {
	photoDao.UpdatePhotosAlbum(data)
}

func (*PhotoService) UpdateDelete(data model.DeleteVO) {
	photoDao.UpdatePhotoDelete(data)
}

func (*PhotoService) Delete(ids []int) {
	dao.Delete(model.Photo{}, "id in ?", ids)
}

func (*PhotoService) GetAlbumList() []model.PhotoAlbum {
	return dao.List([]model.PhotoAlbum{}, "id,album_name,album_desc,album_cover",
		"id desc", "status = 1 AND is_delete = 0")
}

func (*PhotoService) GetListByAlbumId(albumId int, condition model.Condition) (int, model.PhotoDTO) {
	album := dao.GetOne(model.PhotoAlbum{},
		"id = ? AND is_delete = ? AND status = ?", albumId, 0, 1)
	if album.IsEmpty() {
		return r.AlbumNotExist, model.PhotoDTO{}
	}
	photoList := photoDao.PhotoList(condition)
	photos := make([]string, 0)
	for _, photo := range photoList {
		photos = append(photos, photo.PhotoSrc)
	}
	return r.SUCCESS, model.PhotoDTO{
		PhotoAlbumCover: album.AlbumCover,
		PhotoAlbumName:  album.AlbumName,
		PhotoList:       photos,
	}
}
