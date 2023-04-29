package model

import (
	"github.com/oyrabbit/dream-fun-admin/utils/errmsg"
	"gorm.io/gorm"
)

type FCategory struct {
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Priority uint   `gorm:"" json:"priority"`
	Icon     string `gorm:"type:varchar(255);not null" json:"icon"`
}

// CheckCategory 查询分类是否存在
func CheckFCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //2001
	}
	return errmsg.SUCCSE
}

// CreateCate 新增分类
func CreateFCate(data *FCategory) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// GetCateInfo 查询单个分类信息
func GetFCateInfo(id int) (Category, int) {
	var cate Category
	db.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCSE
}

// GetCate 查询分类列表
func GetFCate(pageSize int, pageNum int, fCateName string) ([]FCategory, int64) {
	var cate []FCategory
	var total int64
	err = db.Find(&cate).Where(
		"name LIKE ?", "%"+fCateName+"%",
	).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

// EditCate 编辑分类信息
func EditFCate(id int, data *FCategory) int {
	var cate FCategory
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["priority"] = data.Priority
	maps["icon"] = data.Icon

	err = db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteCate 删除分类
func DeleteFCate(id int) int {
	var cate FCategory
	err = db.Where("id = ? ", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
