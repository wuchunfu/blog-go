package model

import "reflect"

type PhotoAlbum struct {
	Universal
	AlbumName  string `json:"albumName"`
	AlbumDesc  string `json:"albumDesc"`
	AlbumCover string `json:"albumCover"`
	IsDelete   int    `json:"isDelete"`
	Status     int    `json:"status"`
}

func (a *PhotoAlbum) IsEmpty() bool {
	return reflect.DeepEqual(a, &PhotoAlbum{})
}

type BackendPhotoAlbumDTO struct {
	ID         int    `json:"id"`
	AlbumName  string `json:"albumName"`
	AlbumDesc  string `json:"albumDesc"`
	AlbumCover string `json:"albumCover"`
	PhotoCount int    `json:"photoCount"`
	Status     int    `json:"status"`
}
