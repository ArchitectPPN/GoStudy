package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

// @Summary Get an user
// @Description 获取用户信息
// @Tags 用户
// @Param id path int true "ID"
// @Success 200 {string} string "ok"
// @Router /user/{id} [get]
func (c *Controller) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	users := make(map[string]string)
	users["1"] = "Tom"
	users["2"] = "Jerry"
	ctx.String(http.StatusOK, "Hello %s", users[id])
}

type Res struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	UserType string `json:"userType"`
}

// @Summary Add an user
// @Description 提交用户信息
// @Tags 用户
// @Param id formData int true "ID"
// @Param name formData string true "姓名"
// @Param userType formData string true "类别"
// @Success 200 {object} Res
// @Router /user [post]
func (c *Controller) AddUser(ctx *gin.Context) {
	id := ctx.PostForm("id")
	name := ctx.PostForm("name")
	userType := ctx.DefaultPostForm("userType", "student")

	ctx.JSON(200, Res{
		Id:       id,
		Name:     name,
		UserType: userType,
	})
}