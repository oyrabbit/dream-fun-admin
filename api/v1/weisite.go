package v1

import (
	"dream-fun-admin/model"
	"dream-fun-admin/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddWebsite 添加分类
func AddWebsite(c *gin.Context) {
	var data model.Website
	_ = c.ShouldBindJSON(&data)
	code := model.CreateWebsite(&data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetCateInfo 查询分类信息
func GetWebsiteInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetWebsiteInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

// GetCate 查询分类列表
func GetWebsite(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	webName := c.Query("webname")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data := model.GetWebsite(pageSize, pageNum, webName)
	code := errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status": code,
			"data":   data,
			//"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

func GetDefaultWebsite(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data := model.GetDefaultWebsite(id)
	code := errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status": code,
			"data":   data,
			//"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 查询单个分类
//func GetCateInfo(c *gin.Context)  {
//	id, _ := strconv.Atoi(c.Param("id"))
//
//	data,code := model.GetCateInfo(id)
//
//	c.JSON(http.StatusOK, gin.H{
//		"status":  code,
//		"data":    data,
//		"message": errmsg.GetErrMsg(code),
//	})
//}

// EditCate 编辑分类名
func EditWebsite(c *gin.Context) {
	var data model.Website
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.EditWebsite(id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteCate 删除用户
func DeleteWebsite(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteWebsite(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
