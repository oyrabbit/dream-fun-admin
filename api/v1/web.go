package v1

import "C"

//type CustomWebsiteList struct {
//	UserCategoryID   uint   `gorm:"primary_key;auto_increment" json:"user_category_id"`
//	UserCategoryName string `gorm:"type:varchar(255);not null" json:"user_category_name"`
//	UserID           uint   `gorm:"primary_key;auto_increment" json:"user_id"`
//	UserWebsiteID    uint   `gorm:"primary_key;auto_increment" json:"user_website_id"`
//	UserWebsiteName  string `gorm:"type:varchar(255);not null" json:"user_website_name"`
//	UserWebsiteUrl   string `gorm:"type:varchar(255);not null" json:"user_website_url"`
//}

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

//	func GetDefaultWebsite(c *gin.Context) {
//		id, _ := strconv.Atoi(c.Param("id"))
//
//		data := model.GetDefaultWebsite(id)
//		code := errmsg.SUCCSE
//		c.JSON(
//			http.StatusOK, gin.H{
//				"status": code,
//				"data":   data,
//				//"total":   total,
//				"message": errmsg.GetErrMsg(code),
//			},
//		)
//	}
//func GetWebsites(c *gin.Context) {
//	var db *gorm.DB
//	//id, _ := strconv.Atoi(c.Param("id"))
//	var websiteList []CustomWebsiteList
//	//fmt.Println(id)
//	username := C.get("username")
//	db.Raw("SELECT t1.id as user_category_id,t1.`name` as user_category_name,t1.user_id,t2.id as user_website_id,t2.`name` as user_website_name, t2.url as user_website_url FROM `cate` as t1,`web` as t2,`user` as t3 WHERE t1.user_id=t3.id AND t1.id=t2.category_id and t3.username=?", username).Scan(&websiteList)
//	//data := model.GetCustomWebsite(id)
//	code := errmsg.SUCCSE
//	c.JSON(
//		http.StatusOK, gin.H{
//			"status": code,
//			"data":   websiteList,
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
//func EditUserWebsite(c *gin.Context) {
//	var data model.Website
//	id, _ := strconv.Atoi(c.Param("id"))
//	_ = c.ShouldBindJSON(&data)
//	code := model.EditWebsite(id, &data)
//
//	c.JSON(
//		http.StatusOK, gin.H{
//			"status":  code,
//			"message": errmsg.GetErrMsg(code),
//		},
//	)
//}
//
//// DeleteCate 删除用户
//func DeleteUserWebsite(c *gin.Context) {
//	id, _ := strconv.Atoi(c.Param("id"))
//
//	code := model.DeleteWebsite(id)
//
//	c.JSON(
//		http.StatusOK, gin.H{
//			"status":  code,
//			"message": errmsg.GetErrMsg(code),
//		},
//	)
//}
