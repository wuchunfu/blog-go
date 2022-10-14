package model

type UniqueView struct {
	Universal
	ViewsCount int `json:"viewsCount"`
}

type UniqueViewDTO struct {
	Day        string `json:"day"`
	ViewsCount int    `json:"viewsCount"`
}
