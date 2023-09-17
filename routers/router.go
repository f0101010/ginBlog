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

		// Article module routing
	}
	r.Run(utils.HttpPort)
}
