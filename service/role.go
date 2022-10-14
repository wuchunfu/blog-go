package service

import (
	"myblog/dao"
	"myblog/model"
	"myblog/util/r"
)

type RoleService struct {
}

func (*RoleService) GetUserRoleList() []model.Role {
	return dao.List([]model.Role{}, "id,role_name", "", "")
}

func (*RoleService) GetList(condition model.Condition) model.PageResult[[]model.RoleDTO] {
	var count int
	if condition.Keywords != "" {
		count = dao.Count(model.Role{}, "role_name like ?", "%"+condition.Keywords+"%")
	} else {
		count = dao.Count(model.Role{}, "")
	}
	roleDTOList := roleDao.RoleList(condition)
	return model.PageResult[[]model.RoleDTO]{
		Count:      count,
		RecordList: roleDTOList,
	}
}

func (*RoleService) SaveOrUpdate(data model.RoleVO) int {
	existRole := dao.GetOne(model.Role{}, "role_name", data.RoleName)
	if !existRole.IsEmpty() && existRole.ID != data.ID {
		return r.RoleExist
	}
	role := model.Role{
		Universal: model.Universal{ID: data.ID},
		RoleName:  data.RoleName,
		RoleLabel: data.RoleLabel,
		IsDisable: data.IsDisable,
	}
	roleDao.SaveOrUpdate(&role)
	if len(data.ResourceIdList) > 0 {
		dao.Delete(model.RoleResource{}, "role_id = ?", data.ID)
		dao.RemovePolicy(role.RoleLabel)
		var rslist []model.RoleResource
		for _, rid := range data.ResourceIdList {
			rslist = append(rslist, model.RoleResource{
				RoleId:     role.ID,
				ResourceId: rid,
			})
		}
		rlist := dao.List([]model.Resource{}, "url,request_method", "", "id in ?", data.ResourceIdList)
		var rules [][]string
		for _, resource := range rlist {
			if resource.Url != "" && resource.RequestMethod != "" {
				rules = append(rules, []string{role.RoleLabel, resource.Url, resource.RequestMethod})
			}
		}
		dao.Create(&rslist)
		dao.AddPolicies(rules)
	} else if len(data.MenuIdList) == 0 && len(data.ResourceIdList) == 0 {
		dao.Delete(model.RoleResource{}, "role_id = ?", data.ID)
		dao.RemovePolicy(role.RoleLabel)
	}
	if len(data.MenuIdList) > 0 {
		dao.Delete(model.RoleMenu{}, "role_id = ?", data.ID)
		var rmlist []model.RoleMenu
		for _, mid := range data.MenuIdList {
			rmlist = append(rmlist, model.RoleMenu{
				RoleId: role.ID,
				MenuId: mid,
			})
		}
		dao.Create(&rmlist)
	} else if len(data.MenuIdList) == 0 && len(data.ResourceIdList) == 0 {
		dao.Delete(model.RoleMenu{}, "role_id = ?", data.ID)
	}
	return r.SUCCESS
}

func (*RoleService) Delete(ids []int) {
	roleDao.Delete(ids)
}
