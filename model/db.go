package model

import (
	"fmt"
	"ginBlog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	db  *gorm.DB
	err error
)

func InitDb() {
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName)))
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}

	db.AutoMigrate(&User{}, &Article{}, &Category{})

	sqlDB, _ := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
