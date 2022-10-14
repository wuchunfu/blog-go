package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"myblog/config"
	"myblog/router/middleware"
	"net/http"
)

func backendRouter() http.Handler {
	gin.SetMode(ginConfig.AppMode)
	router := gin.Default()

	router.Static("web/markdown", "./web/markdown")

	// 尝试使用redis存储这个session但是实现的效果并不好，和我所期待的效果差太多了，可能是对session的一些东西还不够了解
	//store1, _ := redis.NewStore(10, "", "", "", []byte("df321f"))
	//redis.SetKeyPrefix(store, config.SessionConf.Prefix)
	//_, s := redis.GetRedisStore(store)
	//s.SetSerializer(redistore.JSONSerializer{})

	router.Use(middleware.Cors(), sessions.Sessions(config.SessionConf.Name, store))

	api := router.Group("/api")
	{
		api.POST("/report", blogInfo.Report)
		api.POST("/login", logon.Login)
	}
	api.Use(middleware.Jwt(), middleware.RBAC(), middleware.ListenOnline(), middleware.OperationLog())
	api.POST("/logout", logon.Logout)
	api.GET("/about", blogInfo.GetAbout)

	admin := api.Group("/admin")
	{

		admin.GET("", blogInfo.GetBackendInfo)
		admin.PUT("/about", blogInfo.UpdateAbout)
		admin.POST("/config/images", blogInfo.SaveConfigImage)
		admin.GET("/website/config", blogInfo.GetWebsiteConfig)
		admin.PUT("/website/config", blogInfo.UpdateWebsiteConfig)

		articles := admin.Group("/articles")
		{
			articles.GET("", article.GetBackendList)
			articles.POST("/images", article.UploadImages)
			articles.POST("", article.SaveOrUpdate)
			articles.PUT("/top", article.UpdateTop)
			articles.GET("/:id", article.GetBackendById)
			articles.PUT("", article.UpdateDelete)
			articles.DELETE("", article.Delete)
			articles.POST("/export", article.Export)
			articles.POST("/import", article.Import)
		}

		categories := admin.Group("/categories")
		{
			categories.GET("/search", category.GetListBySearch)
			categories.GET("", category.GetBackendList)
			categories.POST("", category.SaveOrUpdate)
			categories.DELETE("", category.Delete)
		}

		comments := admin.Group("/comments")
		{
			comments.GET("", comment.GetBackendList)
			comments.PUT("/review", comment.UpdateReview)
			comments.DELETE("", comment.Delete)
		}

		menus := admin.Group("/menus")
		{
			menus.GET("", menu.GetList)
			menus.POST("", menu.SaveOrUpdate)
			menus.DELETE("/:menuId", menu.Delete)
		}

		messages := admin.Group("/messages")
		{
			messages.GET("", message.GetBackendList)
			messages.PUT("/review", message.UpdateReview)
			messages.DELETE("", message.Delete)
		}

		operationLog := admin.Group("/operation/logs")
		{
			operationLog.GET("", operation.GetList)
			operationLog.DELETE("", operation.Delete)
		}

		links := admin.Group("/links")
		{
			links.GET("", friendLink.GetBackendList)
			links.POST("", friendLink.SaveOrUpdate)
			links.DELETE("", friendLink.Delete)
		}

		pages := admin.Group("/pages")
		{
			pages.GET("", page.GetList)
			pages.POST("", page.SaveOrUpdate)
			pages.DELETE("/:pageId", page.Delete)
		}

		photos := admin.Group("/photos")
		{
			photos.GET("/albums", photo.GetBackendAlbumList)
			photos.POST("/albums/cover", photo.SaveAlbumCover)
			photos.POST("/albums", photo.SaveOrUpdateAlbum)
			photos.DELETE("/albums/:albumId", photo.DeleteAlbum)
			photos.GET("/albums/info", photo.GetBackendAlbumInfoList)
			photos.GET("/albums/:albumId/info", photo.GetBackendAlbumById)
			photos.GET("", photo.GetList)
			photos.POST("", photo.Save)
			photos.PUT("", photo.Update)
			photos.PUT("/album", photo.UpdatePhotosAlbum)
			photos.PUT("/delete", photo.UpdateDelete)
			photos.DELETE("", photo.Delete)
		}

		resources := admin.Group("/resources")
		{
			resources.GET("", resource.GetList)
			resources.POST("", resource.SaveOrUpdate)
			resources.DELETE("/:resourceId", resource.Delete)
		}

		role_ := admin.Group("/role")
		{
			role_.GET("/resources", resource.OptionList)
			role_.GET("/menus", menu.OptionList)
			role_.POST("", role.SaveOrUpdate)
		}

		roles := admin.Group("/roles")
		{
			roles.GET("", role.GetList)
			roles.DELETE("", role.Delete)
		}

		tags := admin.Group("/tags")
		{
			tags.GET("/search", tag.GetListBySearch)
			tags.GET("", tag.GetBackendList)
			tags.POST("", tag.SaveOrUpdate)
			tags.DELETE("", tag.Delete)
		}

		talks := admin.Group("/talks")
		{
			talks.GET("", talk.GetBackendList)
			talks.GET("/:talkId", talk.GetBackendById)
			talks.POST("/images", talk.SaveImages)
			talks.POST("", talk.SaveOrUpdate)
			talks.DELETE("", talk.Delete)
		}

		admin.GET("/user/menus", menu.GetUserMenuList)

		users := admin.Group("/users")
		{
			users.PUT("/password", user.UpdateAdminPassword)
			users.GET("/area", user.GetUserAreas)
			users.GET("", user.GetList)
			users.GET("/role", role.GetUserRoleList)
			users.PUT("/role", user.UpdateRole)
			users.PUT("/disable", user.UpdateDisable)
			users.GET("/online", user.GetOnlineList)
			users.DELETE("/:userInfoId/online", user.ForceOffline)
		}
	}

	users := api.Group("/users")
	{
		users.PUT("/info", user.UpdateInfo)
		users.POST("/avatar", user.UpdateAvatar)
	}
	return router
}
