package model

import (
	"reflect"
	"time"
)

type Role struct {
	Universal
	RoleName  string `json:"roleName"`
	RoleLabel string `json:"roleLabel"`
	IsDisable int    `json:"isDisable"`
}

func (r *Role) IsEmpty() bool {
	return reflect.DeepEqual(r, &Role{})
}

type RoleDTO struct {
	ID             int       `json:"id"`
	RoleName       string    `json:"roleName"`
	RoleLabel      string    `json:"roleLabel"`
	CreateTime     time.Time `json:"createTime"`
	IsDisable      string    `json:"isDisable"`
	ResourceIdList []int     `json:"resourceIdList" gorm:"-"`
	MenuIdList     []int     `json:"menuIdList" gorm:"-"`
}

type RoleVO struct {
	ID             int    `json:"id"`
	RoleName       string `json:"roleName"`
	RoleLabel      string `json:"roleLabel"`
	IsDisable      int    `json:"isDisable"`
	ResourceIdList []int  `json:"resourceIdList"`
	MenuIdList     []int  `json:"menuIdList"`
}
