package router

import (
	"gin-ranking/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.GET("/info/:id", controllers.UserController{}.GetInfo)
		user.POST("/list", controllers.UserController{}.PostList)
	}

	return r
}
