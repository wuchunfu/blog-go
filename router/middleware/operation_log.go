package middleware

import (
	"bytes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"myblog/dao"
	"myblog/model"
	"strings"
)

var optMap = map[string]string{
	"Article":      "文章",
	"BlogInfo":     "博客信息",
	"Category":     "分类",
	"Comment":      "评论",
	"FriendLink":   "友链",
	"Logon":        "登录",
	"Menu":         "菜单",
	"Message":      "留言",
	"OperationLog": "操作日志",
	"Page":         "页面",
	"Photo":        "相册照片",
	"Resource":     "资源权限",
	"Role":         "角色",
	"Tag":          "标签",
	"Talk":         "说说",
	"User":         "用户",
	"POST":         "新增或修改",
	"PUT":          "修改",
	"DELETE":       "删除",
}

func GetOptString(structName string) string {
	return optMap[structName]
}

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer // 写响应时会写入缓存
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func OperationLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			blw := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
			c.Writer = blw
			uuid, _ := c.Get("uuid")
			sessionInfo := sessions.Default(c).Get(User + uuid.(string)).(model.SessionInfo)
			body, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			operationLog := model.OperationLog{
				OptModule:     GetOptString(strings.Split(c.HandlerName(), ".")[1]),
				OptType:       GetOptString(c.Request.Method),
				OptUrl:        c.Request.RequestURI,
				OptMethod:     c.HandlerName(),
				OptDesc:       GetOptString(c.Request.Method), // 不方便搞了就这样。。
				RequestParam:  string(body),
				RequestMethod: c.Request.Method,
				UserId:        sessionInfo.UserInfoId,
				Nickname:      sessionInfo.Nickname,
				IpAddress:     sessionInfo.IpAddress,
				IpSource:      sessionInfo.IpSource,
			}
			c.Next()
			operationLog.ResponseData = blw.body.String() // 从缓存中获取响应体内容
			dao.Create(&operationLog)
		} else {
			c.Next()
		}
	}
}
