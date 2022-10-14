package model

import "time"

type Menu struct {
	Universal
	Name      string `json:"name"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Icon      string `json:"icon"`
	OrderNum  int    `json:"orderNum"`
	ParentId  int    `json:"parentId"`
	IsHidden  int    `json:"isHidden"`
}

type UserMenuDTO struct {
	Name      string        `json:"name"`
	Path      string        `json:"path"`
	Component string        `json:"component"`
	Icon      string        `json:"icon"`
	Hidden    bool          `json:"hidden"`
	Children  []UserMenuDTO `json:"children"`
}

type MenuDTO struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Path       string    `json:"path"`
	Component  string    `json:"component"`
	Icon       string    `json:"icon"`
	CreateTime time.Time `json:"createTime"`
	OrderNum   int       `json:"orderNum"`
	IsHidden   int       `json:"isHidden"`
	Children   []MenuDTO `json:"children"`
}
