package model

type Page struct {
	Universal
	PageName  string `json:"pageName"`
	PageLabel string `json:"pageLabel"`
	PageCover string `json:"pageCover"`
}
