package model

import (
	"errors"
	"ginBlog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategoryExist 查询分类是否存在
func CheckCategoryExist(username string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", username).First(&cate)
	if cate.ID > 0 {
		return errmsg.ErrorCatenameUsed // 1001
	}
	return errmsg.SUCCESS // 200
}

// CreateCategory 新增分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

// GetCategory 查询分类列表
func GetCategory(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Count(&total).Error
	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		return nil, 0
	}
	return cate, total
}

// EditCategory 编辑分类信息
func EditCategory(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

// DeleteCategory 删除分类
func DeleteCategory(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
