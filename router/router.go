package router

import (
	"gin-ranking/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.GET("/info/:id/:name", controllers.UserController{}.GetInfo)
		user.POST("/list", controllers.UserController{}.PostList)
	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.PostList)
	}

	return r
}
