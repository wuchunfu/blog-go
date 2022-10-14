package model

import (
	"encoding/gob"
	"time"
)

type UserInfo struct {
	Universal
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Intro     string `json:"intro"`
	WebSite   string `json:"webSite"`
	IsDisable int    `json:"isDisable"`
}

type UserInfoVO struct {
	Nickname string `json:"nickname"`
	Intro    string `json:"intro"`
	Website  string `json:"website"`
}

type UserInfoDTO struct {
	ID             int       `json:"id"`
	UserInfoId     int       `json:"userInfoId"`
	Email          string    `json:"email"`
	LoginType      int       `json:"loginType"`
	Username       string    `json:"username"`
	Nickname       string    `json:"nickname"`
	Avatar         string    `json:"avatar"`
	Intro          string    `json:"intro"`
	WebSite        string    `json:"webSite"`
	ArticleLikeSet []string  `json:"articleLikeSet"`
	CommentLikeSet []string  `json:"commentLikeSet"`
	TalkLikeSet    []string  `json:"talkLikeSet"`
	IpAddress      string    `json:"ipAddress"`
	IpSource       string    `json:"ipSource"`
	LastLoginTime  time.Time `json:"lastLoginTime"`
	Token          string    `json:"token"`
}

type UserDetailDTO struct {
	UserInfoDTO
	Password  string   `json:"password"`
	RoleList  []string `json:"roleList"`
	IsDisable int      `json:"isDisable"`
	Browser   string   `json:"browser"`
	Os        string   `json:"os"`
}

type BackendUserDTO struct {
	ID            int       `json:"id"`
	UserInfoId    int       `json:"userInfoId"`
	Avatar        string    `json:"avatar"`
	Nickname      string    `json:"nickname"`
	RoleList      []Role    `json:"roleList" gorm:"many2many:user_role;foreignKey:UserInfoId;joinForeignKey:UserId;"`
	LoginType     int       `json:"loginType"`
	IpAddress     string    `json:"ipAddress"`
	IpSource      string    `json:"ipSource"`
	CreateTime    time.Time `json:"createTime"`
	LastLoginTime time.Time `json:"lastLoginTime"`
	IsDisable     int       `json:"isDisable"`
	Status        int       `json:"status"`
}

type UserOnlineDTO struct {
	UserIndoId    int       `json:"userInfoId"`
	Nickname      string    `json:"nickname"`
	Avatar        string    `json:"avatar"`
	IpAddress     string    `json:"ipAddress"`
	IpSource      string    `json:"ipSource"`
	Browser       string    `json:"browser"`
	OS            string    `json:"os"`
	LastLoginTime time.Time `json:"lastLoginTime"`
}

type SessionInfo struct {
	UserDetailDTO
	IsOffline int `json:"isOffline"`
}

func init() {
	gob.Register(SessionInfo{})
}

type UserVO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code"`
}
