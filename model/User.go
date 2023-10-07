package model

import (
	"encoding/base64"
	"errors"
	"ginBlog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model        // gorm.Model包含ID、CreatedAt、UpdatedAt、DeletedAt
	Username   string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password   string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20"label:"密码"`
	Role       int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// CheckUserExist 查询用户是否存在
func CheckUserExist(username string) (code int) {
	var users User
	db.Select("id").Where("username = ?", username).First(&users)
	if users.ID > 0 {
		return errmsg.ErrorUsernameUsed // 1001
	}
	return errmsg.SUCCESS // 200
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	//data.Password = ScryptPassword(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

// GetUsers 查询用户列表
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var user User
	var users []User
	var total int64

	db.Model(&user).Count(&total)
	if username == "" {
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	}
	err = db.Where("username LIKE ?", username+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error

	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		return nil, 0
	}
	return users, total
}

// EditUser 编辑用户信息
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password = ScryptPassword(u.Password)
	return nil
}

// ScryptPassword 密码加密
func ScryptPassword(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{13, 24, 98, 4, 88, 23, 45, 67}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)

	return fpw
}

// CheckLogin 登录验证
func CheckLogin(username, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ErrorUserNotExist
	}
	if ScryptPassword(password) != user.Password {
		return errmsg.ErrorPasswordWrong
	}
	if user.Role != 1 {
		return errmsg.ErrorUserNoRight
	}
	return errmsg.SUCCESS
}
