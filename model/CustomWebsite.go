package model

import (
	"dream-fun-admin/middleware"
	"dream-fun-admin/utils/errmsg"
	"github.com/gin-gonic/gin"
)

type CustomWebsite struct {
	ID         uint   `gorm:"primary_key;auto_increment;not null" json:"id"`
	Name       string `gorm:"type:varchar(255);not null" json:"name"`
	Url        string `gorm:"type:varchar(255);not null" json:"url"`
	CategoryId uint   `gorm:"not null" json:"category_id"`
}

type CustomWebsiteList struct {
	UserCategoryID   uint   `gorm:"primary_key;auto_increment" json:"user_category_id"`
	UserCategoryName string `gorm:"type:varchar(255);not null" json:"user_category_name"`
	UserID           uint   `gorm:"primary_key;auto_increment" json:"user_id"`
	UserWebsiteID    uint   `gorm:"primary_key;auto_increment" json:"user_website_id"`
	UserWebsiteName  string `gorm:"type:varchar(255);not null" json:"user_website_name"`
	UserWebsiteUrl   string `gorm:"type:varchar(255);not null" json:"user_website_url"`
}

// CreateCate 新增分类
func CreateCustomWebsite(data *CustomWebsite) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// GetWebsites 获取网站列表
func GetCustomWebsite(c *gin.Context) []CustomWebsiteList {
	var websiteList []CustomWebsiteList
	usernames, status := c.Get("username")
	if status {
	}
	usernames1 := usernames.(*middleware.MyClaims)
	username := usernames1.Username
	db.Raw("SELECT t1.id as user_category_id,t1.`name` as user_category_name,t1.user_id,t2.id as user_website_id,t2.`name` as user_website_name, t2.url as user_website_url FROM `custom_category` as t1,`custom_website` as t2,`user` as t3 WHERE t1.user_id=t3.id AND t1.id=t2.category_id and t3.username=?", username).Scan(&websiteList)
	return websiteList
}

// EditCate 编辑分类信息
func EditCustomWebsite(id int, data *CustomWebsite) int {
	var cate CustomWebsite
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["url"] = data.Url
	maps["category_id"] = data.CategoryId

	err = db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// // DeleteCate 删除分类
func DeleteCustomWebsite(id int) int {
	var web CustomWebsite
	err = db.Where("id = ? ", id).Delete(&web).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//func DeleteCustomWebsites(c *gin.Context, id int) int {
//	var cate CustomWebsite
//	usernames, status := c.Get("username")
//	if status {
//		usernames1 := usernames.(*middleware.MyClaims)
//		username := usernames1.Username
//		fmt.Print(username, id)
//		db.Raw("DELETE t1 FROM custom_website t1,custom_category t2,`user` t3 WHERE t3.username=? AND t2.user_id=t3.id AND t1.category_id=t2.id and t1.category_id=?", username, id).Scan(&cate)
//		fmt.Print(cate)
//		if err != nil {
//			return errmsg.ERROR
//		}
//		return errmsg.SUCCSE
//	} else {
//		return errmsg.ERROR
//	}
//}
