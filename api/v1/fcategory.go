package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/oyrabbit/dream-fun-admin/model"
	"github.com/oyrabbit/dream-fun-admin/utils/errmsg"
	"net/http"
	"strconv"
)

// AddFCategory 添加分类
func AddFCategory(c *gin.Context) {
	var data model.FCategory
	_ = c.ShouldBindJSON(&data)
	code := model.CheckFCategory(data.Name)
	if code == errmsg.SUCCSE {
		model.CreateFCate(&data)
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
func GetFCateInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetFCateInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

// GetFCate 查询分类列表
func GetFCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	fCateName := c.Query("fcatename")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, _ := model.GetFCate(pageSize, pageNum, fCateName)

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
//func GetFCateInfo(c *gin.Context)  {
//	id, _ := strconv.Atoi(c.Param("id"))
//
//	data,code := model.GetFCateInfo(id)
//
//	c.JSON(http.StatusOK, gin.H{
//		"status":  code,
//		"data":    data,
//		"message": errmsg.GetErrMsg(code),
//	})
//}

// EditCate 编辑分类名
func EditFCate(c *gin.Context) {
	var data model.FCategory
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.CheckFCategory(data.Name)
	if code == errmsg.SUCCSE {
		model.EditFCate(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteFCate 删除用户
func DeleteFCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteFCate(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
