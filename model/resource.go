package model

import "time"

// Resource == casbin rule
type Resource struct {
	Universal
	Url           string `json:"url"`
	RequestMethod string `json:"requestMethod"`
	ResourceName  string `json:"resourceName"`
	ParentId      int    `json:"parentId"`
	IsAnonymous   int    `json:"isAnonymous"`
}

type ResourceDTO struct {
	ID            int           `json:"id"`
	ResourceName  string        `json:"resourceName"`
	Url           string        `json:"url"`
	RequestMethod string        `json:"requestMethod"`
	IsAnonymous   int           `json:"isAnonymous"`
	CreateTime    time.Time     `json:"createTime"`
	Children      []ResourceDTO `json:"children"`
}
