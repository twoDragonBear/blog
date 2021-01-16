package model

import (
	"blog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// 新增文章
func CreateArticle(data *Article) (code int) {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类下的文章
func GetCateArt(id, pageSize, pageNum int) (article []Article, code int) {
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Where("cid=?",id).Find(&article).Error
	if err != nil {
		return nil,errmsg.ERROR_CATE_NOT_EXIST
	}
	return article,errmsg.SUCCESS
}

// 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

// 查询文章列表
func GetArticle(pageSize int, pageNum int) (article []Article, code int) {
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return article, errmsg.SUCCESS
}

// 编辑文章信息
func EditArticle(id int, data *Article) (code int) {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&Article{}).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除分类
func DeleteArticle(id int) (code int) {
	err = db.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
