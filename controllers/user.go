package controllers

import (
	"gin-ranking/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetInfo(c *gin.Context) {
	// 通过url获取参数
	idStr := c.Param("id")

	id, _ := strconv.Atoi(idStr)
	user, _ := models.GetUserTest(id)

	ReturnSuccess(c, 200, "success", user, 1)
}

type Param struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (u UserController) PostList(c *gin.Context) {
	// 通过post form获取参数
	// id := c.PostForm("id")
	// name := c.PostForm("name")

	// 通过post json获取参数，使用map接收
	// param := make(map[string]interface{})
	// err := c.BindJSON(&param)

	// 通过post json获取参数，使用struct接收
	param := Param{}
	err := c.BindJSON(&param)

	if err == nil {
		ReturnSuccess(c, 200, "success", param, 1)
	} else {
		ReturnError(c, 400, gin.H{"error": err}) // gin.H = map[string]interface{}
	}
}
