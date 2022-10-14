package model

type Message struct {
	Universal
	Nickname       string `json:"nickname"`
	Avatar         string `json:"avatar"`
	MessageContent string `json:"messageContent"`
	IpAddress      string `json:"ipAddress"`
	IpSource       string `json:"ipSource"`
	Time           int    `json:"time"`
	IsReview       int    `json:"isReview"`
}

type ReviewVO struct {
	IdList   []int `json:"idList"`
	IsReview int   `json:"isReview"`
}

type MessageVO struct {
	Nickname       string `json:"nickname"`
	Avatar         string `json:"avatar"`
	MessageContent string `json:"messageContent"`
	Time           int    `json:"time"`
}
