package model

import (
	"dream-fun-admin/middleware"
	"dream-fun-admin/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CustomCategory struct {
	ID     uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name   string `gorm:"type:varchar(255);not null" json:"name"`
	UserId uint   `gorm:"not null" json:"user_id"`
}

// CheckCategory 查询分类是否存在
func CheckCustomCategory(name string) (code int) {
	var cate CustomCategory
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //2001
	}
	return errmsg.SUCCSE
}

// CreateCate 新增分类
func CreateCustomCategory(data *CustomCategory) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// GetCate 查询分类列表
func GetCustomCategory(c *gin.Context) ([]CustomCategory, int) {
	var cate []CustomCategory
	usernames, status := c.Get("username")
	if status {
		usernames1 := usernames.(*middleware.MyClaims)
		username := usernames1.Username
		db.Raw("SELECT t1.id as id,t1.`name` as name FROM `custom_category` as t1,`user` as t2 WHERE t1.user_id=t2.id AND t2.username=?", username).Scan(&cate)
		if err != nil {
			return cate, errmsg.ERROR
		}
		return cate, errmsg.SUCCSE
	} else {
		return cate, errmsg.ERROR
	}
}

// EditCate 编辑分类信息
func EditCustomCate(c *gin.Context, id int, data *CustomCategory) int {
	var cate CustomCategory
	usernames, status := c.Get("username")
	if status {
		usernames1 := usernames.(*middleware.MyClaims)
		username := usernames1.Username
		name := data.Name
		db.Raw("UPDATE custom_category as t1,`user` as t2 SET t1.`name` = ? WHERE t1.user_id=t2.id AND t2.username=? AND t1.id=?", name, username, id).Scan(&cate)
		if err != nil {
			return errmsg.ERROR
		}
		return errmsg.SUCCSE
	} else {
		return errmsg.ERROR
	}
}

// DeleteCate 删除分类
func DeleteCustomCate(c *gin.Context, id int) int {
	var cate CustomCategory
	usernames, status := c.Get("username")
	if status {
		usernames1 := usernames.(*middleware.MyClaims)
		username := usernames1.Username
		db.Raw("DELETE t1 FROM custom_website t1,custom_category t2,`user` t3 WHERE t3.username=? AND t2.user_id=t3.id AND t1.category_id=t2.id and t1.category_id=?", username, id).Scan(&cate)
		db.Raw("DELETE t1 FROM custom_category t1 LEFT JOIN `user` t2 on t1.user_id=t2.id WHERE t2.username=? AND t1.id=?", username, id).Scan(&cate)
		fmt.Print(cate)
		if err != nil {
			return errmsg.ERROR
		}
		return errmsg.SUCCSE
	} else {
		return errmsg.ERROR
	}
}
