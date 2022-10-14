package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"myblog/config"
	"myblog/router/middleware"
	"net/http"
)

func blogRouter() http.Handler {
	gin.SetMode(ginConfig.AppMode)
	router := gin.Default()

	router.Static("web/markdown", "./web/markdown")

	router.Use(middleware.Cors(), sessions.Sessions(config.SessionConf.Name, store))

	api := router.Group("/api")
	{
		api.GET("/", blogInfo.GetHomeInfo)
		api.POST("/login", logon.Login)
		api.GET("/logout", logon.Logout)
		api.GET("/users/code", user.SendCode)
		api.POST("/register", user.Register)
		api.PUT("/users/password", user.UpdatePassword)
		api.POST("/report", blogInfo.Report)
		api.GET("/about", blogInfo.GetAbout)

		articles := api.Group("/articles")
		{
			articles.GET("", article.GetList)
			articles.GET("/:articleId", article.GetById)
			articles.GET("/search", article.GetListBySearch)
			articles.GET("/archives", article.GetArchiveList)
			articles.GET("/condition", article.GetListByCondition)
		}

		categories := api.Group("/categories")
		{
			categories.GET("", category.GetList)
		}

		comments := api.Group("/comments")
		{
			comments.GET("", comment.GetList)
			comments.GET("/:commentId/replies", comment.GetReplyListByCommentId)
		}

		messages := api.Group("/messages")
		{
			messages.GET("", message.GetList)
		}

		links := api.Group("/links")
		{
			links.GET("", friendLink.GetList)
		}

		photos := api.Group("/photos")
		{
			photos.GET("/albums", photo.GetAlbumList)
			api.GET("/albums/:albumId/photos", photo.GetListByAlbumId)
		}

		tags := api.Group("/tags")
		{
			tags.GET("", tag.GetList)
		}

		talks := api.Group("/talks")
		{
			api.GET("/home/talks", talk.GetHomeList)
			talks.GET("", talk.GetList)
			talks.GET("/:talkId", talk.GetById)
		}

	}

	api.POST("/users/avatar", user.UpdateAvatar) // 额，这里前端不知怎么带上 Authorization 了
	// 下面是需要登陆才能进行的操作
	api.Use(middleware.Jwt(), middleware.RBAC(), middleware.ListenOnline())
	{
		api.POST("/articles/:articleId/like", article.SaveLike)
		api.POST("/talks/:talkId/like", talk.SaveLike)
		api.POST("/comments", comment.SaveComment)
		api.POST("/comments/:commentId/like", comment.SaveLike)
		api.PUT("/users/info", user.UpdateInfo)
		api.POST("/users/email", user.SaveEmail)
		api.POST("/messages", message.SaveMessage)
	}

	return router
}
