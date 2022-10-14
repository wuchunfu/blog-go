package dao

import "myblog/model"

type MenuDao struct {
}

func (*MenuDao) MenuListByUserInfoId(userInfoId int) (list []model.Menu) {
	db.Table("user_role ur").
		Distinct("m.id", "name", "path", "component", "icon", "is_hidden", "parent_id", "order_num").
		Where("user_id = ?", userInfoId).
		Joins("JOIN role_menu rm ON ur.role_id = rm.role_id").
		Joins("JOIN menu m ON rm.menu_id = m.id").
		Find(&list)
	return
}

func (*MenuDao) SaveOrUpdate(data model.Menu) {
	if data.ID != 0 {
		Updates(&data, "name", "path", "component", "icon", "order_num", "is_hidden")
	} else {
		Create(&data)
	}
}
