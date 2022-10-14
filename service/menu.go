package service

import (
	"myblog/dao"
	"myblog/model"
	"sort"
)

type MenuService struct {
}

func (*MenuService) GetUserMenuList(userInfoId int) []model.UserMenuDTO {
	menuList := menuDao.MenuListByUserInfoId(userInfoId)
	cataList := catalogList(menuList)
	childrenMap := menuMap(menuList)
	return convertUserMenuList(cataList, childrenMap)
}

func (*MenuService) GetList(condition model.Condition) (list []model.MenuDTO) {
	var menuList []model.Menu
	if condition.Keywords != "" {
		menuList = dao.List([]model.Menu{}, "*", "", "resource_name like ?", condition.Keywords)
	} else {
		menuList = dao.List([]model.Menu{}, "*", "", "")
	}
	cataList := catalogList(menuList)
	sort.Slice(cataList, func(i, j int) bool {
		return cataList[i].OrderNum < cataList[j].OrderNum
	})
	childrenMap := menuMap(menuList)
	for _, item := range cataList {
		clist := childrenMap[item.ID]
		menuDTO := model.MenuDTO{
			ID:         item.ID,
			Name:       item.Name,
			Path:       item.Path,
			Component:  item.Component,
			Icon:       item.Icon,
			CreateTime: item.CreateTime,
			OrderNum:   item.OrderNum,
			IsHidden:   item.IsHidden,
		}
		sort.Slice(clist, func(i, j int) bool {
			return clist[i].OrderNum < clist[j].OrderNum
		})
		var mdto []model.MenuDTO
		for _, menu := range clist {
			dto := model.MenuDTO{
				ID:         menu.ID,
				Name:       menu.Name,
				Path:       menu.Path,
				Component:  menu.Component,
				Icon:       menu.Icon,
				CreateTime: menu.CreateTime,
				OrderNum:   menu.OrderNum,
				IsHidden:   menu.IsHidden,
			}
			mdto = append(mdto, dto)
		}
		menuDTO.Children = mdto
		list = append(list, menuDTO)
		delete(childrenMap, item.ID)
	}
	if len(childrenMap) > 0 {
		var childrenList []model.Menu
		for _, v := range childrenMap {
			for _, menu := range v {
				childrenList = append(childrenList, menu)
			}
		}
		sort.Slice(childrenList, func(i, j int) bool {
			return childrenList[i].OrderNum < childrenList[j].OrderNum
		})
		for _, menu := range childrenList {
			dto := model.MenuDTO{
				ID:         menu.ID,
				Name:       menu.Name,
				Path:       menu.Path,
				Component:  menu.Component,
				Icon:       menu.Icon,
				CreateTime: menu.CreateTime,
				OrderNum:   menu.OrderNum,
				IsHidden:   menu.IsHidden,
			}
			list = append(list, dto)
		}
	}
	return
}

func (*MenuService) OptionList() (list []model.LabelOptionDTO) {
	menuList := dao.List([]model.Menu{}, "id,name,parent_id,order_num", "", "")
	cataList := catalogList(menuList)
	childrenMap := menuMap(menuList)
	for _, item := range cataList {
		var clist []model.LabelOptionDTO
		children := childrenMap[item.ID]
		if len(children) > 0 {
			sort.Slice(children, func(i, j int) bool {
				return children[i].OrderNum < children[j].OrderNum
			})
			for _, menu := range children {
				clist = append(clist, model.LabelOptionDTO{
					ID:    menu.ID,
					Label: menu.Name,
				})
			}
		}
		list = append(list, model.LabelOptionDTO{
			ID:       item.ID,
			Label:    item.Name,
			Children: clist,
		})
	}
	return
}

func (*MenuService) SaveOrUpdate(data model.Menu) {
	menuDao.SaveOrUpdate(data)
}

func (s *MenuService) Delete(menuId int) {
	dao.Delete(model.Menu{}, "id = ?", menuId)
}

func catalogList(menuList []model.Menu) (list []model.Menu) {
	for _, menu := range menuList {
		if menu.ParentId == 0 {
			list = append(list, menu)
		}
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].OrderNum < list[j].OrderNum
	})
	return
}

func menuMap(menuList []model.Menu) map[int][]model.Menu {
	mMap := make(map[int][]model.Menu)
	for _, menu := range menuList {
		if menu.ParentId != 0 {
			mMap[menu.ParentId] = append(mMap[menu.ParentId], menu)
		}
	}
	return mMap
}

func convertUserMenuList(catalogList []model.Menu, childrenMap map[int][]model.Menu) (dtoList []model.UserMenuDTO) {
	for _, catalog := range catalogList {
		var userMenuDao model.UserMenuDTO
		var list []model.UserMenuDTO
		children := childrenMap[catalog.ID]
		if len(children) > 0 {
			userMenuDao = model.UserMenuDTO{
				Name:      catalog.Name,
				Path:      catalog.Path,
				Component: catalog.Component,
				Icon:      catalog.Icon,
			}
			sort.Slice(children, func(i, j int) bool {
				return children[i].OrderNum < children[j].OrderNum
			})
			for _, menu := range children {
				dto := model.UserMenuDTO{
					Name:      menu.Name,
					Path:      menu.Path,
					Component: menu.Component,
					Icon:      menu.Icon,
				}
				if menu.IsHidden == 1 {
					dto.Hidden = true
				}
				list = append(list, dto)
			}
		} else {
			userMenuDao.Path = catalog.Path
			userMenuDao.Component = "Layout"
			list = append(list, model.UserMenuDTO{
				Name:      catalog.Name,
				Path:      "",
				Component: catalog.Component,
				Icon:      catalog.Icon,
			})
		}
		if catalog.IsHidden == 1 {
			userMenuDao.Hidden = true
		}
		userMenuDao.Children = list
		dtoList = append(dtoList, userMenuDao)
	}
	return
}
