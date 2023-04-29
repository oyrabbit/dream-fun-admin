package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/oyrabbit/dream-fun-admin/model"
	"github.com/oyrabbit/dream-fun-admin/utils/errmsg"
	"net/http"
	"strconv"
)

//// AddWebsite 添加分类
//func AddUserWebsite(c *gin.Context) {
//	var data model.Website
//	_ = c.ShouldBindJSON(&data)
//	code := model.CheckWebsite(data.Name)
//	if code == errmsg.SUCCSE {
//		model.CreateWebsite(&data)
//	}
//
//	c.JSON(
//		http.StatusOK, gin.H{
//			"status":  code,
//			"data":    data,
//			"message": errmsg.GetErrMsg(code),
//		},
//	)
//}

//
//// GetCateInfo 查询分类信息
//func GetWebsiteInfo(c *gin.Context) {
//	id, _ := strconv.Atoi(c.Param("id"))
//
//	data, code := model.GetWebsiteInfo(id)
//
//	c.JSON(
//		http.StatusOK, gin.H{
//			"status":  code,
//			"data":    data,
//			"message": errmsg.GetErrMsg(code),
//		},
//	)
//
//}

// GetCate 查询分类列表
func GetWebsites(c *gin.Context) {
	data := model.GetWebsites()
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

//func GetDefaultWebsite(c *gin.Context) {
//	id, _ := strconv.Atoi(c.Param("id"))
//
//	data := model.GetDefaultWebsite(id)
//	code := errmsg.SUCCSE
//	c.JSON(
//		http.StatusOK, gin.H{
//			"status": code,
//			"data":   data,
//			//"total":   total,
//			"message": errmsg.GetErrMsg(code),
//		},
//	)
//}

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
func EditUserWebsite(c *gin.Context) {
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
func DeleteUserWebsite(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteWebsite(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}