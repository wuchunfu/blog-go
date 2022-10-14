package router

import (
	"github.com/gin-contrib/sessions/cookie"
	v1 "myblog/api/v1"
	"myblog/config"
	"net/http"
	"time"
)

var (
	ginConfig  = config.GinConf
	article    v1.Article
	blogInfo   v1.BlogInfo
	category   v1.Category
	comment    v1.Comment
	friendLink v1.FriendLink
	logon      v1.Logon
	menu       v1.Menu
	message    v1.Message
	operation  v1.OperationLog
	page       v1.Page
	photo      v1.Photo
	resource   v1.Resource
	role       v1.Role
	tag        v1.Tag
	talk       v1.Talk
	user       v1.User
	store      = cookie.NewStore([]byte(config.SessionConf.Salt)) // session用来传递一个用户的详细信息
)

func BackendServer() *http.Server {
	return &http.Server{
		Addr:         ginConfig.BackendPort,
		Handler:      backendRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func BlogServer() *http.Server {
	return &http.Server{
		Addr:         ginConfig.Port,
		Handler:      blogRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
