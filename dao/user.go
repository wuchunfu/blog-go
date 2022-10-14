package dao

import (
	"myblog/model"
	"myblog/util"
)

type UserDao struct {
}

func (*UserDao) CountUser(condition model.Condition) int {
	var count int64
	tx := db.Select("COUNT(1)").Table("user_auth ua").
		Joins("LEFT JOIN user_info ui ON ua.user_info_id = ui.id")
	if condition.Keywords != "" {
		tx = tx.Where("nickname like ?", "%"+condition.Keywords+"%")
	}
	if condition.LoginType != 0 {
		tx = tx.Where("login_type = ?", condition.LoginType)
	}
	tx.Count(&count)
	return int(count)
}

func (*UserDao) UserList(condition model.Condition) (list []model.BackendUserDTO) {
	offset := (condition.Current - 1) * condition.Size
	table := db.Select("id,avatar,nickname,is_disable").Table("user_info")
	if condition.LoginType != 0 {
		table = table.Where("id in (SELECT user_info_id FROM tb_user_auth WHERE login_type = ?)", condition.LoginType)
	}
	if condition.Keywords != "" {
		table = table.Where("nickname like", "%"+condition.Keywords+"%")
	}
	table = table.Limit(condition.Size).Offset(offset)
	db.Select("ua.id,user_info_id,avatar,nickname,login_type,"+
		"ip_address,ip_source,ua.create_time,last_login_time,ui.is_disable").
		Table("(?) ui", table).
		Preload("RoleList").
		Joins("LEFT JOIN user_auth ua ON ua.user_info_id = ui.id").
		Find(&list)
	return
}

func (*UserDao) UpdateUsernameAndEmail(uid int, email string) {
	Updates(&model.UserInfo{Universal: model.Universal{ID: uid}, Email: email})
	db.Select("username").
		Where("user_info_id = ?", uid).
		Updates(&model.UserAuth{Username: email})
}

func (*UserDao) UpdateUserAuthByUsername(data model.UserVO) {
	db.Select("password").Where("username = ?", data.Username).
		Updates(&model.UserAuth{Password: util.Generator.MD5(data.Password)})
}
