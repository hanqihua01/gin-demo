package controllers

import "github.com/gin-gonic/gin"

type OrderController struct{}

func (o OrderController) PostList(c *gin.Context) {
	ReturnError(c, 400, "没有相关信息")
}
