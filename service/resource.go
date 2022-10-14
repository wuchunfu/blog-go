package service

import (
	"myblog/dao"
	"myblog/model"
)

type ResourceService struct {
}

func (*ResourceService) GetList(condition model.Condition) (list []model.ResourceDTO) {
	var resourceList []model.Resource
	if condition.Keywords != "" {
		resourceList = dao.List([]model.Resource{}, "*", "",
			"resource_name like ?", condition.Keywords)
	} else {
		resourceList = dao.List([]model.Resource{}, "*", "", "")
	}
	parentList := resourceModuleList(resourceList)
	childrenMap := resourceChildrenList(resourceList)
	for _, item := range parentList {
		clist := childrenMap[item.ID]
		resourceDTO := model.ResourceDTO{
			ID:            item.ID,
			ResourceName:  item.ResourceName,
			Url:           item.Url,
			RequestMethod: item.RequestMethod,
			IsAnonymous:   item.IsAnonymous,
			CreateTime:    item.CreateTime,
		}
		var rdto []model.ResourceDTO
		for _, child := range clist {
			dto := model.ResourceDTO{
				ID:            child.ID,
				ResourceName:  child.ResourceName,
				Url:           child.Url,
				RequestMethod: child.RequestMethod,
				IsAnonymous:   child.IsAnonymous,
				CreateTime:    child.CreateTime,
			}
			rdto = append(rdto, dto)
		}
		resourceDTO.Children = rdto
		list = append(list, resourceDTO)
	}
	return
}

func (*ResourceService) OptionList() (list []model.LabelOptionDTO) {
	resourceList := dao.List([]model.Resource{}, "id,resource_name,parent_id", "", "is_anonymous = ?", 0)
	parentList := resourceModuleList(resourceList)
	childrenMap := resourceChildrenList(resourceList)
	for _, item := range parentList {
		var clist []model.LabelOptionDTO
		children := childrenMap[item.ID]
		if len(children) > 0 {
			for _, re := range children {
				clist = append(clist, model.LabelOptionDTO{
					ID:    re.ID,
					Label: re.ResourceName,
				})
			}
		}
		list = append(list, model.LabelOptionDTO{
			ID:       item.ID,
			Label:    item.ResourceName,
			Children: clist,
		})
	}
	return
}

func (*ResourceService) SaveOrUpdate(data model.Resource) {
	resourceDao.SaveOrUpdate(data)
}

func (*ResourceService) Delete(resourceId int) {
	resourceDao.Delete(resourceId)
}

func resourceChildrenList(resourceList []model.Resource) map[int][]model.Resource {
	childrenMap := make(map[int][]model.Resource)
	for _, resource := range resourceList {
		if resource.ParentId != 0 {
			childrenMap[resource.ParentId] = append(childrenMap[resource.ParentId], resource)
		}
	}
	return childrenMap
}

func resourceModuleList(resourceList []model.Resource) (list []model.Resource) {
	for _, resource := range resourceList {
		if resource.ParentId == 0 {
			list = append(list, resource)
		}
	}
	return
}
