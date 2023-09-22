package model

import (
	"errors"
	"ginBlog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignKey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200);" json:"desc"`
	Content string `gorm:"type:longtext;" json:"content"`
	Img     string `gorm:"type:varchar(100);" json:"img"`
}

// CreateArticle 添加文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

// GetCateArt 查询分类下的所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64
	var err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArtList).Count(&total).Error
	if err != nil {
		return nil, errmsg.ErrorCateNotExist, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}

// GetArticleInfo 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var art Article
	err = db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return Article{}, errmsg.ErrorArtNotExist
	}
	return art, errmsg.SUCCESS
}

// GetArticle 查询文章列表
func GetArticle(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Count(&total).Error
	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// EditArticle 编辑文章
func EditArticle(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

// DeleteArticle 删除文章
func DeleteArticle(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
