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
	Username   string `gorm:"type:varchar(20);not null" json:"username"`
	Password   string `gorm:"type:varchar(20);not null" json:"password"`
	Role       int    `gorm:"type:int" json:"role"`
}

// CheckUserExist 查询用户是否存在
func CheckUserExist(username string) (code int) {
	var users User
	db.Select("id").Where("username = ?", username).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED // 1001
	}
	return errmsg.SUCCESS // 200
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPassword(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		return nil
	}
	return users
}

// EditUser 编辑用户

// DeleteUser 删除用户

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
