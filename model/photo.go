package model

type Photo struct {
	Universal
	AlbumId   int    `json:"albumId"`
	PhotoName string `json:"photoName"`
	PhotoDesc string `json:"photoDesc"`
	PhotoSrc  string `json:"photoSrc"`
	IsDelete  int    `json:"isDelete"`
}

type PhotoVO struct {
	AlbumId      int      `json:"albumId"`
	PhotoUrlList []string `json:"photoUrlList"`
	PhotoIdList  []int    `json:"photoIdList"`
}

type PhotoDTO struct {
	PhotoAlbumCover string   `json:"photoAlbumCover"`
	PhotoAlbumName  string   `json:"photoAlbumName"`
	PhotoList       []string `json:"photoList"`
}
