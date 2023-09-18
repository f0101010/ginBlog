package routers

import (
	v1 "ginBlog/api/v1"
	"ginBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	routerV1 := r.Group("api/v1")
	{
		// User module routing
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUsers)
		routerV1.PUT("user/:id", v1.EditUser)
		routerV1.DELETE("user/:id", v1.DeleteUser)

		// Category module routing
		routerV1.POST("category/add", v1.AddCategory)
		routerV1.GET("category", v1.GetCategory)
		routerV1.PUT("category/:id", v1.EditCategory)
		routerV1.DELETE("category/:id", v1.DeleteCategory)

		// Article module routing
		routerV1.POST("article/add", v1.AddArticle)
		routerV1.GET("article", v1.GetArticle)
		routerV1.PUT("article/:id", v1.EditArticle)
		routerV1.DELETE("article/:id", v1.DeleteArticle)
	}
	r.Run(utils.HttpPort)
}
