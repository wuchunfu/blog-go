package model

type UserRole struct {
	ID     int `json:"id"`
	UserId int `json:"userId"`
	RoleId int `json:"roleId"`
}

type UserRoleVO struct {
	UserInfoId int    `json:"userInfoId"`
	Nickname   string `json:"nickname"`
	RoleIdList []int  `json:"roleIdList"`
}
