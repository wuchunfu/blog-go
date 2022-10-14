package dao

import (
	"myblog/model"
)

type ResourceDao struct {
}

func (*ResourceDao) SaveOrUpdate(data model.Resource) {
	if data.ID != 0 {
		Updates(&data, "resource_name", "url", "request_method", "is_anonymous")
		if data.IsAnonymous == 0 {
			e.RemovePolicy("anonymous", data.Url, data.RequestMethod)
		} else {
			e.AddPolicy("anonymous", data.Url, data.RequestMethod)
		}
	} else {
		Create(&data)
	}
}

func (*ResourceDao) Delete(resourceId int) {
	resource := GetOne(model.Resource{}, "id = ?", resourceId)
	Delete(model.Resource{}, "id", resourceId)
	db.Table("casbin_rule").Where("v1 = ? AND v2 = ?", resource.Url, resource.RequestMethod)
}
