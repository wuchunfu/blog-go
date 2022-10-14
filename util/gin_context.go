package util

import (
	"github.com/gin-gonic/gin"
)

func BindJSON[T any](c *gin.Context) (data T) {
	_ = c.ShouldBindJSON(&data)
	return
}

func BindQuery[T any](c *gin.Context) (data T) {
	_ = c.ShouldBindQuery(&data)
	return
}
