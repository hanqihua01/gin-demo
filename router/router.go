package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.GET("/get", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello, User Get!")
		})

		user.POST("/post", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello, User Post!")
		})

		user.PUT("/put", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello, User Put!")
		})

		user.DELETE("/delete", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello, User Delete!")
		})
	}

	return r
}
