package v1

import (
	"dream-fun-admin/middleware"
	"dream-fun-admin/model"
	"dream-fun-admin/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func EditCustomWebsite(c *gin.Context) {
	var data model.CustomWebsite
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	usernames, status := c.Get("username")
	usernames1 := usernames.(*middleware.MyClaims)
	username := usernames1.Username
	name := c.Param("username")
	if status {
	}
	var code int
	if username == name {
		model.EditCustomWebsite(id, &data)
	} else {
		code = 1008
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteCate 删除用户
func DeleteWebsites(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteWebsite(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetWebsites 获取网站
func GetWebsites(c *gin.Context) {

	data := model.GetCustomWebsite(c)
	code := errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// AddWebsite 添加分类
func AddCustomWebsite(c *gin.Context) {
	var data model.CustomWebsite
	_ = c.ShouldBindJSON(&data)
	usernames, status := c.Get("username")
	usernames1 := usernames.(*middleware.MyClaims)
	username := usernames1.Username
	id := c.Param("username")
	if status {
	}
	var code int
	if username == id {
		code = model.CreateCustomWebsite(&data)
	} else {
		code = 1008
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteCate 删除用户
func DeleteCustomWebsite(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	usernames, status := c.Get("username")
	usernames1 := usernames.(*middleware.MyClaims)
	username := usernames1.Username
	name := c.Param("username")
	if status {
	}
	var code int
	if username == name {
		code = model.DeleteCustomWebsite(id)
	} else {
		code = 1008
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
