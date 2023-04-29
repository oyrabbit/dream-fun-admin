package model

import (
	"fmt"
	"github.com/oyrabbit/dream-fun-admin/utils/errmsg"
)

type Cate struct {
	ID     uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name   string `gorm:"type:varchar(255);not null" json:"name"`
	UserId uint   `gorm:"not null" json:"user_id"`
}

//
//type Result struct {
//	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
//	SName    string `gorm:"type:varchar();not null" json:"s_name"`
//	FName    string `gorm:"type:varchar();not null" json:"f_name"`
//	Priority uint   `gorm:"" json:"priority"`
//	Total    uint   `gorm:"" json:"total"`
//}
//
//type AllCate struct {
//	FID       uint   `gorm:"primary_key;auto_increment" json:"f_id"`
//	FName     string `gorm:"type:varchar(255);not null" json:"f_name"`
//	FPriority uint   `gorm:"" json:"f_priority"`
//	FIcon     string `gorm:"type:varchar(255);not null" json:"f_icon"`
//	ID        uint   `gorm:"primary_key;auto_increment" json:"id"`
//	Name      string `gorm:"type:varchar(255);not null" json:"name"`
//	Priority  uint   `gorm:"" json:"priority"`
//}

// CheckCategory 查询分类是否存在
//func CheckCategory(name string) (code int) {
//	var cate Category
//	db.Select("id").Where("name = ?", name).First(&cate)
//	if cate.ID > 0 {
//		return errmsg.ERROR_CATENAME_USED //2001
//	}
//	return errmsg.SUCCSE
//}

// CreateCate 新增分类
func CreateUserCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// GetCateInfo 查询单个分类信息
//func GetCateInfo(id int) (Category, int) {
//	var cate Category
//	db.Where("id = ?", id).First(&cate)
//	return cate, errmsg.SUCCSE
//}

// GetCate 查询分类列表
//func GetCate(pageSize int, pageNum int, cateName string) []Result {
//	var cate []Result
//	limit := pageSize
//	offset := (pageNum - 1) * pageSize
//	db.Raw("SELECT category.id as id, category.name as s_name, f_category.name as f_name, category.priority as priority, temp.total FROM category,f_category,(SELECT COUNT(*) as total from category,f_category WHERE category.f_category_id=f_category.id and category.`name`LIKE ?) as temp WHERE category.f_category_id=f_category.id  and category.`name`LIKE ? LIMIT ? OFFSET ?", "%"+cateName+"%", "%"+cateName+"%", limit, offset).Scan(&cate)
//	return cate
//}

// GetCate 查询分类列表
func GetUserCategory() []AllCate {
	var cate []AllCate
	db.Raw("SELECT t1.id as f_id, t1.`name` as f_name, t1.icon as f_icon, t1.priority as f_priority,t2.id as id, t2.`name` as name, t2.priority as priority FROM f_category as t1, category as t2 WHERE t1.id = t2.f_category_id ORDER BY t1.priority DESC,t2.priority DESC").Scan(&cate)
	return cate
}

// GetCate 查询分类列表
func GetCateByUid(fCateId int) []Result {
	var cate []Result
	fmt.Println(fCateId)
	db.Raw("SELECT category.id as id, category.name as s_name, f_category.name as f_name, temp.total FROM category,f_category,(SELECT COUNT(*) as total from category,f_category WHERE category.f_category_id=f_category.id and category.f_category_id=?) as temp WHERE category.f_category_id=f_category.id and category.f_category_id=?", fCateId, fCateId).Scan(&cate)
	return cate
}

// EditCate 编辑分类信息
func EditUserCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["f_category_id"] = data.FCategoryId
	maps["priority"] = data.Priority

	err = db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteCate 删除分类
func DeleteUserCate(id int) int {
	var cate Category
	err = db.Where("id = ? ", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
