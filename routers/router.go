package routers

import (
	v1 "ginBlog/api/v1"
	"ginBlog/middleware"
	"ginBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return
	}
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// User module routing
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		auth.PUT("reset/:id", v1.ResetPassword)

		// Category module routing
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)

		// Article module routing
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)

		// Upload file
		auth.POST("upload", v1.Upload)
	}
	routerV1 := r.Group("api/v1")
	{

		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)
		routerV1.GET("user/:id", v1.GetUser)
		routerV1.GET("category/:id", v1.GetCateInfo)
		routerV1.GET("category", v1.GetCategory)
		routerV1.GET("article", v1.GetArticle)
		routerV1.GET("article/list/:id", v1.GetCateArt)
		routerV1.GET("article/info/:id", v1.GetArticleInfo)
		routerV1.POST("login", v1.Login)
	}

	err = r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
