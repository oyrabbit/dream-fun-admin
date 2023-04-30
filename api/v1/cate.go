package v1

import (
	"dream-fun-admin/middleware"
	"dream-fun-admin/model"
	"dream-fun-admin/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddCategory 添加分类
func AddCustomCategory(c *gin.Context) {
	var data model.CustomCategory
	_ = c.ShouldBindJSON(&data)
	usernames, status := c.Get("username")
	usernames1 := usernames.(*middleware.MyClaims)
	username := usernames1.Username
	id := c.Param("username")
	if status {
	}
	var code int
	if username == id {
		code = model.CreateCustomCategory(&data)
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

// GetCate 查询分类列表
func GetCustomCategory(c *gin.Context) {
	data, code := model.GetCustomCategory(c)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// EditCate 编辑分类名
func EditCustomCate(c *gin.Context) {
	var data model.CustomCategory
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.EditCustomCate(c, id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteCate 删除用户
func DeleteCustomCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteCustomCate(c, id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
