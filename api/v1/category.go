package v1

import (
	"dream-fun-admin/model"
	"dream-fun-admin/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCSE {
		model.CreateCate(&data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetCateInfo 查询分类信息
func GetCateInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetCateInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetCate 查询分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	cateName := c.Query("catename")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data := model.GetCate(pageSize, pageNum, cateName)
	code := errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status": code,
			"data":   data,
			//"total":   data.total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetAllCate 查询前台分类列表
func GetAllCate(c *gin.Context) {

	data := model.GetAllCate()
	code := errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status": code,
			"data":   data,
			//"total":   data.total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetCate 查询分类列表
func GetCateByFid(c *gin.Context) {

	fCateId, _ := strconv.Atoi(c.Query("fcateid"))

	data := model.GetCateByFid(fCateId)
	code := errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status": code,
			"data":   data,
			//"total":   data.total,
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
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.EditCate(id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteCate 删除用户
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteCate(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
