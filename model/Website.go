package model

import (
	"dream-fun-admin/utils/errmsg"
)

type Website struct {
	//Category Category `gorm:"foreignkey:category_id"`
	ID          uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Url         string `gorm:"type:varchar(255);not null" json:"url"`
	Priority    uint   `gorm:"not null" json:"priority"`
	CategoryId  uint   `gorm:"not null" json:"category_id"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
}

type WebsiteList struct {
	ID            uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name          string `gorm:"type:varchar(255);not null" json:"name"`
	CategoryName  string `gorm:"type:varchar(255);not null" json:"category_name"`
	FCategoryName string `gorm:"type:varchar(255);not null" json:"f_category_name"`
	Url           string `gorm:"type:varchar(255);not null" json:"url"`
	Priority      uint   `gorm:"not null" json:"priority"`
	Description   string `gorm:"type:varchar(255);not null" json:"description"`
	Total         uint   `gorm:"" json:"total"`
}

type DefaultWebsiteList struct {
	CategoryId         uint   `gorm:"primary_key;auto_increment" json:"category_id"`
	CategoryName       string `gorm:"type:varchar(255);not null" json:"category_name"`
	CategoryPriority   uint   `gorm:"not null" json:"category_priority"`
	WebsiteId          uint   `gorm:"primary_key;auto_increment" json:"website_id"`
	WebsiteName        string `gorm:"type:varchar(255);not null" json:"website_name"`
	WebsiteUrl         string `gorm:"type:varchar(255);not null" json:"website_url"`
	WebsiteDescription string `gorm:"type:varchar(255);not null" json:"website_description"`
}

// CheckWebsite 查询分类是否存在
func CheckWebsite(name string) (code int) {
	var cate Website
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //2001
	}
	return errmsg.SUCCSE
}

// CreateCate 新增分类
func CreateWebsite(data *Website) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// GetCateInfo 查询单个分类信息
func GetWebsiteInfo(id int) (Website, int) {
	var cate Website
	db.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCSE
}

// GetCate 查询分类列表
//func GetWebsite(pageSize int, pageNum int, websiteName string) []Result {
//	var cate []Result
//limit := pageSize
//offset := (pageNum - 1) * pageSize
//db.Raw("SELECT Website.id as id, Website.name as s_name, f_Website.name as f_name, temp.total FROM Website,f_Website,(SELECT COUNT(*) as total from Website,f_Website WHERE Website.f_Website_id=f_Website.id) as temp WHERE Website.f_Website_id=f_Website.id  and Website.`name`LIKE ? LIMIT ? OFFSET ?", "%"+cateName+"%", limit, offset).Scan(&cate)
//	return cate
//}

// GetArt 查询网站列表
func GetWebsite(pageSize int, pageNum int, webName string) []WebsiteList {
	var websiteList []WebsiteList
	limit := pageSize
	offset := (pageNum - 1) * pageSize
	db.Raw("SELECT website.id as id,website.`name` as name,url,website.priority as priority,website.description as description,category.`name` as category_name,f_category.`name` as f_category_name,temp.total as total FROM website,category,f_category,(SELECT COUNT(*) as total from website,category,f_category  WHERE website.category_id=category.id and f_category.id=category.f_category_id and Website.`name`LIKE ?) as temp WHERE website.category_id=category.id and f_category.id=category.f_category_id and Website.`name`LIKE ?  ORDER BY website.id DESC LIMIT ? OFFSET ?", "%"+webName+"%", "%"+webName+"%", limit, offset).Scan(&websiteList)
	return websiteList
}

// GetArt 查询网站列表
func GetDefaultWebsite(id int) []DefaultWebsiteList {
	var websiteList []DefaultWebsiteList
	db.Raw("select t2.id as category_id, t2.`name` as category_name,t2.priority as category_priority,t3.id as website_id,t3.`name` as website_name,t3.url as website_url,t3.description as website_description FROM f_category as t1,category as t2,website as t3 WHERE t1.id=? and t2.f_category_id=t1.id and t3.category_id=t2.id ORDER BY t2.priority DESC,t3.priority DESC", id).Scan(&websiteList)
	return websiteList
}

// EditCate 编辑分类信息
func EditWebsite(id int, data *Website) int {
	var cate Website
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["url"] = data.Url
	maps["priority"] = data.Priority
	maps["category_id"] = data.CategoryId
	maps["description"] = data.Description

	//maps["f_Website_id"] = data.FWebsiteId

	err = db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteCate 删除分类
func DeleteWebsite(id int) int {
	var web Website
	err = db.Where("id = ? ", id).Delete(&web).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
