package model

import (
	"blog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

// 新增分类
func CreateCategory(data *Category) (code int) {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// todo 查询分类下的文章

// 查询分类列表
func GetCategory(pageSize int, pageNum int) (cate []Category) {
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

// 编辑分类信息
func EditCategory(id int, data *Category) (code int) {
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&Category{}).Where("id = ?",id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除分类
func DeleteCategory(id int) (code int) {
	err = db.Where("id = ?",id).Delete(&Category{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
