package model

import (
	"reflect"
	"time"
)

type UserAuth struct {
	Universal
	UserInfoId    int       `json:"userInfoId"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	LoginType     int       `json:"loginType"`
	IpAddress     string    `json:"ipAddress"`
	IpSource      string    `json:"ipSource"`
	LastLoginTime time.Time `json:"lastLoginTime"`
}

func (u *UserAuth) IsEmpty() bool {
	return reflect.DeepEqual(u, &UserAuth{})
}

type UserToken struct {
	UserAuth
	Token string `json:"token"`
}

type PasswordVO struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type UserAreaDTO struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
