package model

import "C"

type UWebsite struct {
	ID         uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name       string `gorm:"type:varchar(255);not null" json:"name"`
	Url        string `gorm:"type:varchar(255);not null" json:"url"`
	CategoryId uint   `gorm:"not null" json:"category_id"`
}

type UWebsiteList struct {
	UserCategoryID   uint   `gorm:"primary_key;auto_increment" json:"user_category_id"`
	UserCategoryName string `gorm:"type:varchar(255);not null" json:"user_category_name"`
	UserID           uint   `gorm:"primary_key;auto_increment" json:"user_id"`
	UserWebsiteID    uint   `gorm:"primary_key;auto_increment" json:"user_website_id"`
	UserWebsiteName  string `gorm:"type:varchar(255);not null" json:"user_website_name"`
	UserWebsiteUrl   string `gorm:"type:varchar(255);not null" json:"user_website_url"`
}

//
//type DefaultWebsiteList struct {
//	CategoryId         uint   `gorm:"primary_key;auto_increment" json:"category_id"`
//	CategoryName       string `gorm:"type:varchar(255);not null" json:"category_name"`
//	CategoryPriority   uint   `gorm:"not null" json:"category_priority"`
//	WebsiteId          uint   `gorm:"primary_key;auto_increment" json:"website_id"`
//	WebsiteName        string `gorm:"type:varchar(255);not null" json:"website_name"`
//	WebsiteUrl         string `gorm:"type:varchar(255);not null" json:"website_url"`
//	WebsiteDescription string `gorm:"type:varchar(255);not null" json:"website_description"`
//}

// CheckWebsite 查询分类是否存在
//func CheckWebsite(name string) (code int) {
//	var cate Website
//	db.Select("id").Where("name = ?", name).First(&cate)
//	if cate.ID > 0 {
//		return errmsg.ERROR_CATENAME_USED //2001
//	}
//	return errmsg.SUCCSE
//}

// CreateCate 新增分类
//func CreateUserWebsite(data *Website) int {
//	err := db.Create(&data).Error
//	if err != nil {
//		return errmsg.ERROR // 500
//	}
//	return errmsg.SUCCSE
//}

//// GetCateInfo 查询单个分类信息
//func GetWebsiteInfo(id int) (UserWebsite, int) {
//	var cate UserWebsite
//	db.Where("id = ?", id).First(&cate)
//	return cate, errmsg.SUCCSE
//}

// GetCate 查询分类列表
//
//	func GetWebsite(pageSize int, pageNum int, websiteName string) []Result {
//		var cate []Result
//
// limit := pageSize
// offset := (pageNum - 1) * pageSize
// db.Raw("SELECT Website.id as id, Website.name as s_name, f_Website.name as f_name, temp.total FROM Website,f_Website,(SELECT COUNT(*) as total from Website,f_Website WHERE Website.f_Website_id=f_Website.id) as temp WHERE Website.f_Website_id=f_Website.id  and Website.`name`LIKE ? LIMIT ? OFFSET ?", "%"+cateName+"%", limit, offset).Scan(&cate)
//
//		return cate
//	}
func GotS() []UWebsiteList {
	var websiteList []UWebsiteList
	username := C.get("username")
	db.Raw("SELECT t1.id as user_category_id,t1.`name` as user_category_name,t1.user_id,t2.id as user_website_id,t2.`name` as user_website_name, t2.url as user_website_url FROM `user_category` as t1,`user_website` as t2,`user` as t3 WHERE t1.user_id=t3.id AND t1.id=t2.category_id and t3.username=?", username).Scan(&websiteList)
	return websiteList
}

// GetArt 查询网站列表
//func GetUserWebsite(id int) []DefaultWebsiteList {
//	var websiteList []DefaultWebsiteList
//	db.Raw("select t2.id as category_id, t2.`name` as category_name,t2.priority as category_priority,t3.id as website_id,t3.`name` as website_name,t3.url as website_url,t3.description as website_description FROM f_category as t1,category as t2,website as t3 WHERE t1.id=? and t2.f_category_id=t1.id and t3.category_id=t2.id ORDER BY t2.priority DESC,t3.priority DESC", id).Scan(&websiteList)
//	return websiteList
//}

// EditCate 编辑分类信息
//func EditUserWebsite(id int, data *Website) int {
//	var cate Website
//	var maps = make(map[string]interface{})
//	maps["name"] = data.Name
//	maps["url"] = data.Url
//	maps["priority"] = data.Priority
//	maps["category_id"] = data.CategoryId
//	maps["description"] = data.Description
//
//	//maps["f_Website_id"] = data.FWebsiteId
//
//	err = db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
//	if err != nil {
//		return errmsg.ERROR
//	}
//	return errmsg.SUCCSE
//}
//
//// DeleteCate 删除分类
//func DeleteUserWebsite(id int) int {
//	var web Website
//	err = db.Where("id = ? ", id).Delete(&web).Error
//	if err != nil {
//		return errmsg.ERROR
//	}
//	return errmsg.SUCCSE
//}
