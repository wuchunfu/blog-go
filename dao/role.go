package dao

import "myblog/model"

type RoleDao struct {
}

func (*RoleDao) RoleListByUserInfoId(userInfoId int) (list []string) {
	db.Select("role_label").Table("role r, user_role ur").
		Where("r.id = ur.role_id AND ur.user_id = ?", userInfoId).
		Find(&list)
	return
}

func (*RoleDao) RoleList(condition model.Condition) (re []model.RoleDTO) {
	var list []model.RoleDTO
	offset := (condition.Current - 1) * condition.Size
	tx := db.
		Select("id,role_name,role_label,create_time,is_disable").
		Table("role")
	if condition.Keywords != "" {
		tx = tx.Where("name like ?", "%"+condition.Keywords+"%")
	}
	table := tx.Limit(condition.Size).Offset(offset)
	db.Select("r.id,role_name,role_label,r.create_time,r.is_disable").
		Table("(?) r", table).
		Order("r.id").Find(&list)
	for _, role := range list {
		var resourceIdList, menuIdList []int
		db.Select("rr.resource_id").
			Table("(?) r", table).
			Joins("LEFT JOIN role_resource rr ON r.id = rr.role_id").
			Where("r.id = ?", role.ID).
			Find(&resourceIdList)
		db.Select("rm.menu_id").
			Table("(?) r", table).
			Joins("LEFT JOIN role_menu rm on r.id = rm.role_id").
			Where("r.id = ?", role.ID).
			Find(&menuIdList)
		role.ResourceIdList = resourceIdList
		role.MenuIdList = menuIdList
		re = append(re, role)
	}
	return
}

func (*RoleDao) SaveOrUpdate(role *model.Role) {
	if role.ID == 0 {
		Create(role)
	}
	rlist := selectResourceByRoleLabel(role.RoleLabel)
	Updates(role)
	for _, r := range rlist {
		Updates(&r, "v0")
	}
}

func (*RoleDao) Delete(ids []int) {
	var rllist []string
	db.Select("role_label").Table("role").Where("id in ?", ids).Find(&rllist)
	db.Exec("DELETE FROM casbin_rule WHERE v0 in ?", rllist)
	Delete(model.RoleMenu{}, "role_id in ?", ids)
	Delete(model.RoleResource{}, "role_id in ?", ids)
	Delete(model.Role{}, "id in ?", ids)
}

func selectResourceByRoleLabel(roleLabel string) (re []model.Resource) {
	rLable := GetOne(model.Role{}, "role_label = ?", roleLabel).RoleLabel
	db.Model(model.Resource{}).Where("ptype = 'g' AND v0 = ?", rLable).Find(&re)
	return
}
