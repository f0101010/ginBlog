package model

import "gorm.io/gorm"

type User struct {
	gorm.Model        // gorm.Model包含ID、CreatedAt、UpdatedAt、DeletedAt
	Username   string `gorm:"type:varchar(20);not null" json:"username"`
	Password   string `gorm:"type:varchar(20);not null" json:"password"`
	Role       int    `gorm:"type:int" json:"role"`
}
